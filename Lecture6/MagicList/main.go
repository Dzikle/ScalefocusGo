package main

import "fmt"

type MagicList struct {
	LastItem *Item
}
type Item struct {
	Value    int
	PrevItem *Item
}

//For review
func add(l *MagicList, value int) {

	// I dont quite get what to do if the it's nil, However the code works!!!
	if l.LastItem == nil {
		l.LastItem = &Item{Value: value}
	} else {
		l.LastItem = &Item{Value: value, PrevItem: l.LastItem}
	}
}

var MList = []int{}

func toSlice(l *MagicList) []int {

	if l.LastItem != nil {
		MList = append(MList, l.LastItem.Value)
		l1 := &MagicList{l.LastItem.PrevItem}
		toSlice(l1)

	}
	revList := []int{}
	for i := len(MList) - 1; i >= 0; i-- {
		revList = append(revList, MList[i])
	}

	return revList
}

func main() {

	l := &MagicList{}
	add(l, 10)
	add(l, 22)
	add(l, 44)
	add(l, 55)
	add(l, 65)
	add(l, 74)
	add(l, 82)
	add(l, 85)
	add(l, 98)
	add(l, 100)

	fmt.Println(toSlice(l))
}
