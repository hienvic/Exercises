package main

import (
	"fmt"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprint(w, "Welcome!\n")
	})

	router.GET("/hello/:name", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
	})

	http.ListenAndServe(":8080", router)
}