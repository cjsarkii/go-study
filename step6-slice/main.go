package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// nil슬라이스와 빈 슬라이스의 비교 - 둘다 실체는 존재함
	var nilSlice []int    // int *p = NULL
	emptySlice := []int{} // int *p = runtime.zerobase라는 빈 더미주소
	testSlice := []int{}

	fmt.Println("======================== 실습 1 ========================")
	fmt.Printf("nil길이: [%d], empty길이: [%d]\n", len(nilSlice), len(emptySlice))
	fmt.Printf("nil수용크기: [%d], empty수용크기: [%d]\n", cap(nilSlice), cap(emptySlice))

	fmt.Println(nilSlice == nil)
	fmt.Println(emptySlice == nil)

	fmt.Println(unsafe.Sizeof(nilSlice))
	fmt.Println(unsafe.Sizeof(emptySlice))

	fmt.Printf("%%v -> nil: %v, empty: %v\n", nilSlice, emptySlice)
	fmt.Printf("%%#v -> nil: %#v, empty: %#v\n", nilSlice, emptySlice)
	fmt.Printf("Address -> nil: [%p], empty: [%p], test: [%p]\n", nilSlice, emptySlice, testSlice)

	fmt.Println("\n======================== 실습 2 ========================")

	nums := []int{10, 20, 30}

	fmt.Printf("Addr >>>\nnums(slice자체): [%p]\nhdr(slice내부 헤더): [%p]\nindex[0]: [%p]\n", nums, &nums, &nums[0])

	fmt.Println("\n======================== 실습 3 ========================")
	buf := make([]int, 3, 5)

	fmt.Println(buf, len(buf), cap(buf))
	buf[0] = 80
	buf[1] = 443

	fmt.Println(buf)
	buf[2] = 8080

	// 공간이 뒤쪽에 남아있더라도 make로 3칸만 할당했기 때문에 나머지 2개는 무시?
	fmt.Println(buf)

	// 쓰겠다고 선언하지않은 곳에 값을 대입하고 출력 -> panic에러
	buf[3] = 3306
	fmt.Println(buf)
}
