package main

import (
	"fmt"
	"regexp"
)

func main() {
	/*
	//Goの正規表現の基本 regexp.MatchString
	match, _ := regexp.MatchString("A", "ABC")
	fmt.Println(match)


	
	// Compile
	re1, _ := regexp.Compile("A")
	match = re1.MatchString("ABC")
	fmt.Println(match)
	


	// MustCompile
	re2 := regexp.MustCompile("A")
	match = re2.MatchString("ABC")
	fmt.Println(match)



	regexp.MustCompile("\\d")
	regexp.MustCompile(`\d`)
	*/

	//正規表現のフラグ

/*
	// フラグ一覧

	i 大文字小文字を区別しない
	m マルチラインモード（^と&が文頭、文末に加えて行頭、行末にマッチ）
	s .が\nにマッチ
	U 最小マッチへの変換（x*はx*?へ、x+はx+?へ）
*/
	/*
	re3 := regexp.MustCompile(`(?i)abc`)
	match := re3.MatchString("ABC")
	fmt.Println(match)
	*/

	//幅を持たない正規表現のパターン
	/*
		パターン一覧

		^ 文頭（mフラグが有効な場合は行頭にも）
		$ 文末（mフラグが有効な場合は行末にも）
		\A 文頭
		\z 文末
		\b ASCIIによるワード協会
		\B 非ASCIIによるワード協会
	*/

	re4 := regexp.MustCompile(`^ABC$`)
	match := re4.MatchString("ABC")
	fmt.Println(match)

	match = re4.MatchString(" ABC ")
	fmt.Println(match)
}