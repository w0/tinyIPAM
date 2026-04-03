package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /{$}", app.home)
	mux.HandleFunc("POST /api/v1/subnet", app.newSubnet)
	mux.HandleFunc("GET /api/v1/subnet/{id}", app.getSubnet)

	return mux
}
