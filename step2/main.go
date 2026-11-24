package main

import "fmt"

func main() {
	var score int
	var name string
	var age int
	var height float64

	fmt.Print("당신의 이름/점수/나이/키를 입력하세요: ")
	fmt.Scan(&name, &score, &age, &height)

	fmt.Println("입력하신 정보가 맞나요??")
	fmt.Printf("이름: [%s], 점수: [%d], 나이: [%d], 키: [%.1f]\n\n", name, score, age, height)

	portinfo()
}
