package main

import "fmt"

func main() {
	nums := []int{10, 20, 30}

	fmt.Printf("nums의 헤더 주소: [%p]\nnums의 값(내부배열 주소): [%p]\n", &nums, nums)
	fmt.Println("nums[] ->", nums)

	fmt.Println("\n============================== 실습 1 ==============================")
	modifyElement(nums)
	fmt.Println("(메인에서 실행) >>> 함수실행 이후 ->", nums)

	fmt.Println("\n============================== 실습 2 ==============================")
	full := []int{10, 20}
	spare := make([]int, 2, 4)

	fmt.Printf("[생성] >>> full: %#v spare: %#v\n", full, spare)
	fmt.Printf("full -> 값: %v, len: [%d], cap: [%d], Data: [%p]\n", full, len(full), cap(full), full)
	fmt.Printf("spare -> 값: %v, len: [%d], cap: [%d], Data: [%p]\n", spare, len(spare), cap(spare), spare)

	appendFail(full)
	appendFail(spare)

	fmt.Println("======================= [main내부 결과] =======================")
	fmt.Println("full:", full)
	fmt.Println("spare:", spare)
	fmt.Printf("\nfull -> Data: [%p]\nspare -> Data: [%p]\n", full, spare)
	fmt.Println("안보이는 값 확인 >>>", spare[:cap(spare)])

	// append를 통해서 heap에 열려있는 공간을 넘어가는 데이터를 추가 하게 되면 새로 동적할당을 진행하게 되는데
	// 이때 새로 열려진 heap영역주소를 못받오니까 안보이는거구나?? 반대로 공간이 충분하면 기존에 열려있는 공간에
	// 추가할수가 있는거고. 그렇다면 free는 어떻게 하는거지??
	// 정상적으로 공간을 열고 그 주소를 받았다면 main에서 처음에 열려있는 공간은 누가 free해주는거지??
	// 반대로 공간을 열긴했지만 주소를 못 받아와서 미아가 됐다면 그거는 또 누가 free하는거지??

	fmt.Println("\n============================== 실습 3 ==============================")
	full = appendSuccess(full, 777)
	fmt.Println("======================= [main내부 결과] =======================")
	fmt.Println("full:", full)
	fmt.Printf("full -> len: [%d], cap: [%d], Data: [%p]\n", len(full), cap(full), full)
}

func modifyElement(s []int) {
	fmt.Println("넘겨받은 매개변수의 주소값 >>>")
	fmt.Printf("&s: [%p]\ns: [%p]\n", &s, s)

	s[0] = 999
	fmt.Println("\ns[0]=999 실행 >>>\n")
	fmt.Printf("&s: [%p]\ns: [%p]\n", &s, s)
	fmt.Println("값:", s)
}

func appendFail(s []int) {
	s = append(s, 777)
	fmt.Println("\n======================= [appendFail 내부결과] =======================")
	fmt.Printf("헤더주소: [%p]\n", &s)
	fmt.Printf("s: %v, len(s): [%d], cap(s): [%d], Data: [%p]\n\n", s, len(s), cap(s), s)
}

func appendSuccess(s []int, val int) []int {
	s = append(s, val)
	fmt.Println("\n======================= [appendSuccess 내부결과] =======================")
	fmt.Printf("헤더주소: [%p]\n", &s)
	fmt.Printf("s: %v, len(s): [%d], cap(s): [%d], Data: [%p]\n\n", s, len(s), cap(s), s)
	return s
}
