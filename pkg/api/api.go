package api

import (
	"Task36a.4.1Aggregator/pkg/storage"
	"github.com/gorilla/mux"
)

// API Программный интерфейс сервера.
type API struct {
	db     storage.Post
	router *mux.Router
}
