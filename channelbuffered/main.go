package main

import "fmt"

func main() {
	chString := make(chan string, 3)

	go func() {
		for i := 0; i < 10; i++ {
			// TODO: send data to channel
		}
	}()

	for i := 0; i < 10; i++ {
		// TODO: receive data from channel
		fmt.Println(data)
	}
}
