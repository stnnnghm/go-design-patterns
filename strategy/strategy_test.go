package strategy

func ExampleFly() {
	fw := &flyWithWings{}
	fw.fly()

	fnw := &flyNoWay{}
	fnw.fly()

	// Unordered output:
	// I'm flying
	// I can't fly
}
