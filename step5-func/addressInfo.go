package main

import "fmt"

func tryModify(arr [3]int) {
	arr[0] = 999
	fmt.Printf("arr[0] address: [%p], arr: %v\n", &arr[0], arr)
}
