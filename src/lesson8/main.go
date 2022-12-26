package main

import (
	"fmt"
)

// switch

// 式スイッチ

/*
func main() {
	n := 3
	switch n {
	case 1, 2:
		fmt.Println("1 or 2")
	case 3, 4:
		fmt.Println("3 or 4")
	default:
		fmt.Println("default")
	}

	switch n := 2; n {
	case 1, 2:
			fmt.Println("1 or 2")
		case 3, 4:
			fmt.Println("3 or 4")
		default:
			fmt.Println("default")
	}

	n := 6
	switch {
	case n > 0 && n < 4:
		fmt.Println("0 < n < 4")
	case n > 3 && n < 7:
		fmt.Println("3 < n < 7")
	}
}
*/

// 型スイッチ
/*
func anything(a interface{}) {
	// fmt.Println(a)
	switch v := a.(type) {
	case string:
		fmt.Println(v + "!?")
	case int:
		fmt.Println(v + 10000)
	}
}

func main() {
	anything("aaa")
	anything(1)

	var x interface{} = 3

	i, isInt := x.(int)

	fmt.Println(i, isInt)
	// fmt.Println(x + 2)

	f, isFloat64 := x.(float64)
	fmt.Println(f, isFloat64)

	if x == nil {
		fmt.Println("none")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, isString)
	} else {
		fmt.Println("I don`t know")
	}

	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don`t know")
	}

	switch v := x.(type) {
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int")
	default:
		fmt.Println("I don`t know")
	}

}
*/

// ラベル付きfor

/*
func main() {
Loop:
	for {
		for {
			for {
				fmt.Println("START")
				break Loop
			}
			fmt.Println("処理をしない")
		}
		fmt.Println("処理をしない")
	}
	fmt.Println("END")

Loop:
	for i := 0; i < 3; i++ {
		for j := 1; j < 3; j++ {
			if j > 1 {
				continue Loop
			}
			fmt.Println(i, j, i*j)
		}
		fmt.Println("処理をしない")
	}
}
*/

// defer
/*
func TestDefer() {
	defer fmt.Println("END")
	fmt.Println("START")
}

func RunDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}
*/
/*
func main() {
	TestDefer()
		defer func() {
			fmt.Println("1")
			fmt.Println("2")
			fmt.Println("3")
		}()
	RunDefer()

	file, err := os.Create("test.txt")
	if  err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file.Write([]byte("Hello"))
}
*/

// panic recover
/*
func main() {
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()
	panic("runtime error")
	fmt.Println("Start")
}
*/

// go goroutin
/*
func sub() {
	for {
		fmt.Println("Sub Loop")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go sub()
	go sub()

	for {
		fmt.Println("Main Loop")
		time.Sleep(200 * time.Millisecond)
	}

}
*/

// init

func init() {
	fmt.Println("init")
}

func init() {
	fmt.Println("init2")
}

func main() {
	fmt.Println("Main")
}
