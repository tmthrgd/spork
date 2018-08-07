package main

import "net/http"

func httpHandlerError(handler func(http.ResponseWriter, *http.Request) error) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := handler(w, r)
		if err == nil {
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
