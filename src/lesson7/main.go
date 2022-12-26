package main

import "fmt"

// 関数
// 無名関数
// 関数を関数
// 関数を引数にとる関数
// クロージャー

func Plus(x, y int) int {
	return x + y
}

func Div(x, y int) (int, int) {
	q := x / y
	r := x % y
	return q, r
}

func ReturnFunc() func() {
	return func() {
		fmt.Println("I`m a function")
	}
}

func Double(price int) (result int) {
	result = price * 2
	return
}

func CallFunction(f func()) {
	f()
}

func Later() func(string) string {
	var store string
	return func(next string) string {
		s := store
		store = next
		return s
	}
}

// ジェネレーターの実装
func integers() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func main() {
	i := Plus(1, 2)
	fmt.Println(i)

	i2, _ := Div(9, 4)
	fmt.Println(i2)

	i4 := Double(1000)
	fmt.Println(i4)

	f := func(x, y int) int {
		return x + y
	}
	i5 := f(1, 2)
	fmt.Println(i5)

	i6 := func(x, y int) int {
		return x + y
	}(1, 2)

	fmt.Println(i6)

	f2 := ReturnFunc()
	f2()

	CallFunction(func() {
		fmt.Println("function")
	})

	f3 := Later()
	fmt.Println(f3("Hey"))

	ints := integers()
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())
	fmt.Println(ints())

}
