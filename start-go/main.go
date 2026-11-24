//Go 환경 세팅 및 시스템 정보 출력 실습

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
