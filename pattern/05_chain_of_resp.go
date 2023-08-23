package main

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
	Паттерн цепочка вызовов относится к паттернам поведенческого типа.
	Этот паттерн позволяет передеавать запросы последовательно по цепочке обработчиков
	Каждый из обработичков решает может ли он обработать запрос сам и стоит ли передавать запрос дальше
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

type Handler interface {
	Send()
}

type Handler1 struct {
	next Handler
}

func (h *Handler1) Send() {
	fmt.Println("Handler 1")
	if h.next != nil {
		h.next.Send()
	}
}

type Handler2 struct {
	next Handler
}

func (h *Handler2) Send() {
	fmt.Println("Handler 2")
	if h.next != nil {
		h.next.Send()
	}
}

type Handler3 struct {
	next Handler
}

func (h *Handler3) Send() {
	fmt.Println("Handler 3")
}

//func main() {
//	handlers := &Handler1{next: &Handler2{next: &Handler3{}}}
//	handlers.Send()
//}
