package main

import "fmt"

type NamedObj struct {
	Name string
}

// method show
func (n NamedObj) show() {
	fmt.Println(n.Name) // "n" is "this"
}

// class Rectangle
type Rectangle1 struct {
	NamedObj      //inheritance
	Width, Height float64
}

// override method show
func (r Rectangle1) show() {
	fmt.Println("Rectangle1 ", r.Name) // "r" is "this"
}

func main() {
	var a = NamedObj{"Joe"}
	var b = Rectangle1{NamedObj{"Richard"}, 10, 20}

	a.show()
	b.show()
}
