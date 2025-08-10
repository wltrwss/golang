package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c *Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r *Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Triangle struct {
	A float64
	B float64
	C float64
}

func (t *Triangle) Area() float64 {
	// Формула Герона
	p := t.Perimeter() / 2
	return math.Sqrt(p * (p - t.A) * (p - t.B) * (p - t.C))
}

func (t *Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

func PrintShapeInfo(s Shape) {
	fmt.Printf("Фигура: %T, Площадь: %.2f, Периметр: %.2f\n", s, s.Area(), s.Perimeter())
}

func main() {
	c := &Circle{10.3}
	r := &Rectangle{23.54, 54.3}
	t := &Triangle{54.2, 76.4, 67.4}

	PrintShapeInfo(c)
	PrintShapeInfo(r)
	PrintShapeInfo(t)
}
