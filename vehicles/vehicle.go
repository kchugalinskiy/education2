package vehicles

import (
	"fmt"

	"github.com/kchugalinskiy/education2/sumrilsht"
)

func g(arr *[]int) {
	tmp := append(*arr, 6)
	*arr = tmp
}

func (c *Car) Drive() {
	fmt.Printf("Drive %q car steering %f\n", c.Brand, c.steering)
}

func (c *Car) explode() {
	println("boom!")
}

type Car struct {
	Brand    string  // public
	steering float64 // private (internal?)
}

func NewCar(brand string) *Car {
	blabla()
	return &Car{Brand: brand, steering: 1.0}
}

func blabla() {
	println("blabla")
}

type Vehicle interface { // contract ?
	Drive()
}

func print() {
	c1 := NewCar("BMW")
	c2 := NewCar("Haval")
	t1 := sumrilsht.NewBSTrain("A123")

	vs := []Vehicle{c1, c2, t1}
	for _, v := range vs {
		v.Drive()
	}
}
