package main

import (
	"in-memory-api/memory"
)

func main() {
	memory := memory.NewMemoryStore()
	memory.LoadData()
}
