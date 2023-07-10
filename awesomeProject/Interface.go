package main

import (
	"math"
)

type Rectangle struct {
	x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := (x1 - x2) * (x1 - x2)
	b := (y1 - y2) * (y1 - y2)
	return math.Sqrt(a + b)
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x2, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

//func main() {
//	r := Rectangle{0, 0, 10, 10}
//	c := Circle{20, 1, 10}
//	fmt.Println(c.area())
//	fmt.Println(r.area())
//	fmt.Println("Total area: ", totalArea(&c, &r))
//}
