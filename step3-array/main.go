package main

import "fmt"

func main() {
	var array [3]int
	allowedPorts := [...]int{22, 80, 443, 3306, 8080}

	fmt.Println(array)
	fmt.Println(allowedPorts)
	fmt.Printf("array: %d, allowedPorts: %d\n", len(array), len(allowedPorts))

	for i := 0; i < len(allowedPorts); i++ {
		fmt.Printf("인덱스: %d -> 포트: %d\n", i, allowedPorts[i])
	}
	fmt.Println()
	for i, v := range allowedPorts {
		fmt.Printf("인덱스: %d -> 포트: %d\n", i, v)
	}
	fmt.Println()
	for _, v := range allowedPorts {
		fmt.Println(v)
	}

	for i, v := range allowedPorts {
		fmt.Printf("[%d번] %d\n", i, v)
	}

	fmt.Print("번호 하나를 입력하세요: ")
	var num int
	fmt.Scan(&num)

	isAllowed := false
	for _, v := range allowedPorts {
		if num == v {
			isAllowed = true
			fmt.Println("[시스템]: 발견!")
			break
		} else {
			fmt.Println("[시스템]: 다음확인")
			continue
		}
	}
	fmt.Println("\n============================ 최종결과 ============================")
	if !isAllowed {
		fmt.Println("접속차단")
	} else {
		fmt.Println("접속허용")
	}
}
