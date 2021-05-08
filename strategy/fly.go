package strategy

import "fmt"

// Interface that all flying behavior classes impl
type flyBehavior interface {
	fly()
}

// Flying behavior impl for ducks that fly
type flyWithWings struct{}

func (fw *flyWithWings) fly() {
	fmt.Println("I'm flying")
}

// Flying behavior impl for ducks that do not fly (rubber/decoy ducks)
type flyNoWay struct{}

func (fnw *flyNoWay) fly() {
	fmt.Println("I can't fly")
}
