package main

import "fmt"

/*
	Реализовать паттерн «фасад».
	Фасад - структурный паттерн, который позволяет работать с большим количеством
	объектов сложного фреймворка\библиотеки, используя, при этом всего 1 простой интерфейс
*/

type TeaFacade struct {
	water *Water
	sugar *Sugar
	tea   *Tea
}

type Water struct {
	temperature float32
}

func (w *Water) boilWater() {
	fmt.Println("Boiling water... Temp:", w.temperature)
	w.temperature = 90
	fmt.Println("Boiling water... Temp:", w.temperature)

}

type Sugar struct {
	sugarType string
	count     int
}

func (s *Sugar) addSugar() {
	fmt.Println("add sugar:", s.sugarType, "Count of", s.count)
}

type Tea struct {
	teaType string
}

func (t *Tea) brewTea() {
	fmt.Println("Brew Tea", t.teaType)
}

func newTeaFacade(typeOfTea string, sugarCount int, sugarType string) *TeaFacade {
	fmt.Println("Starting create tea!")
	teafacade := &TeaFacade{
		water: &Water{temperature: 0},
		sugar: &Sugar{
			sugarType: sugarType,
			count:     sugarCount,
		},
		tea: &Tea{teaType: typeOfTea},
	}
	fmt.Println("Tea created")
	return teafacade
}

func (t *TeaFacade) LetsTeaFacade() {
	fmt.Println("Lets boil the water")
	t.water.boilWater()
	fmt.Println("Temperature is ok")
	fmt.Println("Lets brew the tea")
	t.tea.brewTea()
	fmt.Println("Lets add sugar")
	t.sugar.addSugar()
	fmt.Println("Lets Drink!!!")
}

//
//func main() {
//	teafacade := newTeaFacade("Earl Gray", 1, "Classic")
//	teafacade.LetsTeaFacade()
//}
