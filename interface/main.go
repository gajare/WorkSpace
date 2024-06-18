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

type Triangle struct {
	Base, Height float64
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
func PrintArea(s Shape) {
	fmt.Println("area of:", s.Area())

}
func main() {
	rect := Rectangle{Width: 10, Height: 5}
	tri := Triangle{Base: 6, Height: 4}
	PrintArea(rect)
	PrintArea(tri)

}
