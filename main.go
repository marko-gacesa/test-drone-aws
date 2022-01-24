package main

import (
	"fmt"
	"github.com/satori/go.uuid"
)

func main() {
	fmt.Println("Hello!")

	id := uuid.NewV4()
	fmt.Println(id)
}
