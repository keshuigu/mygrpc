// Package data provides convenience routines to access files in the data
// directory.
package main

import (
	"path/filepath"
	"runtime"
)

var basepath string

func init() {
	_, currentFile, _, _ := runtime.Caller(0) // 0 代表从Caller的调用者
	basepath = filepath.Dir(currentFile)
}

func Path(rel string) string {
	if filepath.IsAbs(rel) {
		return rel
	}
	return filepath.Join(basepath, rel)
}
