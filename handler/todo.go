package handler

import (
	"fmt"
	"net/http"
)

type TODOHandler struct{}

func NewTODOHandler() *TODOHandler {
	return &TODOHandler{}
}

func (h *TODOHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		h.createTODO(w, r)
	case "GET":
		h.readTODO(w, r)
	case "PUT":
		h.updateTODO(w, r)
	case "DELETE":
		h.deleteTODO(w, r)
	}
}

func (h *TODOHandler) createTODO(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "create")
}
func (h *TODOHandler) readTODO(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "read")
}
func (h *TODOHandler) updateTODO(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "update")
}
func (h *TODOHandler) deleteTODO(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "delete")
}
