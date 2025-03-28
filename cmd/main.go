package main

import (
    "log"
    "net/http"
    "in-memory-api/api"
    "in-memory-api/handlers"
    "in-memory-api/store"
)

func main() {
    store := store.NewMemoryStore()
    store.LoadData() // Fetch API data

    artistHandler := &handlers.ArtistHandler{Store: store}

    api.RegisterRoutes(artistHandler)

    log.Println("Server running on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
