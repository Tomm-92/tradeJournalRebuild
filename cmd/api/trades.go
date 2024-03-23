package main

import (
	"fmt"
	"net/http"
)

func (app *application) createTradeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create a new trade entry")
}

func (app *application) getTradeHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "show the details of trade %d\n", id)
}
