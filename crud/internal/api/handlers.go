package api

import (
	"crud/internal/storage"
	"encoding/json"
	"net/http"
)

func (h *Handler) pingHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (h *Handler) instertPostHandler(w http.ResponseWriter, r *http.Request) {
	var postdoc storage.PostDocument
	db := storage.GetDBInstance()

	json.NewDecoder(r.Body).Decode(&postdoc)

	err := h.S.InsertPost(db, postdoc)

	if err != nil {
		http.Error(w, "Failed to inmsert doc", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
