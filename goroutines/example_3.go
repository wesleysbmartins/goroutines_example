package goroutines

import "fmt"

type structExample struct {
	n int
	s string
	b bool
}

func ExampleThre() {
	myStruct := make(chan structExample) // create channel

	go create(&myStruct) // run new goroutine

	close(myStruct) // close channel

	fmt.Println("Value", <-myStruct) // read channel

	fmt.Println("Goroutine principal")

}

func create(myStruct *chan structExample) {
	value := structExample{
		n: 10,
		s: "str",
		b: true,
	}

	*myStruct <- value // write channel
}
