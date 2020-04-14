package main

import "fmt"

// Action any action taken by the user
type Action struct {
	name string
	next *Action
}

// ActionHistory a stack of all actions taken
type ActionHistory struct {
	top  *Action
	size int
}

// Luggage represents real world luggage
type Luggage struct {
	weight    int
	passenger string
}

// NewLuggage is a constructor for Luggage
func NewLuggage(weight int, passenger string) *Luggage {
	l := Luggage{weight, passenger}
	return &l
}

// Belt custom type slice of Luggage
type Belt []*Luggage

// Add luggage to the Belt
func (b *Belt) Add(l *Luggage) {
	*b = append(*b, l)
}

// Take is similar to an shift
func (b *Belt) Take() *Luggage {
	first, rest := (*b)[0], (*b)[1:]
	*b = rest
	return first
}

// Add to the stack
func (ah *ActionHistory) Add(newAction *Action) {
	if ah.top != nil {
		oldTop := ah.top
		newAction.next = oldTop
	}
	ah.top = newAction
	ah.size++
}

// Undo the last action
func (ah *ActionHistory) Undo() *Action {
	topAction := ah.top
	if topAction != nil {
		ah.top = topAction.next
	} else if topAction.next == nil {
		ah.top = nil
	}
	ah.size--
	return topAction
}

func main() {
	belt := &Belt{}

	belt.Add(NewLuggage(15, "Red"))
	belt.Add(NewLuggage(4, "Blue"))
	belt.Add(NewLuggage(7, "Orange"))
	belt.Add(NewLuggage(3, "Black"))

	fmt.Println("Belt:", belt, "Length:", len(*belt))

	f := belt.Take()

	fmt.Println("First luggage:", f)
	fmt.Println("Belt:", belt, "Length:", len(*belt))

}
