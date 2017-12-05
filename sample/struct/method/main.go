package main

import "fmt"

type Rectangle struct {
	width, height float64
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func main() {
	r1 := Rectangle{12, 2}
	fmt.Println(r1.area())
}
