package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"

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
	enc := json.NewEncoder(w)
	ctx := r.Context()
	switch r.Method {
	case "POST":
		res, err := h.createTODO(ctx, r)
		if err != nil {
			log.Fatal(err)
		}
		enc.Encode(&res)
	case "GET":
		var res interface{}
		var err error
		if r.URL.Query().Get("id") != "" {
			res, err = h.ReadTODO(ctx, r)
		} else {
			res, err = h.ReadTODOs(ctx, r)
		}
		if err != nil {
			log.Fatal(err)
		}
		enc.Encode(&res)
	case "PUT":
		h.updateTODO(ctx, r)
	case "DELETE":
		var err error
		if r.URL.Query().Get("ids") != "" {
			err = h.deleteTODOs(ctx, r)
		} else {
			err = h.deleteTODO(ctx, r)
		}
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (h *TODOHandler) createTODO(ctx context.Context, r *http.Request) (*model.TODO, error) {
	var req *model.CreateTODORequest
	err := utils.DecodeBody(r, &req)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	res, err := h.svc.CreateTODO(ctx, &model.CreateTODORequest{Subject: req.Subject, Description: req.Description})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return res, err
}

func (h *TODOHandler) ReadTODO(ctx context.Context, r *http.Request) (*model.TODO, error) {
	id := r.URL.Query().Get("id")

	res, err := h.svc.ReadTODO(ctx, &model.ReadTODORequest{ID: id})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return res, err
}

func (h *TODOHandler) ReadTODOs(ctx context.Context, r *http.Request) (*[]model.TODO, error) {
	skip := r.URL.Query().Get("skip")
	limit := r.URL.Query().Get("limit")

	res, err := h.svc.ReadTODOs(ctx, &model.ReadTODOsRequest{Skip: &skip, Limit: &limit})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return res, err
}

func (h *TODOHandler) updateTODO(ctx context.Context, r *http.Request) (*model.TODO, error) {
	var req *model.UpdateTODORequest
	utils.DecodeBody(r, &req)
	res, err := h.svc.UpdateTODO(ctx, &model.UpdateTODORequest{ID: req.ID})
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return res, err
}

func (h *TODOHandler) deleteTODO(ctx context.Context, r *http.Request) error {
	// var req *model.DeleteTODORequest
	// utils.DecodeBody(r, &req)
	id := r.URL.Query().Get("id")
	err := h.svc.DeleteTODO(ctx, &model.DeleteTODORequest{ID: id})
	if err != nil {
		log.Fatal(err)
	}
	return err
}

func (h *TODOHandler) deleteTODOs(ctx context.Context, r *http.Request) error {
	// var req *model.DeleteTODORequest
	// utils.DecodeBody(r, &req)
	ids := strings.Split(strings.Replace(r.URL.Query().Get("ids"), " ", "", -1), ",")
	err := h.svc.DeleteTODOs(ctx, &model.DeleteTODOsRequest{IDs: ids})
	if err != nil {
		log.Fatal(err)
	}
	return err
}
