package practice

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
}

type Rect struct{
	width, height float64
}

type Circle struct {
	radius float64
}

func (r Rect) area() float64{
	return r.width * r.height
}

func (c Circle) area() float64{
	return math.Pi * c.radius * c.radius
}

func paramInterface(i interface{}){
	fmt.Println(i)
}

func InterfacePractice(){
	fmt.Println(Rect{10,20}.area())
	fmt.Println(Circle{10}.area())
	
	var x interface{}
	x = 1
	paramInterface(x)
	x = "test"
	paramInterface(x)
}