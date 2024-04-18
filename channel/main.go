package main

import "fmt"

func main() {
	chString := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			// TODO: send data to channel
			chString <- fmt.Sprintf("Data ke - %d", i)
		}
	}()

	for i := 0; i < 10; i++ {
		// TODO: receive data from channel
		data := <-chString
		fmt.Println(data)
	}
}
