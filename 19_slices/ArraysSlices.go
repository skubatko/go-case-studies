package main

import "fmt"

func printCap(elements []string) {
	fmt.Printf("cap=%v, %v\n", cap(elements), elements)
}

func main() {
	elements := make([]string, 0, 0)
	capacity := cap(elements)
	printCap(elements)
	for i := 0; i < 100; i++ {
		elements = append(elements, fmt.Sprint(i))
		if capacity != cap(elements) {
			capacity = cap(elements)
			printCap(elements)
		}
	}
}
