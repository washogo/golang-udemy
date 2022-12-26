package main

import "fmt"

// テスト

func IsOne(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	fmt.Println(IsOne(1))
	fmt.Println(IsOne(0))
}
