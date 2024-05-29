package main

import (
	"crud/internal/api"
	"crud/internal/server"
	"crud/internal/storage"
)

func main() {
	h := api.Handler{
		S: storage.New(),
	}

	if err != nil {
		return
	}

	storage.InitDatabase(db)

	server := server.New()
}
