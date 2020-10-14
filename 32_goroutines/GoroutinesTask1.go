package main

import (
	"fmt"
)

func sourceItems(downstream chan string) {
	for _, v := range []string{"hello", "hello", "world", "world", "goodbye", "goodbye all"} {
		downstream <- v
	}
	close(downstream)
}

func filterDuplicates(upstream, downstream chan string) {
	var s string
	for {
		item, ok := <-upstream
		if !ok {
			close(downstream)
			return
		}
		if item != s {
			downstream <- item
			s = item
		}
	}
}

func printItems(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go sourceItems(c1)
	go filterDuplicates(c1, c2)
	printItems(c2)
}
