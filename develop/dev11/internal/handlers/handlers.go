package handlers

import (
	"dev11/internal/middleware"
	"dev11/internal/usecase"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Server struct {
	Server  *http.Server
	usecase usecase.IUseCase
}

func CreateListenServer(port string) error {
	server := Server{
		Server:  &http.Server{Addr: port},
		usecase: usecase.NewUseCase(),
	}
	server.Server.Handler = server.routes()
	fmt.Println("Starting server at port", port)
	err := server.Server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func (s *Server) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/events_for_day", middleware.Log(s.Get))
	mux.HandleFunc("/events_for_week", middleware.Log(s.Get))
	mux.HandleFunc("/events_for_month", middleware.Log(s.Get))
	mux.HandleFunc("/create_event", middleware.Log(s.Create))
	mux.HandleFunc("/update_event", middleware.Log(s.Update))
	mux.HandleFunc("/delete_event", middleware.Log(s.Delete))
	//mux.HandleFunc("/snippet/create", s.createSnippet)

	return mux
}

func (s *Server) Get(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode("Unable method")
		return
	}
	var result []byte
	var err error
	switch r.URL.Path {
	case "/events_for_day":
		result, err = s.usecase.Get("day")
	case "/events_for_week":
		result, err = s.usecase.Get("week")
	case "/events_for_month":
		result, err = s.usecase.Get("month")
	default:
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Unable method")
		return
	}

	if err != nil {
		w.WriteHeader(503)
		json.NewEncoder(w).Encode("error: " + err.Error())
	}
	w.Write(result)
}

func (s *Server) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode("Unable method")
		return
	}
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Unable to read body: " + err.Error())
		return
	}

	err = s.usecase.Create(requestBody)
	if err != nil {
		w.WriteHeader(503)
		json.NewEncoder(w).Encode("error: " + err.Error())
		return
	}
	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode("Event was created")
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode("Unable method")
		return
	}
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Unable to read body: " + err.Error())
		return
	}

	err = s.usecase.Update(requestBody)
	if err != nil {
		w.WriteHeader(503)
		json.NewEncoder(w).Encode("error: " + err.Error())
		return
	}
	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode("Event was update")
	if err != nil {
		log.Fatal(err)
	}
}

func (s *Server) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(500)
		json.NewEncoder(w).Encode("Unable method")
		return
	}
	requestBody, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode("Unable to read body: " + err.Error())
		return
	}

	err = s.usecase.Delete(requestBody)
	if err != nil {
		w.WriteHeader(503)
		json.NewEncoder(w).Encode("error: " + err.Error())
		return
	}
	w.WriteHeader(200)
	err = json.NewEncoder(w).Encode("Event was delete")
	if err != nil {
		log.Fatal(err)
	}
}
