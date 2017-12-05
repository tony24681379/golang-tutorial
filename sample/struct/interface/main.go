package main

import "fmt"

type writeGolang interface {
	SayHi()
	writeGo()
}

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s, you can call me on %s\n", h.name, h.phone)
}

//  Employee writeGo
func (s Employee) writeGo() {
	fmt.Printf("Hi, I am %s, I can write golang\n", s.name)
}

//  Student writeGo
func (s Student) writeGo() {
	fmt.Printf("Hi, I am %s, I can write golang\n", s.name)
}

func main() {
	mark := Student{Human{"Mark", 21, "0912345678"}, "NTHU"}
	sam := Employee{Human{"Sam", 21, "0912345678"}, "google"}
	var i writeGolang

	i = mark
	i.writeGo() //Hi, I am Mark, I can write golang
	i.SayHi()   //Hi, I am Mark, you can call me on 0912345678
	i = sam
	i.writeGo() //Hi, I am Sam, I can write golang
}
