package api

import (
	"in-memory-api/handlers"
	"net/http"
)

func RegisterRoutes(artistHandler *handlers.ArtistHandler) {
	http.HandleFunc("/artists", artistHandler.GetArtists)
}
