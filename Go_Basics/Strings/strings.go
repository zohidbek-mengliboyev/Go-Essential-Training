package main

import (
	"fmt"
)

func main() {
	book := "The color of magic"
	fmt.Println(book)

	fmt.Println(len(book))

	fmt.Printf("book[0] = %v (type %T)\n", book[0], book[0])
	// uint8 is byte
	// this is my string task
}
