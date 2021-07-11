package main 

import (
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"davidlares/timezone/api/handlers"
	"davidlares/timezone/api/repository"
)

func main() {
	
	// new repository
	repo := repository.NewRepository("mongodb://localhost:27017", "timezones", "timezones")
	
	// closing connection
	defer repo.Close() 

	// instance
	h := handlers.Handlers {  
		Repo: repo,
	}
	
	// mux instance
	r := mux.NewRouter()
	r.HandleFunc("/timezones", h.All).Methods("GET")
	r.HandleFunc("/timezones/{timezone}", h.GetByTimezone).Methods("GET")
	r.HandleFunc("/timezones", h.Insert).Methods("POST")
	r.HandleFunc("/timezones/{timezone}", h.Delete).Methods("DELETE")
	r.HandleFunc("/timezones/{timezone}", h.Update).Methods("PATCH")
	
	// logging
	log.Fatal(http.ListenAndServe(":9000", r))
}