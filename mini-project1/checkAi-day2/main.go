// 전역변수 금지, 모든 데이터는 main 안에서 선언
// 비즈니스 로직 함수는 데이터 처리와 결과 반환만 담당(SRP), 화면 출력과 로그 기록은 main이 전담

package main

import (
	"fmt"
)

func main() {
	// [수정 - 주석 교정]
	// Before: logs := make([]string, 0, 10) // heap-memory
	// After:  logs 변수 자체는 24B 슬라이스 헤더이고, make로 미리 확보한 10칸짜리 뒤쪽 배열이 힙에 할당됨
	inventory := map[string]int{"coffee": 10, "water": 0, "ramen": 5} // 8B 포인터 -> 힙 해시 테이블
	logs := make([]string, 0, 10)                                     // 24B 헤더(len=0, cap=10) -> 힙 배열
	backupLog := make([]string, 0, 10)                                // 증분 백업 저장용 슬라이스

	// [수정 - C 스타일 초기화 -> Go 슬라이스 리터럴]
	// Before: command := make([]string, 9) 후 command[0] = ... 9줄 대입
	// After:  리터럴 한 줄로 선언 (메뉴 안내 출력용)
	commands := []string{
		"check", "restock", "sell", "remove",
		"status", "view", "backup", "backupview", "exit",
	}

	// [수정 - go vet 경고 방지]
	// Before: fmt.Println("====...\n") -> Println 자체가 개행을 붙이므로 \n 중복 경고 발생
	// After:  끝의 \n 제거 후 필요하면 빈 fmt.Println() 사용
	fmt.Println("================================================================")
	fmt.Println("사용가능한 명령어 >>>")
	fmt.Println(commands)
	fmt.Println("================================================================")
	fmt.Println()

	// [수정 - 변수 선언 위치 최소화 (C89 습관 교정)]
	// Before: inputOption, inputName, inputCount, possbleLogCount를 모두 여기 상단에 몰아 선언
	// After:  루프를 돌며 값이 계속 누적되어야 하는 logPosition만 for 바깥에 선언
	logPosition := 0

	for {
		var inputOption string
		fmt.Print("\n원하는 작업을 입력하세요(exit 입력 시 종료): ")
		fmt.Scan(&inputOption)

		// [수정 - 매직 넘버 인덱스(command[0]~[8]) 제거]
		// Before: switch inputOption { case command[0]: ... }
		// After:  Go의 switch는 C와 달리 문자열 직접 비교를 지원하므로 명령어 이름을 바로 명시
		switch inputOption {
		case "exit":
			fmt.Println("프로그램을 종료합니다.")
			// [Go 포인트] switch 안에서 break를 쓰면 for가 아니라 switch만 빠져나감!
			// main 함수를 즉시 끝낼 때는 return이 가장 깔끔함
			return

		case "check":
			var inputName string
			fmt.Print("[조회] >>> 물품의 이름을 입력해주세요: ")
			fmt.Scan(&inputName)

			// [수정 - SRP 적용] checkStock이 문자열만 반환하고 출력은 main이 담당
			fmt.Println(checkStock(inventory, inputName))

		case "restock":
			var inputName string
			var inputCount int
			fmt.Print("[입고/신규등록] >>> 물품의 이름과 수량을 입력해주세요: ")
			fmt.Scan(&inputName, &inputCount)

			// [수정 - SRP 통일 & 불필요한 break 제거]
			// Before: restock이 bool만 반환하고 함수 내부에서 fmt.Printf 출력 + if 안에 break 사용
			// After:  sell/removeItem과 동일하게 (bool, string) 반환으로 통일
			ok, reason := restock(inventory, inputName, inputCount)
			if ok {
				msg := fmt.Sprintf("[입고] %s +%d개", inputName, inputCount)
				logs = addLog(logs, msg)
			}
			fmt.Println(reason)

		case "sell":
			var inputName string
			var inputCount int
			fmt.Print("[판매] >>> 판매하실 물품과 수량을 입력해주세요: ")
			fmt.Scan(&inputName, &inputCount)

			// [수정 - 중복 출력 제거]
			// Before: if/else 양쪽 블록 끝에 fmt.Println(reason)을 중복 작성
			// After:  성공 시 로그만 기록하고, 결과 메시지는 if 밖에서 한 번만 출력
			ok, reason := sell(inventory, inputName, inputCount)
			if ok {
				msg := fmt.Sprintf("[판매] %s -%d개", inputName, inputCount)
				logs = addLog(logs, msg)
			}
			fmt.Println(reason)

		case "remove":
			var inputName string
			fmt.Print("[품목 삭제] >>> 물품의 이름을 입력해주세요(기존 수량까지 전부 삭제): ")
			fmt.Scan(&inputName)

			// [수정 - SRP 통일 & 불필요한 Sprintf 방지]
			// Before: 삭제 전에 무조건 msg 생성 + removeItem에 logs/msg까지 넘겨 3개 반환
			// After:  삭제 전 수량만 잠깐 기록해 두고, removeItem은 inventory 삭제만 전담((bool, string))
			oldQty := inventory[inputName]
			ok, reason := removeItem(inventory, inputName)
			if ok {
				msg := fmt.Sprintf("[삭제] %s (기존 수량: %d개)", inputName, oldQty)
				logs = addLog(logs, msg)
			}
			fmt.Println(reason)

		case "status":
			fmt.Println("\n=================================== 재고 현황 ===================================")
			if len(inventory) == 0 {
				fmt.Println("등록된 상품이 없습니다.")
				break
			}
			// [수정] 재고 0인 품목에 [품절] 표시 추가
			for name, qty := range inventory {
				if qty == 0 {
					fmt.Printf("품명: [%s], 수량: [%d] [품절]\n", name, qty)
				} else {
					fmt.Printf("품명: [%s], 수량: [%d]\n", name, qty)
				}
			}

		case "view":
			fmt.Println("\n=================================== 로그 확인 ===================================")
			// [수정] 로그가 비어 있을 때 안내 문구 추가
			if len(logs) == 0 {
				fmt.Println("기록된 로그가 없습니다.")
				break
			}
			for idx, msg := range logs {
				fmt.Printf("[로그 %d] -> %s\n", idx+1, msg)
			}

		case "backup":
			// [수정 - 오타 교정 및 스코프 축소] possbleLogCount -> possibleLogCount (이 case 안에서만 사용)
			possibleLogCount := len(logs) - logPosition
			if possibleLogCount == 0 {
				fmt.Println("현재 백업 가능한 신규 로그가 없습니다.")
				break // 여기서 break는 switch를 빠져나가 다음 for 반복으로 이동
			}

			// [수정 - 안내 문구와 실제 동작 일치]
			// Before: "참고-마지막로그부터 기록합니다."
			// After:  실제 코드는 logPosition 0번부터 미백업 로그를 순서대로 쌓는 증분 백업 방식
			var inputCount int
			fmt.Printf("백업할 로그 개수를 입력해주세요(미백업 %d건, 오래된 미백업 로그부터 순차 기록): ", possibleLogCount)
			fmt.Scan(&inputCount)

			// [수정 - 패닉 버그 방어]
			// Before: inputCount <= 0 검사 누락 -> 음수 입력 시 make([]string, -3)에서 즉시 런타임 패닉!
			// After:  1 미만 입력 차단
			if inputCount <= 0 {
				fmt.Println("[오류] >>> 백업 개수는 1 이상이어야 합니다.")
				break
			}

			if inputCount > possibleLogCount {
				fmt.Printf("[알림] -> 원본 로그 [%d개], 백업 완료 [%d개], 남은 [%d개]만 백업합니다.\n",
					len(logs), logPosition, possibleLogCount)
				inputCount = possibleLogCount
			}

			// 예시(원본 로그가 총 6개일 때):
			// 1회차: inputCount=2, logPosition=0 -> logs[0:2] 복사 (2개) -> logPosition=2
			// 2회차: inputCount=10 입력 -> 남은 4개로 조정, logPosition=2 -> logs[2:6] 복사 (4개) -> logPosition=6
			chunk := backupRecentLogs(logs, inputCount, logPosition)
			backupLog = append(backupLog, chunk...)
			logPosition += inputCount

			fmt.Printf("[백업완료] >>> %d건 백업 완료 (누적 백업: %d건)\n", inputCount, len(backupLog))

		case "backupview":
			fmt.Println("\n=================================== 백업로그 확인 ===================================")
			// [수정] 백업 로그가 비어 있을 때 안내 문구 추가
			if len(backupLog) == 0 {
				fmt.Println("백업된 로그가 없습니다.")
				break
			}
			for idx, msg := range backupLog {
				fmt.Printf("[백업로그 %d] -> %s\n", idx+1, msg)
			}

		default:
			fmt.Println("잘못된 명령어입니다. 사용 가능한 명령어:", commands)
		}
	}
}

// [수정 - SRP 적용] 화면 출력 대신 결과 문자열 반환
func checkStock(inv map[string]int, item string) string {
	count, ok := inv[item]
	if !ok {
		return fmt.Sprintf("[%s] >>> 등록되지 않은 제품입니다.", item)
	}
	if count == 0 {
		return fmt.Sprintf("[%s] >>> 품절입니다.", item)
	}
	return fmt.Sprintf("[%s] 현재 재고: %d개", item, count)
}

// [수정 - 시그니처 통일] bool -> (bool, string), 함수 내부 fmt.Printf 제거
func restock(inv map[string]int, item string, qty int) (bool, string) {
	if qty <= 0 {
		return false, "[입고실패] >>> 수량은 1 이상이어야 합니다."
	}

	if _, ok := inv[item]; ok {
		inv[item] += qty
		return true, fmt.Sprintf("[입고완료] >>> 기존 제품 수량 추가: [%s] +%d개 (현재: %d개)", item, qty, inv[item])
	}

	inv[item] = qty
	return true, fmt.Sprintf("[등록완료] >>> 신규 제품 등록: [%s] %d개", item, qty)
}

// 슬라이스 append는 Len(또는 힙 주소)이 바뀌므로 반드시 바뀐 24B 헤더를 return 해야 함
func addLog(logs []string, msg string) []string {
	return append(logs, msg)
}

func sell(inv map[string]int, item string, qty int) (bool, string) {
	// [수정 - 입력값 검증을 최우선으로(Fail-Fast)] 수량이 이상하면 맵 조회조차 할 필요 없음
	if qty <= 0 {
		return false, "[판매실패] >>> 판매 수량은 1 이상이어야 합니다."
	}

	// [수정 - 맵 중복 조회 제거]
	// Before: if _, ok := inv[item]; !ok { ... } 이후 다시 if qty > inv[item] 조회
	// After:  stock, ok := inv[item]으로 해시 조회 1회만 수행
	stock, ok := inv[item]
	if !ok {
		return false, "[판매실패] >>> 존재하지 않는 상품입니다."
	}
	if qty > stock {
		return false, fmt.Sprintf("[판매실패] >>> 재고가 부족합니다. (현재 재고: %d개)", stock)
	}

	inv[item] -= qty
	return true, fmt.Sprintf("[판매완료] >>> [%s] %d개 판매 (남은 재고: %d개)", item, qty, inv[item])
}

// [수정 - SRP 통일]
// Before: removeItem(inv, item, logs, msg) (bool, string, []string) -> 삭제 함수가 로그 슬라이스까지 떠안음
// After:  inventory 삭제만 책임지고 (bool, string) 반환 (restock, sell과 완벽히 동일한 패턴)
func removeItem(inv map[string]int, item string) (bool, string) {
	if _, ok := inv[item]; !ok {
		return false, "[삭제실패] >>> 존재하지 않는 상품입니다."
	}

	delete(inv, item)
	return true, fmt.Sprintf("[삭제완료] >>> [%s] 품목이 삭제되었습니다.", item)
}

func backupRecentLogs(logs []string, n int, logPosition int) []string {
	// [수정 - 함수 자체 경계 방어] 호출부가 실수하더라도 패닉이 나지 않도록 방어
	if n <= 0 || logPosition < 0 || logPosition >= len(logs) {
		return nil
	}
	end := logPosition + n
	if end > len(logs) {
		end = len(logs)
	}

	// [수정 - 변수명 오타 교정 & 슬라이싱 주석 교정]
	// Before: cashLog := make([]string, n)
	//         // 두번째 루프(n=4, logPosition=2) -> logs[2:4] -> logs[2], logs[3] 까지 복사... (오류!)
	// After:  n=4, logPosition=2 일 때 end = 2 + 4 = 6 이므로
	//         logs[2:6] -> logs[2], logs[3], logs[4], logs[5] 총 4개가 복사됨!
	chunk := make([]string, end-logPosition)
	copy(chunk, logs[logPosition:end])
	return chunk
}
