package strategy

import "fmt"

// Interface that all quacking behavior classes impl
type quackBehavior interface {
	quack()
}

// Quack behavior impl for ducks that quack
type quack struct{}

func (q *quack) quack() {
	fmt.Println("Quack")
}

// Quack behavior impl for ducks that do not quack
type muteQuack struct{}

func (mq *muteQuack) quack() {
	fmt.Println("Silence")
}

// Quack behavior impl for ducks that squeak
type squeak struct{}

func (s *squeak) quack() {
	fmt.Println("Squeak")
}
