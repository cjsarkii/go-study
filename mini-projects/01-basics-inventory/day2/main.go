// 전역변수 금지, 모든 데이터는 main안에서 선언
// 반드시 함수 매개변수로 주고받으며 구현

package main

import (
	"fmt"
)

func main() {
	inventory := map[string]int{"coffee": 10, "water": 0, "ramen": 5} // HashTable
	logs := make([]string, 0, 10)                                     // heap-memory
	command := make([]string, 9)
	backupLog := make([]string, 0, 10)

	command[0] = "check"
	command[1] = "restock"
	command[2] = "sell"
	command[3] = "remove"
	command[4] = "status"
	command[5] = "view"
	command[6] = "backup"
	command[7] = "backupview"
	command[8] = "exit"

	fmt.Println("================================================================")
	fmt.Println("사용가능한 명령어 >>>")
	fmt.Println(command)
	fmt.Println("================================================================\n")

	var inputOption string
	var inputName string
	var inputCount int
	var possbleLogCount int
	var logPosition int

	for {
		fmt.Print("원하는 작업을 입력하세요(exit입력시 종료): ")
		fmt.Scan(&inputOption)

		if inputOption == command[8] {
			fmt.Println("프로그램을 종료합니다.")
			break
		}

		switch inputOption {
		case command[0]:
			fmt.Print("[조회] >>> 물품의 이름을 입력해주세요: ")
			fmt.Scan(&inputName)
			checkStock(inventory, inputName)
		case command[1]:
			fmt.Print("[입고/신규등록] >>> 물품의 이름과 수량을 입력해주세요: ")
			fmt.Scan(&inputName, &inputCount)
			if restock(inventory, inputName, inputCount) {
				msg := fmt.Sprintf("[입고] %s +%d개", inputName, inputCount)
				logs = addLog(logs, msg)
				break
			}
			fmt.Println("[오류] >>> 유효하지 않은 개수 입력됨")
		case command[2]:
			fmt.Print("[판매] >>> 판매하실 물품과 수량을 입력해주세요: ")
			fmt.Scan(&inputName, &inputCount)
			if ok, reason := sell(inventory, inputName, inputCount); ok {
				msg := fmt.Sprintf("[판매] %s, %d개", inputName, inputCount)
				logs = addLog(logs, msg)
				fmt.Println(reason)
			} else {
				fmt.Println(reason)
			}
		case command[3]:
			fmt.Print("[품목 삭제] >>> 물품의 이름을 입력해주세요(참고-기존 수량까지 전부 삭제됩니다.): ")
			fmt.Scan(&inputName)
			msg := fmt.Sprintf("[삭제] %s(기존 수량: %d개)", inputName, inventory[inputName])
			if ok, reason, rcvlog := removeItem(inventory, inputName, logs, msg); ok {
				logs = rcvlog
				fmt.Println(reason)
			} else {
				fmt.Println(reason)
			}
		case command[4]:
			fmt.Println("\n=================================== 재고 현황 ===================================")
			for name, value := range inventory {
				fmt.Printf("품명: [%s], 수량: [%d]\n", name, value)
			}
		case command[5]:
			fmt.Println("\n=================================== 로그 확인 ===================================")
			for idx, msg := range logs {
				fmt.Printf("[로그 %d] -> %s\n", idx+1, msg)
			}
		case command[6]:
			if len(logs)-len(backupLog) == 0 {
				fmt.Println("현재 백업가능한 로그가 없습니다.")
				break
			}

			// 로그 원본이 존재할 때만 진행
			fmt.Print("백업할 로그 개수를 입력해주세요(참고-마지막로그부터 기록합니다.): ")
			fmt.Scan(&inputCount)

			// 백업로그 개수 확인??
			// 몇개를 백업할지 결정(입력한 숫자 > 원본로그 개수일 경우 대비)
			possbleLogCount = len(logs) - len(backupLog)
			if inputCount > possbleLogCount {
				fmt.Printf("[알림] -> 현재 원본로그: [%d개], 기존 로그: [%d개], 가능한 로그수[%d개]를 백업합니다.\n", len(logs), len(backupLog), possbleLogCount)
				inputCount = possbleLogCount
			}

			// 첫번째 루프 -> inputCount: 2, logPosition: 0 -> 2개 로그 저장
			// 두번째 루프 -> inputCount: 10 -> 위에서 가능한 만큼인 4로 변경 -> logPosition: 2
			backupLog = append(backupLog, backupRecentLogs(logs, inputCount, logPosition)...)

			// 마지막 위치 설정 후 탈출
			logPosition += inputCount // logPosition = 0 + 2 = 2
		case command[7]:
			fmt.Println("\n=================================== 백업로그 확인 ===================================")
			for idx, msg := range backupLog {
				fmt.Printf("[백업로그 %d] -> %s\n", idx+1, msg)
			}
		default:
			fmt.Println("잘못된 명령어 입니다.")
		}
	}
}

func checkStock(inv map[string]int, item string) {
	if count, ok := inv[item]; ok {
		if count == 0 {
			fmt.Printf("[%s] >>> 품절입니다.\n", item)
			return
		}
		fmt.Printf("[%s]개수: %d\n", item, count)
		return
	}
	fmt.Println("제품이 없습니다.\n")
}

func restock(inv map[string]int, item string, qty int) bool {
	if qty <= 0 {
		return false
	}

	if _, ok := inv[item]; ok {
		fmt.Printf("기존제품 수량추가 -> [%s], [%d]개\n", item, qty)
		inv[item] += qty
		return true
	}

	fmt.Printf("신규제품 재고등록 완료 -> [%s], [%d]개\n", item, qty)
	inv[item] = qty
	return true
}

func addLog(logs []string, msg string) []string {
	logs = append(logs, msg)
	return logs
}

func sell(inv map[string]int, item string, qty int) (bool, string) {
	// 1. 입력받은 상품이 존재하지 않는 경우
	if _, ok := inv[item]; !ok {
		return false, "[판매실패] >>> 존재하지 않는 상품입니다."
	}

	// 2. 존재하지만 (판매수량 > 재고수량)인 경우
	if qty > inv[item] {
		return false, "재고가 부족합니다."
	}

	// 3. 존재하지만 (판매수량 <= 0)인 경우
	if qty <= 0 {
		return false, "판매수량은 1이상이어야 합니다."
	}

	inv[item] -= qty
	return true, "[판매완료] >>> 정상적으로 처리되었습니다."
}
func removeItem(inv map[string]int, item string, logs []string, msg string) (bool, string, []string) {
	// 삭제를 입력받은 품목이 map에 없을 경우
	if _, ok := inv[item]; !ok {
		return false, "[삭제실패] >>> 존재하지 않는 상품입니다.", logs
	}

	logs = addLog(logs, msg)
	delete(inv, item)
	return true, "[삭제완료] >>> 정상적으로 처리되었습니다.", logs
}

func backupRecentLogs(logs []string, n int, logPosition int) []string {
	cashLog := make([]string, n) // 첫번째 루프때 n = 2, logPosition: 0
	// 두번째 루프 -> n = 4, logPosition: 2
	copy(cashLog, logs[logPosition:logPosition+n]) // logs[0:2] -> logs[0], logs[1] 까지 복사
	// logs[2:4] -> logs[2], logs[3] 까지 복사...
	return cashLog
}
