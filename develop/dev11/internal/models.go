package internal

import "time"

type Event struct {
	Id   int       `json:"id"`
	Name string    `json:"name"`
	Info string    `json:"info"`
	Date time.Time `json:"date"`
}

//type Response struct {
//	Error string `json:"error"`
//	Event Event  `json:"model"`
//}
