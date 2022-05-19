package main

import (
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	urls := []string{"Google", "Favebook", "twitter", "A1On"}

	fetchURLs(urls, 3)
}

var wg sync.WaitGroup

func fetchURLs(urls []string, concurency int) {
	//buffered
	processQueue := make(chan string, concurency)
	//unbuffered
	done := make(chan string)
	var wg sync.WaitGroup
	go func() {
		for _, urlToProcces := range urls {
			wg.Add(1)
			processQueue <- urlToProcces

			go func(url string) {
				defer wg.Done()
				//simulate work
				time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
				log.Println("fetched: ", url)
				<-processQueue
				done <- url
			}(urlToProcces)
		}

	}()
	for range urls {
		log.Println("Done: ", <-done)
	}
	wg.Done()

}
