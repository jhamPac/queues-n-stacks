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

// Apply add to the stack
func (ah *ActionHistory) Apply(newAction *Action) {
	if ah.top != nil {
		oldTop := ah.top
		newAction.next = oldTop
	}
	ah.top = newAction
	ah.size++
}

func main() {

}
