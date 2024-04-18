package main

func add(a chan int, b chan int) {
	// TODO: receive data from channel a

	// TODO: add 2 to x

	// TODO: send data to channel b
}

func main() {
	a := make(chan int)
	b := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			a <- i
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			add(a, b)
		}
	}()

	for i := 0; i < 10; i++ {
		data := <-b
		println(data)
	}
}
