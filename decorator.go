package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

//IDShouldBeInt ...
func IDShouldBeInt(h httprouter.Handle, name string) httprouter.Handle {
	return CommonHeaders(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		idParam := ps.ByName("todoId")
		_, err := strconv.Atoi(idParam)
		if err != nil {
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusInternalServerError)
			if errr := json.NewEncoder(w).Encode(err); errr != nil {
				return
			}
			return
		}

		h(w, r, ps)
	}, name)
}

//CommonHeaders ....
func CommonHeaders(h httprouter.Handle, name string) httprouter.Handle {
	return Logging(func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		h(w, r, ps)
	}, name)
}
