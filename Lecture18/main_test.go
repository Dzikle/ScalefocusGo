package main

import (
	"testing"
	"time"
)

const max = 100

func Benchmark100PrimesWith0MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primesAndSleep(max, time.Millisecond*0)
	}
}
func Benchmark100PrimesWith5MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primesAndSleep(max, time.Millisecond*5)
	}
}
func Benchmark100PrimesWith10MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primesAndSleep(max, time.Millisecond*10)
	}
}

func Benchmark100GoPrimesWith0MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primeNumbers(max, time.Millisecond*0)
	}
}
func Benchmark100GoPrimesWith5MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primeNumbers(max, time.Millisecond*5)
	}
}
func Benchmark100GoPrimesWith10MSSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		primeNumbers(max, time.Millisecond*10)
	}
}
