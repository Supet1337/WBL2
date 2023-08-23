package main

/*
	Реализовать паттерн «строитель».
	Паттерн строитель это порождающий паттерн, который позволяет создавать сложные объекты с разынм количеством параметров шаг за шагом
	https://en.wikipedia.org/wiki/Builder_pattern
*/

type Builder interface {
	MakeWheels()
	MakeBody()
	MakeInterrior()
	GetCar() Car
}

func getBuilder(carType string) Builder {
	switch carType {
	case "Offroad":
		return NewOffroadCarBuilder()
	case "Sport":
		return NewSportCarBuilder()
	}
	return nil
}

type Car struct {
	wheels    string
	body      string
	interrior string
}

type SportCarBuilder struct {
	wheels    string
	body      string
	interrior string
}

func NewSportCarBuilder() *SportCarBuilder {
	return &SportCarBuilder{}
}

func (s *SportCarBuilder) MakeWheels() {
	s.wheels = "Sport Wheels"
}

func (s *SportCarBuilder) MakeBody() {
	s.body = "Sport Body"
}

func (s *SportCarBuilder) MakeInterrior() {
	s.interrior = "Sport Interrior"
}

func (s *SportCarBuilder) GetCar() Car {
	return Car{
		wheels:    s.wheels,
		body:      s.body,
		interrior: s.interrior,
	}
}

type OffroadCarBuilder struct {
	wheels    string
	body      string
	interrior string
}

func NewOffroadCarBuilder() *OffroadCarBuilder {
	return &OffroadCarBuilder{}
}

func (o *OffroadCarBuilder) MakeWheels() {
	o.wheels = "Offroad Wheels"
}

func (o *OffroadCarBuilder) MakeBody() {
	o.body = "Offroad Body"
}

func (o *OffroadCarBuilder) MakeInterrior() {
	o.interrior = "Offroad Interrior"
}

func (o *OffroadCarBuilder) GetCar() Car {
	return Car{
		wheels:    o.wheels,
		body:      o.body,
		interrior: o.interrior,
	}
}

//
//func main() {
//	sportcar := getBuilder("Sport")
//	offroadcar := getBuilder("Offroad")
//
//	sportcar.MakeInterrior()
//	sportcar.MakeBody()
//	sportcar.MakeWheels()
//	fmt.Println(sportcar.GetCar())
//
//	offroadcar.MakeInterrior()
//	offroadcar.MakeBody()
//	offroadcar.MakeWheels()
//	fmt.Println(offroadcar.GetCar())
//}
