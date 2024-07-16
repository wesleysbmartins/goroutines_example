package goroutines

import (
	"fmt"
	"time"
)

func ExampleOne() {
	go Print("Goroutine 01") // run new goroutine
	time.Sleep(1 * time.Second)
	fmt.Println("Goroutine Principal")
}
