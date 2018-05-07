package main

import (
	"fmt"
)

type Car interface {
	GetName() string
	Run()
	DiDi()
}

type BMW struct {
	Name string
}
type BYD struct {
	Name string
}

func (p *BMW) GetName() string {
	return p.Name
}

func (p *BMW) Run() {
	fmt.Printf("%s is running\n", p.Name)
}

func (p *BMW) DiDi() {
	fmt.Printf("%s is didi\n", p.Name)
}

func (p *BYD) GetName() string {
	return p.Name
}

func (p *BYD) Run() {
	fmt.Printf("%s is running\n", p.Name)
}

func (p *BYD) DiDi() {
	fmt.Printf("%s is didi\n", p.Name)
}

func main() {
	var car Car
	fmt.Println(car)

	byd := &BYD{
		Name: "BYD",
	}
	car = byd
	car.Run()
}
