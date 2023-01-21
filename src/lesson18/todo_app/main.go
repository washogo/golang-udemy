package main

import (
	"fmt"
	"todo-app/golang-udemy/app/controller"
	"todo-app/golang-udemy/app/models"
)

func main() {
	fmt.Println(models.Db)

	controller.StartMainServer()
}
