package main

import (
	"fmt"
	"time"
)

func main() {

	ch := generateThrothled("Foo", 2, time.Second)

	for {
		msg := <-ch + <-ch
		fmt.Println(msg)
	}

}

func generateThrothled(data string, buffer int, clearInterval time.Duration) <-chan string {
	ch := make(chan string, buffer)

	go func(data string, ci time.Duration) {
		for i := 0; i < 10; i++ {
			ch <- data
			time.Sleep(ci)
		}
	}(data, clearInterval)

	return ch
}
