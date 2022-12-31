package main

import (
	"fmt"
	"log"
	"os"
)

// os
/*
func main() {
	// Exit
	// os.Exit(1)
	// fmt.Println("Start")

	defer func() {
		fmt.Println("defer")
	}()
	os.Exit(0)
}
*/

/*
func main() {
	//log.Fatal
	_, err := os.Open("A.txt")
	if err != nil {
		log.Fatalln(err)
	}
}
*/

/*
func main() {
	// Args
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])
	fmt.Println(os.Args[2])
	fmt.Println(os.Args[3])

	//os.Argsの要素数を表示
	fmt.Printf("length=%d\n", len(os.Args))
	//os.Argsの中身を全て表示
	for i, v := range os.Args {
		fmt.Println(i, v)
	}
}
*/

/*
func main() {
	//ファイル操作
	//Open
	f, err := os.Open("test1.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

}
*/

/*
func main() {
	//ファイル操作
	// Create
	f, _ := os.Create("foo.txt")

	f.Write([]byte("Hello\n"))

	f.WriteAt([]byte("Golang"), 6)

	f.Seek(0, os.SEEK_END)

	f.WriteString("Yaah")
}
*/
/*
func main() {
	//ファイル操作
	//Read
	f, err := os.Open("foo.txt")
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	bs := make([]byte, 128)

	n, err := f.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))

	bs2 := make([]byte, 128)

	nn, err := f.ReadAt(bs2, 10)
	fmt.Println(nn)
	fmt.Println(string(bs2))
}
*/

func main() {
	// ファイル操作
	//OpenFile
	//O_RDONLY 読み込み専用
	//O_WRONLY 書き込み専用
	//O_RDWR 読み書き専用
	//O_APPEND ファイル末尾に追記
	//O_CREATE ファイルがなければ作成
	//O_TRUNC 可能であればファイルの内容をオープン時に空にする
	f, err := os.OpenFile("foo.txt", os.O_RDONLY | os.O_CREATE, 0666)
	if err != nil {
		log.Fatalln(err)
	}

	defer f.Close()

	bs := make([]byte, 128)
	n, err := f.Read(bs)
	fmt.Println(n)
	fmt.Println(string(bs))

}