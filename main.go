package main

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

type Luggage struct {
	weight    int
	passenger string
}

func NewLuggage(weight int, passenger string) *Luggage {
	l := Luggage{weight, passenger}
	return &l
}

type Belt []*Luggage

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

}
