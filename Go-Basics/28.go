package main

import "fmt"

type Square struct {
	length float64
	width  float64
}

func (s *Square) area() float64 {
	return s.length * s.width
}

func (s *Square) parameter() float64 {
	perimeter := s.length + s.width
	return 2 * perimeter
}

func main() {

	s := Square{
		length: 3,
		width:  4,
	}

	fmt.Println(s.area())
	fmt.Println(s.parameter())

}
