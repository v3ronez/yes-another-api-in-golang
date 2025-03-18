package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

func (app *application) storeMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new movie")
}

func (app *application) showMovie(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseInt(httprouter.
		ParamsFromContext(r.Context()).
		ByName("id"), 10, 64)
	if err != nil || id < 1 {
		http.NotFound(w, r)
	}
	fmt.Fprintf(w, "show a movie %d", id)
}
