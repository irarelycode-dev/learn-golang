package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perim() float32
}

type rectangle struct {
	width, height float32
}

type circle struct {
	radius float64
}

func (r rectangle) area() float64 {
	return float64(r.width * r.height)
}

func (r rectangle) perim() float32 {
	return 2 * (r.width + r.height)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float32 {
	return float32(2 * math.Pi * c.radius)
}

func geometricalMethods(g geometry) {
	fmt.Println("g:", g)
	fmt.Println("area:", g.area())
	fmt.Println("perim:", g.perim())
}

func callInterfaceMethods() {
	geometricalMethods(rectangle{height: 5, width: 2})
	geometricalMethods(circle{radius: 10})
}
