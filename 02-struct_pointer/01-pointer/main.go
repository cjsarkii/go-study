package main

import (
	"fmt"
	"unsafe"
)

func main() {
	timeoutSec := 30 // 그냥 일반 정수형을 담는 변수
	p := &timeoutSec // 정수형을 담는 변수의 주소를 담은 포인터 변수

	fmt.Println(unsafe.Sizeof(timeoutSec), unsafe.Sizeof(p), p)
	// 결과 -> 8 8 0x3eab43cae0b0, timeoutSec자체는 정수를 담는 변수인데 int형이 go에서 8바이트인가?

	// 첫번째 케이스 -> 정상 범위일때
	fmt.Println("======================= 정상범위 =======================")
	ok, reason := applyTimeout(p, 55)
	fmt.Printf("bool: [%v], return: [%s]\n", ok, reason)
	fmt.Printf("p(&timeoutSec): [%p], *p(timeoutSec): [%d], timeoutSec: [%d]\n\n", p, *p, timeoutSec)

	fmt.Println("======================= 오류범위 =======================")
	ok, reason = applyTimeout(p, 0)
	fmt.Printf("bool: [%v], return: [%s]\n", ok, reason)
	fmt.Printf("p(&timeoutSec): [%p], *p(timeoutSec): [%d], timeoutSec: [%d]\n\n", p, *p, timeoutSec)

	fmt.Println("======================= nil테스트 =======================")
	var q *int
	// q = &timeoutSec
	fmt.Printf("q == nil: [%v]\n\n", q == nil)

	ok, reason = applyTimeout(q, 10)
	fmt.Printf("bool: [%v], return: [%s]\n", ok, reason)

	if q != nil {
		fmt.Printf("q: [%p], *q: [%d]\n\n", q, *q)
	} else {
		fmt.Printf("q: [%p] (nil포인터라 역참조 생략)\n\n", q)
	}

	fmt.Println("======================= swapPorts =======================")
	portA := 7777
	portB := 9999

	fmt.Printf("A: [%d], B: [%d]\n", portA, portB)
	fmt.Println("(포인터 변경 X, 값만 변경)swap실행 >>>") // 이거를 역참조 값 변경했다라고 이해하면 되나??

	ok = swapPorts(&portA, &portB)
	if ok {
		fmt.Println("결과 >>>")
		fmt.Printf("A: [%d], B: [%d]\n\n", portA, portB)
	} else {
		fmt.Println("nil값이 있는지 확인하세요")
	}

	fmt.Println("======================= 탈출분석 실험 =======================")
	defaultPort := createDefaultPort()

	fmt.Println("main >>>")
	fmt.Printf("defaultPort: [%p], *defaultPort: [%d]\n", defaultPort, *defaultPort)
}

func applyTimeout(target *int, newSec int) (bool, string) {
	if target == nil {
		return false, "유효하지 않은 포인터입니다."
	}
	if newSec < 1 || newSec > 300 {
		return false, "타임아웃은 1~300초 범위여야 합니다."
	}

	// go에서도 c처럼 포인터가 찌르는 곳의 값(이거를 참조한다고 표현하는건가??)을 다룰때는 *(별표?? 와일드카드??)를 앞에 붙이는구나??
	*target = newSec
	return true, "타임아웃 변경 완료"
}
func swapPorts(a, b *int) bool {
	if a == nil || b == nil {
		return false
	}

	*a, *b = *b, *a
	return true
}

func createDefaultPort() *int {
	port := 80

	// fmt.Println("func >>>")
	// fmt.Printf("&port: [%p], port: [%d]\n", &port, port)
	return &port
}
