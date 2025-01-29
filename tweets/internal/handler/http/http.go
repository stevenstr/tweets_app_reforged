package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"tweets.com/tweets/internal/controller/tweets"
)

type Handler struct {
	ctrl *tweets.Controller
}

func New(ctrl *tweets.Controller) *Handler {
	return &Handler{ctrl: ctrl}
}

func (h *Handler) HandleGetSingleTweet(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	switch req.Method {
	case http.MethodGet:
		v, err := h.ctrl.Get(req.Context(), id)
		if err != nil && errors.Is(err, tweets.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		} else if err != nil {
			w.WriteHeader(http.StatusNotFound)
		}

		if err := json.NewEncoder(w).Encode(v); err != nil {
			log.Printf("Response encode error: %v\n", err)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (h *Handler) HandlePutSingleTweet(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest) // 400
		return
	}

	switch req.Method {
	case http.MethodPut:
		if err := h.ctrl.Put(req.Context(), req.FormValue("id"), req.FormValue("message")); err != nil {
			log.Printf("Repository pu error: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError) // 500
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}

func (h *Handler) HandleGetAllTweet(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		v, err := h.ctrl.GetAll(req.Context())
		if err != nil && errors.Is(err, tweets.ErrNotFound) {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if err := json.NewEncoder(w).Encode(v); err != nil {
			log.Printf("Response encode error: %v\n", err)
		}
	default:
		w.WriteHeader(http.StatusBadRequest)
	}
}
