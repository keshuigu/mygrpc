package main

import "fmt"

func get() (int32, int32) {
	return 42, 24
}

func main() {
	if n1, n2 := get(); n1 != 42 {
		fmt.Println(n2)
	} else if n1+n2 == 66 {
		fmt.Println(n1)
	}
}
