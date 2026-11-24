package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var s []int

	fmt.Println(unsafe.Sizeof(s), s == nil)
	fmt.Println("nil슬라이스 선언 후 상태 ->", s, len(s), cap(s))
	fmt.Printf("선언 후 헤더 Data값: [%p]\n", s)
	fmt.Println()

	fmt.Println("\n======================== 실습 1 ========================")
	for i := 1; i <= 10; i++ {
		if len(s) < cap(s) {
			fmt.Println("아직 공간 있음, 기존공간 그대로")
		} else {
			fmt.Println("수용공간 부족, 새로운 공간 할당")
		}

		fmt.Printf("before: [%p]\n", s)
		s = append(s, i*10)
		fmt.Printf("after: [%p]\n", s)

		fmt.Printf("[%d회차] -> 슬라이스의 사용가능한 길이: [%d], 슬라이스의 최대 수용길이: [%d]\n", i, len(s), cap(s))
	}

	fmt.Println("\n======================== 실습 2 ========================")

	// make함수를 사용할때 주의점!!
	good := make([]int, 0, 5)
	trap := make([]int, 5)

	fmt.Printf("good -> len: [%d], cap: [%d], addr: [%p]\n", len(good), cap(good), good)
	fmt.Printf("trap -> len: [%d], cap: [%d], addr: [%p]\n\n", len(trap), cap(trap), trap)

	good = append(good, 80, 443, 8080)
	trap = append(trap, 80, 443, 8080)

	fmt.Printf("good -> value: %v, len: [%d], cap: [%d], addr: [%p]\n", good, len(good), cap(good), good)
	fmt.Printf("trap -> value: %v, len: [%d], cap: [%d], addr: [%p]\n", trap, len(trap), cap(trap), trap)
}
