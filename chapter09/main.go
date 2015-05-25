package main

import (
	"fmt"
	"math"
)

// Shape is a shape
type Shape interface {
	area() float64
	perimeter() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

// MultipleShape is MultipleShape
type MultipleShape struct {
	shapes []Shape
}

func (m *MultipleShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
}

// Embedded Types

// Person is a person
type Person struct {
	Name string
}

// Android is a person
type Android struct {
	Person
	Model string
}

// Talk is a function
func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

// Circle is a circle.
type Circle struct {
	x, y, r float64
}

// Rectangle is a Rectangle
type Rectangle struct {
	x1, y1, x2, y2 float64
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}
func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func rectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x2, y2)
	w := distance(x1, y1, x2, y2)
	return l * w
}

func circleArea(x, y, r float64) float64 {
	return math.Pi * r * r
}

func circleAreaStruct(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.r
}

func (r *Rectangle) perimeter() float64 {
	return r.x1 + r.x2 + r.y1 + r.y2
}

func main() {

	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	var cx, cy, cr float64 = 0, 0, 5

	fmt.Println(rectangleArea(rx1, ry1, rx2, ry2))
	fmt.Println(circleArea(cx, cy, cr))

	//both ways can be used to create an instance
	//var c Circle
	//c := new(Circle)
	//c := Circle{x: 0, y: 0, r: 5}
	c := Circle{0, 0, 5}
	fmt.Println(circleAreaStruct(&c))
	fmt.Println(c.area())

	a := new(Android)
	a.Person.Talk()

	r := Rectangle{0, 0, 10, 10}
	fmt.Println(r.area())

	fmt.Println(totalArea(&c, &r))
}
