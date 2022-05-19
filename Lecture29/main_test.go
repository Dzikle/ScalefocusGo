package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGroupBy(t *testing.T) {
	//Arrange
	expectRes := map[string][]Order{"John": {{Customer: "John", Amount: 1000}, {Customer: "John", Amount: 1200}}, "Sara": {{Customer: "Sara", Amount: 2000}, {Customer: "Sara", Amount: 1800}}}

	var input = []Order{
		{Customer: "John", Amount: 1000},
		{Customer: "Sara", Amount: 2000},
		{Customer: "Sara", Amount: 1800},
		{Customer: "John", Amount: 1200},
	}

	result := GroupBy(input, func(o Order) string { return o.Customer })

	res1 := reflect.DeepEqual(result, expectRes)
	fmt.SprintF("Is Map %s is equal to Map %s: ", result, expectRes)

}
