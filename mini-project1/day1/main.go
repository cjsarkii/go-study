// 전역변수 금지, 모든 데이터는 main안에서 선언
// 반드시 함수 매개변수로 주고받으며 구현

package main

import (
	"fmt"
)

func main() {
	inventory := map[string]int{"coffee": 10, "water": 0, "ramen": 5} // HashTable
	logs := make([]string, 0, 10)                                     // heap-memory
	command := make([]string, 7)

	command[0] = "check"
	command[1] = "restock"
	command[2] = "sell"
	command[3] = "remove"
	command[4] = "addlog"
	command[5] = "viewlog"
	command[6] = "exit"

	fmt.Println("================================================================")
	fmt.Println("사용가능한 명령어 >>>")
	fmt.Println(command)
	fmt.Println("================================================================\n")

	_ = inventory
	_ = logs

	var input string

	for {
		fmt.Print("원하는 작업을 입력하세요(exit입력시 종료): ")
		fmt.Scan(&input)

		if input == command[6] {
			fmt.Println("프로그램을 종료합니다.")
			break
		}

		switch input {
		case command[0]:
			fmt.Print("[조회] >>> 물품의 이름을 입력해주세요: ")
			fmt.Scan(&input)
			checkStock(inventory, input)
		case command[1]:

		case command[2]:

		case command[3]:

		case command[4]:

		case command[5]:

		}
	}
}

func checkStock(inv map[string]int, item string) {

}
func restock(inv map[string]int, item string, qty int) bool {

}
