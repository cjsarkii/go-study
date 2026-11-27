// 구조체 선언과 배치 순서에 따른 메모리 차지비율 차이
// C와 거의 동일한 방식으로 작용
package main

import (
	"fmt"
	"unsafe"
)

type GoodSession struct {
	Port   uint16
	Active bool
	IP     string
	UserID int64
}

type BadSession struct {
	Active bool
	UserID int64
	Port   uint16
	IP     string
}

func main() {
	var good GoodSession
	var bad BadSession

	fmt.Println("Good:", unsafe.Sizeof(good))
	fmt.Println("Bad:", unsafe.Sizeof(bad))

	fmt.Println("--- GoodSession Offsets ---")
	fmt.Println("Port:", unsafe.Offsetof(good.Port))
	fmt.Println("Active:", unsafe.Offsetof(good.Active))
	fmt.Println("IP:", unsafe.Offsetof(good.IP))
	fmt.Println("UserID:", unsafe.Offsetof(good.UserID))

	fmt.Println("--- BadSession Offsets ---")
	fmt.Println("Active:", unsafe.Offsetof(bad.Active))
	fmt.Println("UserID:", unsafe.Offsetof(bad.UserID))
	fmt.Println("Port:", unsafe.Offsetof(bad.Port))
	fmt.Println("IP:", unsafe.Offsetof(bad.IP))

}
