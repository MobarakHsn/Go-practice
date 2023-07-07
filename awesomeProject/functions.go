package main

import (
	"fmt"
	"math"
)

func average(xs []float64) float64 {
	fmt.Println("Nothing to do")
	xs[0] = 5
	return 0
}
func f() (a int, b int) {
	return 5, 6
}

func add(x ...int) int {
	//fmt.Printf("%T", x)
	total := 0
	for _, v := range x {
		total += v
	}
	return total
}
func makeEvenGenerator() func() int {
	i := 0
	return func() (ret int) {
		ret = i
		i += 2
		return
	}
}
func first() {
	fmt.Println("First")
}
func second() {
	fmt.Println("Second")
}
func third() {
	fmt.Println("Third")
}

func zero(ptr *int) {
	*ptr = 0
}

type Circle struct {
	x, y, r float64
}

// This is a method
func (cir *Circle) area() float64 {
	return math.Pi * cir.r * cir.r
}

// Here we can see is a relationship with structures
type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi my name is", p.Name)
}

type Android struct {
	Person
	Model string
}

func main() {

	an := new(Android)
	an.Name = "Mobarak"
	an.Talk()
	return
	//Struct
	var cir Circle
	cir.r = 10
	fmt.Println(cir)

	var cir1 = new(Circle)
	cir1 = &cir
	fmt.Println(cir1.y)

	cir3 := Circle{20, 1, 10}
	fmt.Println(cir3)
	//Methods
	fmt.Println(cir3.area())
	//Pointer
	var pptr *int
	x := 5
	pptr = &x
	z := 20
	pptr = &z
	fmt.Println(*pptr)
	zero(&x)
	fmt.Println(x)
	ptr := new(int)
	fmt.Println(*ptr)

	//Panic & recover
	defer func() {
		str := recover()
		fmt.Println(str)
	}()
	panic("Onek PanicxD")

	//Defer
	defer first()
	third()
	defer second()

	//Closure
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven()) //0
	fmt.Println(nextEven()) //2
	fmt.Println("Hello")
	arr := []float64{1, 2, 3}
	brr := []int{1, 2, 3}

	//multiple Args
	fmt.Println(add(brr[0:1]...))
	fmt.Println(add(1, 2, 3))
	fmt.Println(average(arr))
	fmt.Println(arr)
	x, y := f()
	fmt.Println(x, y)

	divide := func(a, b int) int {
		return a / b
	}
	fmt.Println(divide(3, 2))
	in := 0
	increment := func() int {
		in++
		return in
	}
	fmt.Println(increment())
}
