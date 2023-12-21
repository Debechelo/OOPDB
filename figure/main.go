package main

import (
	"fmt"
	"math"
)

// В golang не надо специально указывать,
// что структуры применяют определенный интерфейс.
// В Go используется правило регистра первой буквы имени —
// если название начинается заглавной буквы — это public-доступ,
// если со строчной — private

type GeometricFigure interface {
	calculateArea() float64
	calculatePerimeter() float64
}

// Все структуры наследуют интерфейс GeometricFigure
// Square наследует Rectangle

// Circle //////////////////////
type Circle struct {
	radius float64
}

func NewCircle(radius float64) *Circle {
	return &Circle{radius: radius}
}

func (c *Circle) calculateArea() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *Circle) calculatePerimeter() float64 {
	return 2 * math.Pi * c.radius
}

////////////////////////////////

// Rectangle ///////////////////
type Rectangle struct {
	length float64
	width  float64
}

func NewRectangle(length, width float64) *Rectangle {
	return &Rectangle{length: length, width: width}
}

func (r *Rectangle) calculateArea() float64 {
	return r.length * r.width
}

func (r *Rectangle) calculatePerimeter() float64 {
	return 2 * (r.length + r.width)
}

////////////////////////////////

// Square //////////////////////
type Square struct {
	side float64
	Rectangle
}

func NewSquare(side float64) *Square {
	return &Square{side: side, Rectangle: Rectangle{length: side, width: side}}
}

////////////////////////////////

func main() {
	circle := NewCircle(5)
	fmt.Printf("Circle Area: %.2f\n", circle.calculateArea())
	fmt.Printf("Circle Perimeter: %.2f\n", circle.calculatePerimeter())

	rectangle := NewRectangle(5, 3)
	fmt.Printf("Rectangle Area: %.2f\n", rectangle.calculateArea())
	fmt.Printf("Rectangle Perimeter: %.2f\n", rectangle.calculatePerimeter())

	square := NewSquare(3)
	fmt.Printf("Square Area: %.2f\n", square.calculateArea())
	fmt.Printf("Square Perimeter: %.2f\n", square.calculatePerimeter())
}
