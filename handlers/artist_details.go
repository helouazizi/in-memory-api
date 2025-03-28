package handlers

import (
	"encoding/json"
	"in-memory-api/memory"
	"net/http"
)

type ArtistHandler struct {
	Store *memory.Memory_Store
}

func (h *ArtistHandler) GetArtists_by_id(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(h.Store.Artists)
}
