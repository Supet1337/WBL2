package main

import "fmt"

/*
	Реализовать паттерн «фабричный метод».
	Это порождающий паттерн, который позволяет определить общий интерфейс для создания объектов в спуер классе,
	Позволяя подклассам изменять тип создаваемых объектов
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

type Gadget interface {
	getName()
}

type Phone struct {
	name string
}

func (p *Phone) getName() {
	fmt.Println(p.name)
}

func NewPhone() Gadget {
	return &Phone{name: "Phone gadget"}
}

type PC struct {
	name string
}

func (p *PC) getName() {
	fmt.Println(p.name)
}
func NewPc() Gadget {
	return &PC{name: "Pc gadget"}
}

func GetGadget(name string) Gadget {
	switch name {
	case "phone":
		return NewPhone()
	case "computer":
		{
			return NewPc()
		}
	}
	return nil
}

//func main() {
//	phone := GetGadget("phone")
//	computer := GetGadget("computer")
//
//	phone.getName()
//	computer.getName()
//}
