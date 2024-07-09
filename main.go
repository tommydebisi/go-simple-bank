package main

import (
	"fmt"
	"time"
)

func main() {
	// channels is used to communicate seamlessly between goroutines
	food := make(chan string)
	food2 := make(chan string)

	go func() {
		time.Sleep(2 *time.Second)
		food <- "eba"
	}()

	go func() {
		time.Sleep(4 *time.Second)
		food <- "fish"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg := <-food:
			fmt.Println(msg)
		case msg := <-food2:
			fmt.Println(msg)
		// case <-time.After(1 * time.Second):
		// 	fmt.Println("timeout")
		}
	}

	// fmt.Println(<-food)
}