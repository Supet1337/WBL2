package main

import "fmt"

/*
	Реализовать паттерн «состояние».
	Паттерн состояние является поведенческим паттерном, который позволяет объектам менять своё поведение в зависимости от состояния
	В отличии от паттерна Стратегия, состояния могут знать о друг-друге и переключаться между собой
	https://en.wikipedia.org/wiki/State_pattern
*/

type Tv struct {
	isOn  State
	isOff State

	currentState State
}

func newTv() *Tv {
	t := &Tv{}
	t.isOn = &OnState{tv: t}
	t.isOff = &OffState{tv: t}
	t.setState(t.isOff)
	return t
}

func (t *Tv) setState(s State) {
	t.currentState = s
}

func (t *Tv) changeChannel() {
	t.currentState.changeChannel()
}

func (t *Tv) onTV() {
	t.currentState.onTV()
}

func (t *Tv) offTV() {
	t.currentState.offTV()
}

type State interface {
	changeChannel()
	offTV()
	onTV()
}

type OnState struct {
	tv *Tv
}

func (o *OnState) changeChannel() {
	fmt.Println("Channel change success, because TV ON")
}

func (o *OnState) onTV() {
	fmt.Println("Can't ON tv, because TV already ON")
}

func (o *OnState) offTV() {
	fmt.Println("TV OFF")
	o.tv.setState(o.tv.isOff)
}

type OffState struct {
	tv *Tv
}

func (o *OffState) changeChannel() {
	fmt.Println("Channel change error, because TV OFF")
}

func (o *OffState) onTV() {
	fmt.Println("TV ON")
	o.tv.setState(o.tv.isOn)
}

func (o *OffState) offTV() {
	fmt.Println("Can't OFF tv, because TV already OFF")
}

//func main() {
//	tv := newTv()
//	tv.onTV()
//	tv.changeChannel()
//	tv.offTV()
//	tv.changeChannel()
//}
