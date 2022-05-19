package main

import (
	"log"
	"testing"
)

func TestReverseString(t *testing.T) {

	input := "apple"
	output := "elppa"

	res := NewReversStringReader(input)

	if res.input != output {
		log.Fatalf("The %s, should be equal to %s", res.input, output)
	}
}
