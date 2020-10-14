package main

import (
	"fmt"
	"strings"
)

func sourceStrings(downstream chan string) {
	for _, v := range []string{"hello world goodbye all", "a b c d"} {
		downstream <- v
	}
	close(downstream)
}

func splitStrings(upstream, downstream chan string) {
	for {
		item, ok := <-upstream
		if !ok {
			close(downstream)
			return
		}
		for _, word := range strings.Fields(item) {
			downstream <- word
		}
	}
}

func printWords(upstream chan string) {
	for v := range upstream {
		fmt.Println(v)
	}
}

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	go sourceStrings(c1)
	go splitStrings(c1, c2)
	printWords(c2)
}
