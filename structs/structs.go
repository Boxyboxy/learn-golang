package main

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

type Circle struct {
	Radius float64
}

func (circle Circle) Perimeter() float64 {
	return 2 * math.Pi * circle.Radius
}
func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

type Triangle struct {
	Height float64
	Base   float64
}

func (triangle Triangle) Area() float64 {
	return 0.5 * triangle.Height * triangle.Base
}

type Shape interface {
	Area() float64
}

// introduction to interfaces
// interface resolution is IMPLICIT
/*
Rectangle has a method called Area that returns a float64 so it satisfies the Shape interface

Circle has a method called Area that returns a float64 so it satisfies the Shape interface

string does not have such a method, so it doesn't satisfy the interface
*/
