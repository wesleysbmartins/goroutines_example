# Golang Goroutines
Exemplos de como as Goroutines funcionam em Golang.

- Exemplo 01

Goroutines comuns com utilização de timeSleep para que de tempo de finalizar a goroutine.

```go
func ExampleOne() {
	go Print("Goroutine 01") // run new goroutine
	time.Sleep(1 * time.Second)
	fmt.Println("Goroutine Principal")
}
```

- Exemplo 02

Goroutines utilizando channels, uma abordagem para maniular os dados entre as goroutines, onde um channel criado, um valor é inserido, e só após a inserção a leitura pode ser realizada.
Sendo assim, sua aplicação vai esperar a liberação do channel para seguir com as próximas instruções.

```go
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
```
- Exmplo 3

Goroutines com channels do tipo struct, onde a diferença é apenas o tipo esperado pelo channel.

```go
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
```
