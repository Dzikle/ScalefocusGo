package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

var wg sync.WaitGroup

func primesAndSleep(n int, sleep time.Duration) []int {
	res := []int{}
	for k := 2; k < n; k++ {
		for i := 2; i < n; i++ {
			if k%i == 0 {
				time.Sleep(sleep)
				if k == i {
					res = append(res, k)
				}
				break
			}
		}
	}
	return res
}

func goPrimesAndSleep(n int, sleep time.Duration) []int {
	res := []int{}
	for k := 2; k < n; k++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 2; i < n; i++ {
				if k%i == 0 {
					time.Sleep(sleep)
					if k == i {
						res = append(res, k)
					}
					break
				}
			}
		}()
		wg.Wait()
	}
	return res
}
func primeNumbers(max int, sleep time.Duration) []int {
	var primes []int

	for i := 2; i < max; i++ {
		isPrime := true
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					if i%j == 0 {
						time.Sleep(sleep)
						isPrime = false
					}
				}()
				break
			}
			if isPrime {
				primes = append(primes, i)
			}
		}()
		wg.Wait()
	}

	return primes
}

func main() {
	fmt.Println(primeNumbers(200, time.Millisecond*1))
	fmt.Println(goPrimesAndSleep(200, time.Millisecond*1))
	fmt.Println(primesAndSleep(200, time.Millisecond*1))
}
