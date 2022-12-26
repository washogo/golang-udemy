package main

import (
	"fmt"
)

// スライス

// 宣言、操作
/*
func main() {
	var sl []int
	fmt.Println(sl)

	var sl2 []int = []int{100, 200}
	fmt.Println(sl2)

	sl3 := []string{"A", "B"}
	fmt.Println(sl3)

	sl4 := make([]int, 5)
	fmt.Println(sl4)

	sl2[0] = 1000
	fmt.Println(sl2)

	sl5 := []int{1, 2, 3, 4, 5}
	fmt.Println(sl5)

	fmt.Println(sl5[0])

	fmt.Println(sl5[2:4])

	fmt.Println(sl5[:2])

	fmt.Println(sl5[2:])

	fmt.Println(sl5[:])

	fmt.Println(sl5[cap(sl5)-1])

	fmt.Println(sl5[1 : len(sl5)-1])
}
*/

// append make len cap
/*
func main() {
	sl := []int{100, 200}
	fmt.Println(sl)

	sl = append(sl, 300)
	fmt.Println(sl)

	sl = append(sl, 400, 500, 600)
	fmt.Println(sl)

	sl2 := make([]int, 5)
	fmt.Println(sl2)

	fmt.Println(len(sl2))

	fmt.Println(cap(sl2))

	sl3 := make([]int, 5, 10)

	fmt.Println(len(sl3))

	fmt.Println(cap(sl3))

	sl3 = append(sl3, 1, 2, 3, 4, 5, 6, 7)

	fmt.Println(len(sl3))

	fmt.Println(cap(sl3))

}
*/

// copy
/*
func main(){
	sl := []int{100, 200}
	sl2 := sl

	sl2[0] = 1000
	fmt.Println(sl)

	var i int = 10
	i2 := i
	i2 = 100
	fmt.Println(i, i2)

	sl := []int{1, 2, 3, 4, 5}
	sl2 := make([]int, 5, 10)
	fmt.Println(sl2)

	n := copy(sl2, sl)

	fmt.Println(n, sl2)
}
*/

// for
/*
func main() {
	sl := []string{"A", "B", "C"}
	fmt.Println(sl)

	// for i := range sl {
	// 	fmt.Println(i, sl[i])
	// }

	for i := 0; i < len(sl); i++ {
		fmt.Println(sl[i])
	}

}
*/

// 可変長引数
/*
func Sum(s ...int) int {
	n := 0
	for _, v := range s {
		n += v
	}
	return n
}

func main() {
	fmt.Println(Sum(1, 2, 3))

	fmt.Println(Sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))

	fmt.Println(Sum())

	sl := []int{1, 2, 3}
	fmt.Println(Sum(sl...))

}
*/

// マップ

/*
func main() {
	var m = map[string]int{"A": 100, "B": 200}
	fmt.Println(m)

	m2 := map[string]int{"A": 100, "B": 200}
	fmt.Println(m2)

	m3 := map[int]string{
		1: "A",
		2: "B",
	}
	fmt.Println(m3)

	m4 := make(map[int]string)
	fmt.Println(m4)

	m4[1] = "japan"
	m4[2] = "usa"
	fmt.Println(m4)

	fmt.Println(m["A"])
	fmt.Println(m4[2])
	fmt.Println(m4[3])

	_, ok := m4[3]
	if !ok {
		fmt.Println("error")
	}
	// fmt.Println(s, ok)

	m4[2] = "us"
	fmt.Println(m4)

	m4[3] = "china"
	fmt.Println(m4)

	delete(m4, 3)
	fmt.Println(m4)

	fmt.Println(len(m4))

}
*/

// for
/*
func main() {
	m := map[string]int{
		"apple": 100,
		"banana": 200,
	}

	for k := range m {
		fmt.Println(k)
	}
}
*/

// channel
// 複数のゴルーチン間でのデータの受け渡しをするために設計されたデータ構造
// 宣言、操作
/*
func main() {
	var ch1 chan int

	// 受信専用
	// var ch2 <-chan int

	// 送信専用
	// var ch3 chan<- int

	ch1 = make(chan int)

	ch2 := make(chan int)

	fmt.Println(cap(ch1))
	fmt.Println(cap(ch2))

	ch3 := make(chan int, 5)
	fmt.Println(cap(ch3))

	ch3 <- 1
	fmt.Println(len(ch3))

	ch3 <- 2
	ch3 <- 3
	fmt.Println("len", len(ch3))

	i := <-ch3
	fmt.Println(i)
	fmt.Println("len", len(ch3))

	i2 := <-ch3
	fmt.Println(i2)
	fmt.Println("len", len(ch3))

	fmt.Println(<-ch3)
	fmt.Println("len", len(ch3))

	ch3 <- 1
	fmt.Println("len", len(ch3))
	fmt.Println(<-ch3)
	fmt.Println("len", len(ch3))
	ch3 <- 2
	ch3 <- 3
	ch3 <- 4
	ch3 <- 5
	ch3 <- 6


}
*/

// チャンネルとゴルーチン
/*
func receiver(c chan int) {
	for {
		i := <-c
		fmt.Println(i)
	}
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// fmt.Println(<-ch1)

	go receiver(ch1)
	go receiver(ch2)

	i := 0
	for i < 100 {
		ch1 <- i
		ch2 <- i
		time.Sleep(50 * time.Millisecond)
		i++
	}

}
*/

// close
/*
func receiver(name string, ch chan int){
	for{
		i, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(name, i)
	}
	fmt.Println(name + "END")
}

func main() {
	ch1 := make(chan int, 2)

	// ch1 <- 1

	// close(ch1)

	// // ch1 <- 1

	// // fmt.Println(<-ch1)

	// i, ok := <-ch1
	// fmt.Println(i, ok)

	// i2, ok := <-ch1
	// fmt.Println(i2, ok)

	go receiver("1.goroutin", ch1)
	go receiver("2.goroutin", ch1)
	go receiver("3.goroutin", ch1)

	i := 0
	for i < 100 {
		ch1 <- i
		i++
	}
	close(ch1)
	time.Sleep(3 * time.Second)
}
*/

// for
/*
func main() {
	ch1 := make(chan int, 3)
	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	close(ch1)
	for i := range ch1 {
		fmt.Println(i)
	}
}
*/

// select

func main() {
	ch1 := make(chan int, 2)
	ch2 := make(chan string, 2)

	// ch2 <- "A"
	// ch1 <- 1
	// ch2 <- "B"
	// ch1 <- 2
	// v1 := <-ch1
	// v2 := <-ch2
	// fmt.Println(v1)
	// fmt.Println(v2)

	select {
	case v1 := <-ch1:
		fmt.Println(v1 + 1000)
	case v2 := <-ch2:
		fmt.Println(v2 + "!?")
	default:
		fmt.Println("どちらでもない")
	}

	ch3 := make(chan int)
	ch4 := make(chan int)
	ch5 := make(chan int)

	// receiver
	go func() {
		for {
			i := <-ch3
			ch4 <- i * 2
		}
	}()

	go func() {
		for {
			i2 := <-ch4
			ch5 <- i2 - 1
		}
	}()

	n := 0
L:
	for {
		select {
		case ch3 <- n:
			n++
		case i3 := <-ch5:
			fmt.Println("receiver", i3)
		default:
			if n > 100 {
				break L
			}
		}
		// if n > 100 {
		// 	break
		// }
	}

}
