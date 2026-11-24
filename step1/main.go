package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("hello")
	fmt.Println(runtime.Version(), runtime.GOOS, runtime.GOARCH, runtime.NumCPU())
	fmt.Printf("version: %s\nOS: %s\nArch: %s\nCPU: %d\n", runtime.Version(), runtime.GOOS, runtime.GOARCH, runtime.NumCPU())
}
