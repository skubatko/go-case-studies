package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	c := make(chan int) // Делает канал для связи

	for i := 0; i < 5; i++ {
		go sleepyGopherSelect(i, c)
	}

	timeout := time.After(3 * time.Second)
	for i := 0; i < 5; i++ {
		select { // Оператор select
		case gopherID := <-c: // Ждет, когда проснется гофер
			fmt.Println("gopher ", gopherID, " has finished sleeping")
		case <-timeout: // Ждет окончания времени
			fmt.Println("my patience ran out")
			return // Сдается и возвращается
		}
	}
}

func sleepyGopherSelect(id int, c chan int) {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	c <- id
}
