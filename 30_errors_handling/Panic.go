package main

import (
	"fmt"
	"net/url"
	"os"
)

func main() {
	//var zero int
	//_ = 42 / zero // Runtime error: integer divide by zero - целое число делится на ноль

	defer func() {
		if e := recover(); e != nil { // Приходит в себя после panic
			fmt.Println(e) // Выводит: Я забыл свое полотенце
		}
	}()

	//panic("Я забыл свое полотенце") // Приводит к panic

	u, err := url.Parse("https://a b.com/")

	if err != nil {
		fmt.Println(err)         // Выводит: parse https://a b.com/: invalid character “ ” in host name
		fmt.Printf("%#v\n", err) // Выводит: &url.Error{Op:“parse”, URL:“https://a b.com/”, Err:” “}

		if e, ok := err.(*url.Error); ok {
			fmt.Println("Op:", e.Op)   // Выводит: Op: parse
			fmt.Println("URL:", e.URL) // Выводит: URL: https://a b.com/
			fmt.Println("Err:", e.Err) // Выводит: Err: invalid character “ ” in host name
		}
		os.Exit(1)
	}

	fmt.Println(u)
}
