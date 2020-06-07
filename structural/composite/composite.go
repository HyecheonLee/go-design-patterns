package composite

type Athlete struct {
}

func (r *Athlete) Train() {
	println("Training")
}
func Swim() {
	println("Swimming!")
}

type CompositeSwimmerA struct {
	MyAthlete Athlete
	MySwim    func()
}

//

type Animal struct{}

func (r *Animal) Eat() {
	println("Eating")
}

type Shark struct {
	Animal
	Swim func()
}

type Swimmer interface {
	Swim()
}

type Trainer interface {
	Train()
}
type SwimmerImpl struct{}

func (s *SwimmerImpl) Swim() {
	println("Swimming!")
}

type CompositeSwimmerB struct {
	Trainer
	Swimmer
}

type Tree struct {
	LeafValue int
	Right     *Tree
	Left      *Tree
}
type Parent struct {
	SomeField int
}

type Son struct {
	P Parent
}

func GetParentField(p Parent) int {
	return p.SomeField
}
