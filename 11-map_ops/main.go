package main

import (
	"fmt"
)

func main() {
	var val string
	// isExist := false
	blockCount := map[string]int{"192.168.1.10": 0, "192.168.1.20": 5}

	fmt.Println("============================= 실습 1 =============================")
	for {
		fmt.Print("조회하려는 IP 끝자리를 입력하세요: ")
		fmt.Scan(&val)

		// 음.. Go는 C에서처럼 캐스팅같은 문법이 없나??

		val = "192.168.1." + val
		// fmt.Printf("[%s] 조회를 시작합니다.\n", val)

		if count, ok := blockCount[val]; !ok {
			fmt.Println("조회완료 >>> 해당하는 IP가 없습니다.\n")
			continue
		} else {
			fmt.Printf("조회완료 >>> [%s]의 차단횟수: [%d]\n", val, count)
			break
		}
	}

	fmt.Println("\n============================= 실습 2 =============================")
	var addIp int
	fmt.Print("IP를 몇개 추가 하실건가요??")
	fmt.Scan(&addIp)

	// 너무 C스럽나..?? Go언어답게 for range addIp 같은거로 해보고싶은데 마땅히 방법이 생각이 안나네
	for i := 0; i < addIp; i++ {
		fmt.Print("\n추가하실 IP 끝자리를 입력하세요: ")
		fmt.Scan(&val)
		if val == "0" {
			fmt.Println("[경고] >>> 0입력 불가")
			i--
			continue
		}

		val = "192.168.1." + val
		if _, ok := blockCount[val]; ok {
			fmt.Println("존재하는 IP ->", val)
			i--
		} else {
			blockCount[val] = 0
			fmt.Println("IP 추가완료 ->", val)
		}
	}

	fmt.Println("\n저장완료 >>> blockCount[] ->", blockCount)

	fmt.Println("\n============================= 실습 3 =============================")
	for {
		fmt.Print("삭제할 IP를 입력하세요(0입력시 종료): ")
		fmt.Scan(&val)
		if val == "0" {
			fmt.Println("실습3을 종료합니다.")
			break
		}

		val = "192.168.1." + val
		delete(blockCount, val)
		fmt.Println("\n결과 >>>", blockCount)
	}

	fmt.Println("\n============================= 실습 4 =============================")
	var repeat int
	fmt.Print("blockCount 출력을 몇번 반복할까요? ")
	fmt.Scan(&repeat)

	for i := range repeat {
		fmt.Printf("[%d]회차 >>>\n", i+1)
		for ip := range blockCount {
			fmt.Println(ip)
		}
		fmt.Println()
	}

	fmt.Println("\n============================= 실습 5(IP추가 함수) =============================")
	var cnt int
	fmt.Printf("blockCount addr in main -> &blockCount: [%p], HashTable: [%p]\n\n", &blockCount, blockCount)
	fmt.Print("추가할 IP와 차단횟수 입력(우선 겹치지 않는 IP로 진행): ")
	fmt.Scan(&val, &cnt)

	val = "192.168.1." + val

	addRule(blockCount, val, cnt)

	fmt.Println("\n============================= 실습 6(IP제거 함수) =============================")
	fmt.Print("제거할 IP 입력(우선 겹치지 않는 IP로 진행): ")
	fmt.Scan(&val)

	val = "192.168.1." + val
	removeRule(blockCount, val)
}

func addRule(m map[string]int, ip string, count int) {
	fmt.Println("=========================== addRule 내부 ===============================")
	fmt.Printf("Params addr in addRule -> &m: [%p], m: [%p]\n\n", &m, m)
	m[ip] = count

	fmt.Println("현재 상태(blockCount) ->", m)
}

func removeRule(m map[string]int, ip string) {
	fmt.Println("=========================== removeRule 내부 ===============================")
	fmt.Printf("Params addr in removeRule -> &m: [%p], m: [%p]\n\n", &m, m)

	delete(m, ip)
	fmt.Println("현재 상태(blockCount) ->", m)
}
