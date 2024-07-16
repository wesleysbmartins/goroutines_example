package goroutines

import (
	"fmt"
	"time"
)

func ExampleTwo() {
	channel := make(chan string) //create channel

	go write("Loop", &channel)

	close(channel) // close channel

	fmt.Println("Value", <-channel) // read channel

	fmt.Println("Goroutine principal")
}

func write(s string, channel *chan string) {

	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s, i)
	}

	*channel <- "Finished" // write channel

	fmt.Println("Write finish")
}
