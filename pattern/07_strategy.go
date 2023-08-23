package main

import "fmt"

/*
	Реализовать паттерн «стратегия».
	Паттерн стратегия является поведенческим паттерном, который определяет семейство алгоритмов и помещает каждый из них в собственный класс
	После чего эти стратегии можно взаимозаменять
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

type Strategy interface {
	calculate(a, b int)
}

type Sum struct{}

func (s *Sum) calculate(a, b int) {
	fmt.Println(a + b)
}

type Diff struct{}

func (d *Diff) calculate(a, b int) {
	fmt.Println(a - b)
}

type Calculator struct {
	operation Strategy
}

func (c *Calculator) setOps(ops Strategy) {
	c.operation = ops
}

func (c *Calculator) execute(a, b int) {
	c.operation.calculate(a, b)
}

//func main() {
//	calc := Calculator{}
//	calc.setOps(&Sum{})
//	calc.execute(1, 2)
//
//	calc.setOps(&Diff{})
//	calc.execute(2, 1)
//}
