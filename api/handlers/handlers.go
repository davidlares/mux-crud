package handlers 

import (
 "encoding/json"
 "net/http"
 "github.com/gorilla/mux"
 "davidlares/timezone/api/repository"
)

type Handlers struct {
	Repo *repository.Repository
}

func (h *Handlers) All(w http.ResponseWriter, r *http.Request) {
	// getting all timezones
	tzcs, err := h.Repo.FindAll()
	if err != nil {
		error500(w, err)
		return
	}
	// marshaling ORM results into json
	jr, err := json.Marshal(tzcs)
	// handling error
	if err != nil {
		error500(w, err)
		return
	}
	// converting into string
	ok200(w, string(jr))
}

func (h *Handlers) GetByTimezone(w http.ResponseWriter, r *http.Request) {
	// getting queryString param timezone
	params := mux.Vars(r)
	// timezone
	tz, ok := params["timezone"]
	if !ok {
		error400(w, "Timezone required")
		return
	}
	// getting record
	tzc, err := h.Repo.FindByTimezone(tz)
	if err != nil {
		error404(w, "Timezone not found")
	}
	// marshaling ORM results
	jr, err := json.Marshal(tzc)
	// handling error
	if err != nil {
		error500(w, err)
		return
	}
	// converting into string
	ok200(w, string(jr))
}


func (h *Handlers) Delete(w http.ResponseWriter, r *http.Request) {
	// deleting queryString Param timezone
	params := mux.Vars(r)
	tz, ok := params["timezone"]
	if !ok {
		error400(w, "Timezone is required")
		return
	}
	// asking for convertion
	tzc := repository.TZConvertion {
		TimeZone: tz,
	}
	// handling error
	err := h.Repo.Delete(tzc)
	if err != nil {
		error500(w, err)
		return
	}
	// sending response
	ok200(w, "Element deleted")
}

func (h *Handlers) Insert(w http.ResponseWriter, r *http.Request) {
	// inserting timezone
	defer r.Body.Close() // execute last
	// instance
	var tzc repository.TZConvertion
	err := json.NewDecoder(r.Body).Decode(&tzc)
	if err != nil {
		error400(w, "Invalid json")
		return
	}
	// inserting json
	err = h.Repo.Insert(tzc)
	// handling error
	if err != nil {
		error500(w, err)
		return
	}
	// sending response
	ok200(w, "Element Successfully inserted")
}

func (h *Handlers) Update(w http.ResponseWriter, r *http.Request) {
	// updating queryString param timezone
	params := mux.Vars(r)
	tz, ok := params["timezone"]
	if !ok {
		error400(w, "Timezone is required")
		return
	}

	var tzc repository.TZConvertion
	err := json.NewDecoder(r.Body).Decode(&tzc)
	if err != nil {
		error400(w, "Invalid json")
		return
	}

	// handling errors
	err = h.Repo.Update(tz, tzc) 
	if err != nil {
		error500(w, err)
		return
	}
	// sending response
	ok200(w, "Element Successfully updated")
}