package main

import (
	"log"
	"os"
)

func main() {
	//ログの出力先を変更
	log.SetOutput(os.Stdout)
/*
	log.Print("Log\n")
	log.Println("Log2")
	log.Printf("Log%d\n", 3)
*/
/*
	log.Fatal("Log\n")
	log.Fatalln("Log2")
	log.Fatalf("Log%d\n", 3)
*/
/*
	log.Panic("Log\n")
	log.Panicln("Log2")
	log.Panicf("Log%d\n", 3)
*/

// 任意のファイルの作成し、出力先に指定
//os.Create ファイルの作成
f, err := os.Create("test.log")
if err != nil {
	return
}
//作成したio.Writer型のファイルを出力先に指定
log.SetOutput(f)
log.Println("ファイルに書き込む")

}
