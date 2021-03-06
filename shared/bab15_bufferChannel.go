package shared

import (
	"fmt"
	"runtime"
)

func ContohBufferChannel() {
	runtime.GOMAXPROCS(1)

	messages := make(chan int, 1)

	go func() {
		for {
			i := <-messages
			fmt.Println("receive data", i)
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Println("send data", i)
		messages <- i
	}
}
