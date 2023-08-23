package main

import "fmt"

/*
	Реализовать паттерн «посетитель».
	Паттерн посетитель позволяет добавить новую операцию для иерархии классов, не изменяя код этих классов
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

type Animal interface {
	getName()
	accept(Visitor)
}

type Dog struct{}

func (d *Dog) getName() {
	fmt.Println("Dog")
}

func (d *Dog) accept(v Visitor) {
	v.visitDog(d)
}

type Cat struct{}

func (c *Cat) getName() {
	fmt.Println("Cat")
}

func (c *Cat) accept(v Visitor) {
	v.visitCat(c)
}

type Visitor interface {
	visitDog(*Dog)
	visitCat(*Cat)
}

type AnimalSound struct { //New feature
	sound string
}

func (a *AnimalSound) visitDog(d *Dog) {
	a.sound = "Wof"
	fmt.Println("Dog say:", a.sound)
}

func (a *AnimalSound) visitCat(c *Cat) {
	a.sound = "Meow"
	fmt.Println("Cat say:", a.sound)
}

//func main() {
//	dog := new(Dog)
//	cat := new(Cat)
//
//	AnimalSay := new(AnimalSound)
//
//	dog.getName()
//	dog.accept(AnimalSay)
//
//	cat.getName()
//	cat.accept(AnimalSay)
//}
