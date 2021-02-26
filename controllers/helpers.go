package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
)

func parseForm(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}

	dec := schema.NewDecoder()
	err := dec.Decode(dst, r.PostForm)
	if err != nil {
		panic(err)
	}
	return nil
}
