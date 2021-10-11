package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/sin392/todo-back/handler"
	"github.com/sin392/todo-back/service"
	"github.com/sin392/todo-back/utils"
)

func main() {
	db := utils.GetConnection()
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
	http.Handle("/todo", handler.NewTODOHandler(service.NewTODOService(db)))

	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal("ListenAndServe failed.", err)
	}
}
