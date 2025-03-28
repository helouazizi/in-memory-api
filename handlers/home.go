package handlers

import (
	"in-memory-api/store"
	"net/http"
)

type HomeHandler struct {
	Store *store.Memory_Store
}

func (h *HomeHandler) HomePage(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var artists []struct {
		ID    string
		Name  string
		Image string
	}

	for _, artist := range h.Store.Artists {
		artists = append(artists)
	}
}
