package main

import (
	"fmt"
	"math/rand"
	"time"
)

//rand

func main() {
	rand.Seed(256)

	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	fmt.Println(rand.Float64())
	/*
		0.23143513611531310
		0.85451313513135135
		0.13259548796632156
	*/

	//現在の情報をシードに使った疑似乱数の生成
	fmt.Println(time.Now().UnixNano())
	rand.Seed(time.Now().UnixNano())
	//0〜99の間の疑似乱数
	fmt.Println(rand.Intn(100))
	//int型の疑似乱数
	fmt.Println(rand.Int())
	// int32型の疑似乱数
	fmt.Println(rand.Int31())
	// int64型の疑似乱数
	fmt.Println(rand.Int63())
	// int32型の疑似乱数
	fmt.Println(rand.Uint32())
}
