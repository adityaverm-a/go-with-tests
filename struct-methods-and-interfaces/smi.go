package smi

import "math"

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * (t.Base * t.Height)
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

// Quick Summary

// - Declaring structs to create your own data types which lets you bundle related data together and make the intent of your code clearer

// - Declaring interfaces so you can define functions that can be used by different types

// - Adding methods so you can add functionality to your data types and so you can implement interfaces

// - Table driven tests to make your assertions clearer and your test suites easier to extend & maintain
