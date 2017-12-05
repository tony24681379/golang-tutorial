package main

import "fmt"

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

func main() {
	tony := Human{"Tony", 24, "0912345678"}
	mark := Student{Human{"Mark", 21, "0912345678"}, "NTHU"}
	sam := Employee{Human{"Sam", 21, "0912345678"}, "google"}
	tony.SayHi()
	mark.SayHi()
	sam.SayHi()
}
