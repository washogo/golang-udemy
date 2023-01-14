package main

import (
	"fmt"

	"github.com/google/uuid"
)

//UUID

func main() {
	uuidObj, _ := uuid.NewUUID()
	fmt.Println("  ", uuidObj.String())

	uuidObj2, _ := uuid.NewRandom()
	fmt.Println("  ", uuidObj2.String())
}
