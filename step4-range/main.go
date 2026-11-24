package main

import "fmt"

func main() {
	origin := [3]int{10, 20, 30}
	copied := origin

	copied[0] = 999

	fmt.Println(origin)
	fmt.Println(copied)

	fmt.Println("Test >>> 루프안에서 변수값 변경")

	for _, v := range origin {
		v += 100
		fmt.Println(v)
	}

	fmt.Println(origin)

	for i := range origin {
		origin[i] += 100
	}

	fmt.Println(origin)

	for i, _ := range origin {
		origin[i] += 100
	}
	fmt.Println(origin)
}
