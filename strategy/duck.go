package strategy

import "fmt"

type abstractDuck struct {
	Display func()
	// Reference vars for the behavior interface types
	// All ducks inherit these (name/interface)
	flyingBehavior   flyBehavior
	quackingBehavior quackBehavior
}

// Delegate fly and quack behavior to respective behavior classes
func (d *abstractDuck) PerformFly() {
	d.flyingBehavior.fly()
}

func (d *abstractDuck) PerformQuack() {
	d.quackingBehavior.quack()
}

func (d *abstractDuck) PerformSwim() {
	fmt.Println("All ducks float")
}

// Setters
func (d *abstractDuck) setFlyingBehavior(fb flyBehavior) {
	d.flyingBehavior = fb
}

func (d *abstractDuck) setQuackingBehavior(qb quackBehavior) {
	d.quackingBehavior = qb
}

type mallardDuck struct {
	*abstractDuck
}

func NewMallardDuck() *mallardDuck {
	// Impl the abstract method
	d := &abstractDuck{
		Display: func() {
			fmt.Println("I'm a mallard duck")
		},
		flyingBehavior:   &flyWithWings{},
		quackingBehavior: &quack{},
	}
	return &mallardDuck{d}
}

type modelDuck struct {
	*abstractDuck
}

func NewModelDuck() *modelDuck {
	// Impl the abstract method
	d := &abstractDuck{
		Display: func() {
			fmt.Println("I'm a model duck")
		},
		flyingBehavior:   &flyNoWay{},
		quackingBehavior: &muteQuack{},
	}
	return &modelDuck{d}
}
