package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

//Index ...
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprintf(w, "Welcome")
}

//TodoIndex ...
func TodoIndex(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

//TodoShow ...
func TodoShow(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("todoId"))
	t := RepoFindTodo(id)
	if t.ID == 0 && t.Name == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

//TodoCreate ...
func TodoCreate(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var todo Todo

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()

	if err := json.Unmarshal(body, &todo); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		if errr := json.NewEncoder(w).Encode(err); errr != nil {
			panic(errr)
		}
		return
	}

	t := RepoCreateTodo(todo)
	location := fmt.Sprintf("http://%s/%d", r.Host, t.ID)
	w.Header().Set("Location", location)
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err)
	}
}

//TodoDelete ...
func TodoDelete(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, _ := strconv.Atoi(ps.ByName("todoId"))

	if err := RepoDestroyTodo(id); err != nil {
		w.WriteHeader(http.StatusNotFound)
		if errr := json.NewEncoder(w).Encode(err); errr != nil {
			panic(errr)
		}
		return
	}

	w.Header().Del("Content-Type")
	w.WriteHeader(http.StatusNoContent)
}
