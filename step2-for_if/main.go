package main

import "fmt"

func main() {
	var port int

	for {
		fmt.Print("포트 번호를 입력하세요 (0입력시 탈출): ")
		fmt.Scan(&port)

		if port == 0 {
			fmt.Print("관제 종료")
			break
		} else if port < 0 || port > 65535 {
			fmt.Println("유효하지 않은 포트입니다.")
		} else if port <= 1023 {
			fmt.Println("Well-known 포트")
		} else {
			fmt.Println("등록/동적 포트")
		}
	}
}
