package customroundrobin

import (
	"encoding/json"
	"fmt"
	"sync/atomic"

	_ "google.golang.org/grpc" // to register pick_first
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/endpointsharding"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/serviceconfig"
)

var gracefulSwitchPickFirst serviceconfig.LoadBalancingConfig

func init() {
	balancer.Register(customRoundRobinBuilder{})
	var err error
	gracefulSwitchPickFirst, err = endpointsharding.ParseConfig(json.RawMessage(endpointsharding.PickFirstConfig))
	if err != nil {
		logger.Fatal(err)
	}
}

const customRRName = "custom_round_robin"

type customRRConfig struct {
	serviceconfig.LoadBalancingConfig `json:"-"`

	// json 导出符号需要大写
	ChooseSecond uint32 `json:"chooseSecond,omitempty"`
}

type customRoundRobinBuilder struct {
}

func (customRoundRobinBuilder) ParseConfig(s json.RawMessage) (serviceconfig.LoadBalancingConfig, error) {
	ldConfig := &customRRConfig{
		ChooseSecond: 3,
	}

	if err := json.Unmarshal(s, ldConfig); err != nil {
		return nil, fmt.Errorf("custom-round-robin: unable to unmarshal customRRConfig: %v", err)
	}

	return ldConfig, nil
}

func (customRoundRobinBuilder) Name() string {
	return customRRName
}

func (customRoundRobinBuilder) Build(cc balancer.ClientConn, bOpts balancer.BuildOptions) balancer.Balancer {
	crr := &customRoundRobin{
		ClientConn: cc,
		bOpts:      bOpts,
	}
	crr.Balancer = endpointsharding.NewBalancer(crr, bOpts)
	return crr
}

var logger = grpclog.Component("example")

type customRoundRobin struct {
	// 此负载均衡器的所有状态和操作要么在构建时初始化并在之后只读，要么仅作为其
	// balancer.Balancer API 的一部分进行访问（来自子负载均衡器的 UpdateState 也仅来自
	// balancer.Balancer 调用，并且子负载均衡器一次一个地被调用），
	// 在这些调用中保证是同步进行的。因此，此负载均衡器不需要额外的同步。
	balancer.Balancer
	balancer.ClientConn
	bOpts balancer.BuildOptions
	cfg   atomic.Pointer[customRRConfig]
}

func (crr *customRoundRobin) updateClientConnState(state balancer.ClientConnState) error {
	crrCfg, ok := state.BalancerConfig.(*customRRConfig)
	if !ok {
		return balancer.ErrBadResolverState
	}
	if el := state.ResolverState.Endpoints; len(el) != 2 {
		return fmt.Errorf("UpdateClientConnState wants two endpoints, got: %v", el)
	}
	crr.cfg.Store(crrCfg)
	// 调用 UpdateClientConnState 应始终生成一个新的 Picker。
	// 这是有保证的，因为聚合器在其 UpdateClientConnState 中总是会调用 UpdateChildState。
	return crr.Balancer.UpdateClientConnState(balancer.ClientConnState{
		BalancerConfig: gracefulSwitchPickFirst,
		ResolverState:  state.ResolverState,
	})
}

func (crr *customRoundRobin) UpdateState(state balancer.State) {
	if state.ConnectivityState == connectivity.Ready {
		childStates := endpointsharding.ChildStatesFromPicker(state.Picker)
		var readyPickers []balancer.Picker
		for _, childState := range childStates {
			if childState.State.ConnectivityState == connectivity.Ready {
				readyPickers = append(readyPickers, childState.State.Picker)
			}
		}
		if len(readyPickers) == 2 {
			picker := &customRoundRobinPicker{
				pickers:      readyPickers,
				chooseSecond: crr.cfg.Load().ChooseSecond,
				next:         0,
			}
			crr.ClientConn.UpdateState(balancer.State{
				ConnectivityState: connectivity.Ready,
				Picker:            picker,
			})
			return
		}
	}
}

type customRoundRobinPicker struct {
	pickers      []balancer.Picker
	chooseSecond uint32
	next         uint32
}

func (crrp *customRoundRobinPicker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	next := atomic.AddUint32(&crrp.next, 1)
	index := 0
	if next != 0 && next%crrp.chooseSecond == 0 {
		index = 1
	}
	childPicker := crrp.pickers[index%len(crrp.pickers)]
	return childPicker.Pick(info)
}
