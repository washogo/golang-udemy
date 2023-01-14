package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//net/http
//クライアント

func main() {

	//GET
	/*
		res, _ := http.Get("https://exaple.com")

		fmt.Println(res.StatusCode)

		fmt.Println(res.Proto)

		fmt.Println(res.Header["Date"])
		fmt.Println(res.Header["Content-type"])

		fmt.Println(res.Request.Method)
		fmt.Println(res.Request.URL)

		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)
		fmt.Print(string(body))
	*/

	//-------------------------------------------
	//POST

	vs := url.Values{}

	vs.Add("id", "1")
	vs.Add("message", "メッセージ")
	fmt.Println(vs.Encode())

	res, err := http.PostForm("http://example.com/", vs)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Print(string(body))

}
