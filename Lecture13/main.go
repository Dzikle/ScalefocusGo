package main

import (
	"context"
	"fmt"
	"time"
)

type BufferedContext struct {
	context.Context
	buffer chan string
	context.CancelFunc
}

func NewBufferedContext(timeout time.Duration, bufferSize int) *BufferedContext {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	buff := make(chan string, bufferSize)
	NewBufferedCTX := &BufferedContext{Context: ctx, buffer: buff, CancelFunc: cancel}
	return NewBufferedCTX

}
func (bc *BufferedContext) Done() <-chan struct{} {
	if len(bc.buffer) == cap(bc.buffer) {
		bc.CancelFunc()
		fmt.Println("TIME LIMIT")
		time.Sleep(time.Millisecond * 200)
	}

	// ctx, cancel := context.WithTimeout(bc, bc.timeout)
	// defer cancel()
	// return ctx.Done()
	return bc.Context.Done()
}

func (bc *BufferedContext) Run(fn func(context.Context, chan string)) {
	// ch := make(chan string, bc.buffer)
	// for {
	// 	if len(ch) != cap(ch) {
	// 		ch <- "bar"
	// 		time.Sleep(time.Millisecond * 2000)
	// 	}
	fn(bc, bc.buffer)
	// 	close(ch)
	// 	return
	// }
}

func main() {
	ctx := NewBufferedContext(55*time.Second, 5)
	fmt.Println("start")
	ctx.Run(func(ctx context.Context, buffer chan string) {
		for {
			select {
			case <-ctx.Done():
				return
			case buffer <- "bar":
				time.Sleep(time.Millisecond * 20)
				fmt.Println("bar")
			}
		}
	})
}
