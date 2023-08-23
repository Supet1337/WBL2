package main

import "fmt"

/*
	Реализовать паттерн «комманда».
	паттерн команда позволяет заворачивать запросы или простые операции в отдельные объекты
	https://en.wikipedia.org/wiki/Command_pattern
*/

type Engine interface {
	engineOn()
	engineOff()
}

type CarEngine struct {
	isRunning bool
}

func (e *CarEngine) engineOn() {
	e.isRunning = true
	fmt.Println("Engine on")
}

func (e *CarEngine) engineOff() {
	e.isRunning = false
	fmt.Println("Engine off")
}

type Command interface {
	execute()
}

type Starter struct {
	command Command
}

func (s *Starter) press() {
	s.command.execute()
}

type StartComand struct {
	engine Engine
}

func (s *StartComand) execute() {
	s.engine.engineOn()
}

type StopComand struct {
	engine Engine
}

func (s *StopComand) execute() {
	s.engine.engineOff()
}

//func main() {
//	car := &CarEngine{}
//	startCommand := &StartComand{engine: car}
//	stopCommand := &StopComand{engine: car}
//
//	buttonOn := &Starter{startCommand}
//	buttonOn.press()
//
//	buttonOff := &Starter{stopCommand}
//	buttonOff.press()
//
//}
