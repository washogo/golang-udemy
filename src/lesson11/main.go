package main

import "fmt"

// 構造体
/*
type User struct {
	Name string
	Age  int
	// X, Y int
}

func UpdateUser(user User) {
	user.Name = "A"

	user.Age = 1000
}

func UpdateUser2(user *User) {
	user.Name = "A"
	user.Age = 1000
}

func main() {
	var user1 User
	fmt.Println(user1)
	user1.Name = "user1"
	user1.Age = 10
	fmt.Println(user1)

	user2 := User{}
	fmt.Println(user2)
	user2.Name = "user2"
	fmt.Println(user2)

	user3 := User{Name: "user3", Age: 30}
	fmt.Println(user3)

	user4 := User{"user4", 40}
	fmt.Println(user4)

	// user5 := User{30, "user5"}
	// fmt.Println(user5)

	user6 := User{Name: "user6"}
	fmt.Println(user6)

	user7 := new(User)
	fmt.Println(user7)

	user8 := &User{}
	fmt.Println(user8)

	UpdateUser(user1)
	UpdateUser2(user8)

	fmt.Println(user1)
	fmt.Println(user8)
}
*/

// structメソッド
/*
type User struct {
	Name string
	Age int
	// X, Y int
}

func (u User) SayName() {
	fmt.Println(u.Name)
}

func (u User) SetName(name string) {
	u.Name = name
}

func (u *User) SetName2(name string) {
	u.Name = name
}

func main() {
	user1 := User{Name: "user1"}
	user1.SayName()

	user1.SetName("A")
	user1.SayName()

	user1.SetName2("A")
	user1.SayName()

	user2 := &User{Name: "user2"}
	user2.SetName2("B")
	user2.SayName()
}
*/

// struct 埋め込み
/*
type T struct {
	User User
}

type User struct {
	Name string
	Age int
	//X, Y int
}

func (u *User) SetName() {
	u.Name = "A"
}

func main() {
	t := T{User: User{Name: "user1", Age: 10}}

	fmt.Println(t)

	fmt.Println(t.User)

	fmt.Println(t.User.Name)

	// fmt.Println(t.Name)

	t.User.SetName()
	fmt.Println(t)
}
*/

// type User struct {
// 	Name string
// 	Age int
// }

// struct コンストラクト
/*
func NewUser(name string, age int) *User {
	return &User{Name: name, Age: age}
}

func main() {
	user1 := NewUser("user1", 10)
	
	fmt.Println(user1)

	fmt.Println(*user1)

}
*/

// struct スライス
/*
type Users []*User

// type Users struct {
// 	Users []*User
// }

func main() {
	user1 := User{Name: "user1", Age: 10}
	user2 := User{Name: "user2", Age: 20}
	user3 := User{Name: "user3", Age: 30}
	user4 := User{Name: "user4", Age: 40}

	users := Users{}

	users = append(users, &user1)
	users = append(users, &user2)
	users = append(users, &user3, &user4)

	for _, u := range users{
		fmt.Println(*u)
	}

	users2 := make([]*User, 0)
	users2 = append(users2, &user1, &user2)

	for _, u := range users2{
		fmt.Println(*u)
	}

}
*/

// struct map
/*
func main() {
	m := map[int]User{
		1: {Name: "user1", Age: 10},
		2: {Name: "user1", Age: 20},
	}

	fmt.Println(m)

	m2 := map[User]string{
		{Name: "user1", Age: 10}: "tokyo",
		{Name: "user2", Age: 10}: "LA",
	}
	fmt.Println(m2)

	m3 := make(map[int]User)
	fmt.Println(m3)
	m3[1] = User{Name: "user3"}
	fmt.Println(m3)

	for _, v := range m3 {
		fmt.Println(v)
	}

}
*/

// struct 独自型
/*
type MyInt int

func (mi MyInt) Print() {
	fmt.Println(mi)
}

func main() {
	var mi MyInt
	fmt.Println(mi)
	fmt.Printf("%T\n", mi)

	// i := 100
	// fmt.Println(mi + i)

	mi.Print()
}
*/

// interface
// 最もポピュラーな使い方、異なる型に共通の性質を付与する。

/*
type Stringfy interface {
	ToString() string
}

type Person struct {
	Name string
	Age int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Name=%v, Model=%v", c.Number, c.Model)
}

func main() {
	vs := []Stringfy{
		&Person{Name: "Taro", Age: 21},
		&Car{Number: "123-456", Model: "AB-1234"},
	}

	for _, v := range vs {
		fmt.Println(v.ToString())
	}

}
*/

// カスタムエラー

// type error interface {
// 	Error() string
// }
/*
type MyError struct {
	Message string
	ErrCode int
}

func (e *MyError) Error() string {
	return e.Message
}

func RaiseError() error {
	return &MyError{Message: "カスタムエラーが発生したよ", ErrCode: 1234}
}

func main() {
	err := RaiseError()
	fmt.Println(err.Error())

	e, ok := err.(*MyError)
	if ok {
		fmt.Println(e.ErrCode)
	}

}
*/

// interface
// fmt.Stringer

// type Stringer interface {
// 	String() string
// }

type Point struct {
	A int
	B string
}

func (p *Point) String() string {
	return fmt.Sprintf("<<%v, %v>>", p.A, p.B)
}

func main() {
	p := &Point{100, "ABC"}
	fmt.Println(p)
}
