package main

import (
	"fmt"
	"math"
	"reflect"
)

// interface use for fix Polymorphism
func main() {
	r := rectangle{width: 10, height: 10}
	fmt.Println(getArea(r))
	c := circle{10}
	fmt.Println(getArea(c))
	showInfo(c)
	printFields(r)
	castToRectangle(c)
}

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width, height float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s shape) float64 {
	return s.area()
}

// check Type
func showInfo(s shape) {
	t := reflect.TypeOf(s).Name()
	switch t {
	case "rectangle":
		r := s.(rectangle)
		fmt.Printf("width: %v, height: %v\n", r.width, r.height)
	case "circle":
		c := s.(circle)
		fmt.Printf("radius: %v\n", c.radius)
	}
}

func printFields(s shape) {
	v := reflect.ValueOf(s)
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		name := t.Field(i).Name
		var value interface{}
		if field.Kind() == reflect.Interface {
			value = field.Interface()
		} else {
			value = field
		}
		fmt.Printf("%s: %v\n", name, value)
	}
}

func castToRectangle(s shape) {
	r, ok := s.(rectangle) // _ is data casting if not use you must use _
	if !ok {
		fmt.Println("Casting Error")
	} else {
		fmt.Printf("Casting Success w: %f,h:%f", r.width, r.height)
	}
}
