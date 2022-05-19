package main

func GroupBy[T any, U comparable](col []T, keyFn func(T) U) map[U][]T {

	seen := make(map[U][]T, len(col))

	for _, item := range col {

		seen[keyFn(item)] = append(seen[keyFn(item)], item)
	}

	return seen
}

type Order struct {
	Customer string
	Amount   int
}

var input = []Order{
	{Customer: "John", Amount: 1000},
	{Customer: "Sara", Amount: 2000},
	{Customer: "Sara", Amount: 1800},
	{Customer: "John", Amount: 1200},
}

func main() {
	GroupBy(input, func(o Order) string { return o.Customer })
}
