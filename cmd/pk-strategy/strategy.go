package main

import "github.com/stnnnghm/go-design-patterns/strategy"

func main() {
	mallardDuck := strategy.NewMallardDuck()
	modelDuck := strategy.NewModelDuck()

	mallardDuck.Display()
	mallardDuck.PerformFly()
	mallardDuck.PerformQuack()
	mallardDuck.PerformSwim()

	modelDuck.Display()
	modelDuck.PerformFly()
	modelDuck.PerformQuack()
	modelDuck.PerformSwim()
}
