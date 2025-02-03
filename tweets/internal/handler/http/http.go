package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/stevenstr/tweets_app_reforged/tweets/internal/controller/tweets"
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
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) // 400
		return
	}

	switch req.Method {
	case http.MethodGet:
		v, err := h.ctrl.Get(req.Context(), id)
		if err != nil {
			if errors.Is(err, tweets.ErrNotFound) {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound) // 404
				return
			}
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}

		if err := json.NewEncoder(w).Encode(v); err != nil {
			log.Printf("Response encode error: %v\n", err)
		}
	default:
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}

func (h *Handler) HandleDeleteSingleTweet(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	if id == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest) // 400
		return
	}

	switch req.Method {
	case http.MethodDelete:
		if err := h.ctrl.Delete(req.Context(), id); err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		}
	default:
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}

func (h *Handler) HandlePostSingleTweet(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	if id == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	switch req.Method {
	case http.MethodPost:
		if err := h.ctrl.Post(req.Context(), req.FormValue("id"), req.FormValue("message")); err != nil {
			log.Printf("Repository post error: %v\n", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	default:
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}

func (h *Handler) HandlePutSingleTweet(w http.ResponseWriter, req *http.Request) {
	id := req.FormValue("id")
	if id == "" {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	switch req.Method {
	case http.MethodPut:
		if err := h.ctrl.Put(req.Context(), req.FormValue("id"), req.FormValue("message")); err != nil {
			log.Printf("Repository put error: %v\n", err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	default:
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}

func (h *Handler) HandleGetAllTweet(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		v, err := h.ctrl.GetAll(req.Context())
		if err != nil && errors.Is(err, tweets.ErrNotFound) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		if err := json.NewEncoder(w).Encode(v); err != nil {
			log.Printf("Response encode error: %v\n", err)
		}
	default:
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
	}
}

func (h *Handler) HandleTime(w http.ResponseWriter, req *http.Request) {
	t := time.Now().Format(time.RFC1123)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		log.Printf("Response encode error: %v\n", err)
	}
}
