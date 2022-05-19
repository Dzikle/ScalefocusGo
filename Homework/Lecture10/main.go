package main

import (
	"fmt"
	"sync"
	"time"
)

type ConcurrentPrinter struct {
	sync.WaitGroup
	sync.Mutex
	last  string
	count int
}

func (cp *ConcurrentPrinter) printFoo(times int) {
	cp.Add(1)

	go func() {
		defer cp.Done()
		for {
			if cp.count == times {
				break
			}

			cp.Lock()
			if cp.last != "foo" {
				fmt.Print("foo")
				cp.last = "foo"
				cp.count++
			}
			cp.Unlock()
			time.Sleep(100 + time.Millisecond)
		}
	}()

}
func (cp *ConcurrentPrinter) printBar(times int) {
	cp.Add(1)

	go func() {
		defer cp.Done()
		for {
			if cp.count == times {
				break
			}

			cp.Lock()
			if cp.last != "bar" {
				fmt.Print("bar")
				cp.last = "bar"
				cp.count++
			}
			cp.Unlock()
			time.Sleep(100 + time.Millisecond)
		}
	}()
}

func main() {
	times := 10
	cp := &ConcurrentPrinter{}

	cp.printBar(times)
	cp.printFoo(times)
	cp.Wait()

}
