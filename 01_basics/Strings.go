package main

import "fmt"

func main() {
	smile := '😃'
	fmt.Printf("%c %[1]v\n", smile)

	hello := "hello"
	for i := 0; i < len(hello); i++ {
		fmt.Printf("%c\n", hello[i])
	}
}
