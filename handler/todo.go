package handler

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/sin392/todo-back/model"
	"github.com/sin392/todo-back/service"
	"github.com/sin392/todo-back/utils"
)

type TODOHandler struct {
	svc *service.TODOService
}

func NewTODOHandler(svc *service.TODOService) *TODOHandler {
	return &TODOHandler{
		svc: svc,
	}
}

func (h *TODOHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	ctx := r.Context()
	switch r.Method {
	case "POST":
		todos, _ := h.createTODO(ctx, r)
		fmt.Fprintln(w, todos)
	case "GET":
		h.readTODO(ctx, r)
	case "PUT":
		h.updateTODO(ctx, r)
	case "DELETE":
		h.deleteTODO(ctx, r)
	}
}

func (h *TODOHandler) createTODO(ctx context.Context, r *http.Request) (*model.CreateTODOResponse, error) {
	var req *model.CreateTODORequest
	err := utils.DecodeBody(r, &req)
	if err != nil {
		log.Println(err)
	}
	// fmt.Fprintln(w, "create")
	// res := &model.TODO{
	// 	ID:          1,
	// 	Subject:     req.Subject,
	// 	Description: req.Description,
	// 	CreatedAt:   time.Now(),
	// 	UpdatedAt:   time.Now(),
	// }
	todo, err := h.svc.CreateTODO(ctx, req.Subject, req.Description)
	println(todo)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &model.CreateTODOResponse{TODO: *todo}, err
}

func (h *TODOHandler) readTODO(ctx context.Context, r *http.Request) {
	var req *model.ReadTODORequest
	utils.DecodeBody(r, &req)
}
func (h *TODOHandler) updateTODO(ctx context.Context, r *http.Request) {
	var req *model.UpdateTODORequest
	utils.DecodeBody(r, &req)
}
func (h *TODOHandler) deleteTODO(ctx context.Context, r *http.Request) {
	var req *model.DeleteTODORequest
	utils.DecodeBody(r, &req)
}
