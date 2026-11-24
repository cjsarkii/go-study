package main

import "fmt"

func main() {
	port := [3]int{22, 80, 443}
	fmt.Printf("port[0] address: [%p], port: %v\n", &port[0], port)

	tryModify(port) // c처럼 주소값을 넘기는게 아니라 깊은복사가 이루어짐
	fmt.Printf("port[0] address: [%p], port: %v\n", &port[0], port)

	// 첫번째 방식 - 둘다 받기
	// 근데 Go는 받아놓고 안쓰면 뭐라고 하니까 어딘가 한번은 써야함
	msg, result := classifyPort(8080)
	_ = msg

	if result {
		// _ = msg
		fmt.Println("정상 결과")
	} else {
		fmt.Println("경고")
	}

	// 두번째 방식 - 필요한 것만 받기
	_, result = classifyPort(-55)
	if result {
		fmt.Println("정상 결과")
	} else {
		fmt.Println("경고")
	}

	// 두번째 방식에서 연장 - 필요한 것만 받아서 사용하기
	onlyMsg, _ := classifyPort(9999)
	fmt.Println(onlyMsg)

	msg, result = classifyPort(10101)
	if result {
		fmt.Println("정상 결과:", msg)
	} else {
		fmt.Println("경고:", msg)
	}
}

func classifyPort(port int) (string, bool) {
	// 문자열이나 실수형태는 없다고 가정
	if port < 1 || port > 65535 {
		return "유효하지 않은 포트", false
	} else if port <= 1023 {
		return "Well-known 포트", true
	}
	return "등록/동적 포트", true

}
