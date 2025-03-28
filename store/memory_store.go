package store

import (
	"in-memory-api/models"
	"sync"
)

type Memory_Store struct {
	Artists   []models.Artist
	Loactions []models.Location
	Dates     []models.Date
	Relations []models.Relation
	Mutex     sync.Mutex
}

func NewMemoryStore() *Memory_Store {
	return &Memory_Store{}
}
