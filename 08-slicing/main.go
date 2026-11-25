package main

import "fmt"

func main() {
	origin := []int{10, 20, 30, 40, 50}
	sub := origin[1:3]

	fmt.Println("======================== 실습 1 ========================")
	fmt.Println("origin[]:", origin)
	fmt.Println("sub[]:", sub)
	fmt.Println()

	fmt.Printf("[origin] >>> len: [%d], cap: [%d], &origin[1]: [%p]\n[sub] >>> len: [%d], cap: [%d], &sub[0]: [%p]\n",
		len(origin), cap(origin), &origin[1], len(sub), cap(sub), sub)

	fmt.Println("\n======================== 실습 2 ========================")
	fmt.Println("sub = append(sub, 777) 실행 >>>")
	sub = append(sub, 777)
	fmt.Println("origin[]:", origin)
	fmt.Println("sub[]:", sub)
	fmt.Printf("\n[origin] >>> len: [%d], cap: [%d], &origin[1]: [%p]\n[sub] >>> len: [%d], cap: [%d], &sub[0]: [%p]\n",
		len(origin), cap(origin), &origin[1], len(sub), cap(sub), sub)

	fmt.Println("\n======================== 실습 3 ========================")
	fmt.Println("sub = append(sub, 888, 999) 실행(자신이 가지고 있는 용량을 벗어나게 유도) >>>")
	sub = append(sub, 888, 999)
	fmt.Println("origin[]:", origin)
	fmt.Println("sub[]:", sub)
	fmt.Printf("\n[origin] >>> len: [%d], cap: [%d], &origin[1]: [%p]\n[sub] >>> len: [%d], cap: [%d], &sub[0]: [%p]\n",
		len(origin), cap(origin), &origin[1], len(sub), cap(sub), sub)

	fmt.Println("\n======================== 실습 4 ========================")
	fmt.Println("이 상태에서 sub[0]값 변경해보기 >>>\n")
	sub[0] = 1111

	fmt.Println("origin[]:", origin)
	fmt.Println("sub[]:", sub)

	fmt.Println("\n[결론] >>>\n원본에서 떼어져 나온 슬라이스는 깊은복사가 아니다.\n따라서 남아있는 뒤쪽 공간으로 확장이 가능할 때는 원본값을 건드리고\n뒤쪽 공간 부족으로 이사가야 할때는 깊은 복사를 하게 된다.\n이렇게 깊은 복사가 이루어지면 더이상 원본값을 건드리지 않게 된다.")

	fmt.Println("\n======================== 실습 5 ========================")

	base := []int{80, 443, 8080, 3306}

	// make의 2번째 인자인 len은 이미 cap사이즈만큼 확보된 공간에서 어디까지 허용할 것인가를 결정해준다.
	cloned := make([]int, 2)
	test := make([]int, 0, 2)
	fmt.Println(base, cloned, test)

	copy(cloned, base[1:3])
	copy(test, base[1:3])

	fmt.Println(base, cloned, test)

	fmt.Println("\n각 변수들 주소값 검증 >>>")
	fmt.Printf("base[1]: [%p], cloned: [%p], test: [%p]\n", &base[1], cloned, test)
}
