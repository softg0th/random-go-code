package api

import "crud/internal/storage"

type Handler struct {
	S storage.Storer
}
