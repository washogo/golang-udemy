package main

import (
	"fmt"
	"strconv"
)

//定数
const Pi = 3.14
const (
	URL = "http://xxx.com"
	SiteName = "test"
)
const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

const (
	c0 = iota
	c1
	c2
)

func main() {
	fmt.Println(Pi)

	// Pi = 3
	// fmt.Println(Pi)

	fmt.Println(URL, SiteName)
	fmt.Println(A, B, C, D, E, F)
	fmt.Println(c0, c1, c2)
}

func byteX() {
	byteA := []byte{72, 73}
	fmt.Println(byteA)

	fmt.Println(string(byteA))

	c := []byte("HI")
	fmt.Println(c)

	fmt.Println(string(c))
}

func array() {
	var arr1 [3]int
	fmt.Println(arr1)
	fmt.Printf("%T\n", arr1)

	var arr2 [3]string = [3]string{"A", "B"}
	fmt.Println(arr2)

	arr3 := [3]int{1, 2, 3}
	fmt.Println(arr3)

	arr4 := [...]string{"C", "D"}
	fmt.Println(arr4)
	fmt.Printf("%T\n", arr4)

	fmt.Println(arr2[0])

	arr2[2] = "C"
	fmt.Println(arr2)

	var arr5 [4]int
	// arr5 = arr1
	fmt.Println(arr5)
}

func interfaceX() {
	var x interface{}
	fmt.Println(x)

	x = 1
	fmt.Println(x)
	x = "A"
	fmt.Println(x)
	x = [3]int{1, 2, 3}
	fmt.Println(x)
}

func changeType() {
	// var i int = 1
	// fl64 := float64(i)
	// fmt.Println(fl64)
	// fmt.Printf("i = %T\n", i)
	// fmt.Printf("fl64 = %T\n", fl64)

	// i2 := int(fl64)
	// fmt.Printf("i2 = %T\n", i2)

	// fl := 10.5
	// i3 := int(fl)
	// fmt.Printf("i3 = %T\n", i3)
	// fmt.Println(i3)

	var s string = "100"
	fmt.Printf("s = %T\n", s)

	i, _ := strconv.Atoi("A")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	fmt.Println(i)
	fmt.Printf("%T\n", i)

	var i2 int = 200
	s2 := strconv.Itoa(i2)
	fmt.Println(s2)
	fmt.Printf("s2 = %T\n", s2)

}
