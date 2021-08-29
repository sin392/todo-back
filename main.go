package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sin392/todo-back/handler"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	http.Handle("/edit", handler.NewTODOHandler())

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("ListenAndServe failed.", err)
	}
}
