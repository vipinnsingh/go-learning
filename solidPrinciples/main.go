package main

import "fmt"

type Shape interface {
	Area() int
}

type Rectangle struct {
	Length  int
	Breadth int
}

func (r *Rectangle) Area() int {
	return r.Length * r.Breadth
}

type Square struct {
	Side int
}

func (s *Square) Area() int {
	return s.Side * s.Side
}

func PrintArea(s Shape) {
	fmt.Println(s.Area())
}

func mainn() {

	r := Rectangle{
		Length:  2,
		Breadth: 5,
	}

	s := Square{
		Side: 6,
	}

	PrintArea(&r)
	PrintArea(&s)

}
