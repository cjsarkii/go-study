package main

import "fmt"

func portinfo() {
	var port int

	fmt.Print("포트 번호를 입력하세요: ")
	fmt.Scan(&port)

	if port == 80 || port == 443 {
		fmt.Println("웹(HTTP/HTTPS) 기본 포트입니다!")
		return
	}

	if port >= 1 && port <= 1023 {
		fmt.Println("Well-Known 포트입니다.")
	} else if port >= 1024 && port <= 49151 {
		fmt.Println("Registered 포트입니다.")
	} else if port >= 49152 && port <= 65535 {
		fmt.Println("Dynamic 포트입니다.")
	} else {
		fmt.Println("유효하지 않은 포트 번호입니다.")
	}

}
