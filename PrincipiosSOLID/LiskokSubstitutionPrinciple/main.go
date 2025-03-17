//Los objetos deben ser reemplazables por instancias de sus subtipos sin alterar el correcto funcionamiento del programa.

package main

import "fmt"

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width, Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func PrintArea(shape Shape) {
	fmt.Println("Area:", shape.Area())
}

func main() {
	r := Rectangle{Width: 5, Height: 10}
	s := Square{Side: 7}

	PrintArea(r)
	PrintArea(s)
}
