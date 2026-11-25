package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var nilMap map[string]int
	emptyMap := make(map[string]int)

	fmt.Println("기본사항 체크 >>>")
	fmt.Println("[SIZE]: nilMap emptyMap ->", unsafe.Sizeof(nilMap), unsafe.Sizeof(emptyMap))
	fmt.Println("[is_nil]: nilMap emptyMap ->", nil == nilMap, nil == emptyMap)

	fmt.Printf("\n[nilMap Address info]: &nilMap nilMap -> [%p] [%p]\n", &nilMap, nilMap)
	fmt.Printf("[emptyMap Address info]: &emptyMap emptyMap -> [%p] [%p]\n", &emptyMap, emptyMap)

	fmt.Println("\n=========================== 실습 1 ===========================")
	check := nilMap["HTTP"]
	fmt.Println(check)

	fmt.Println("\n=========================== 실습 2 ===========================")
	emptyMap["SSH"] = 22
	fmt.Println(emptyMap)
	fmt.Println(emptyMap["ssh"])
	fmt.Println(emptyMap["SSH"])

	nilMap = make(map[string]int)
	nilMap["HTTP"] = 80
	fmt.Println(nilMap)

	fmt.Println("\n=========================== 실습 3 ===========================")
	services := map[string]int{"HTTP": 80, "HTTPS": 443}
	fmt.Println(services)

	fmt.Println("\nMySQL추가 >>>")
	services["MySQL"] = 3306
	fmt.Println(services)

	fmt.Println("\nHTTP 수정 >>>")
	services["HTTP"] = 8080
	fmt.Println(services)

	fmt.Println("\n현재 상태 map조회 >>>")
	fmt.Printf("map: %v, map(origin): %#v, len: [%d]\n", services, services, len(services))

	fmt.Println("\n없는 Key값 조회 >>>")
	fmt.Println(services["FTP"])
}
