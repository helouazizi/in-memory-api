package handlers

import (
	"encoding/json"
	"in-memory-api/store"
	"net/http"
)

type ArtistHandler struct {
	Store *store.Memory_Store
}

func (h *ArtistHandler) GetArtists(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.Store.Artists)
}
