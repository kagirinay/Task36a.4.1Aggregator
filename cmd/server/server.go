package main

import (
	"Task36a.4.1Aggregator/pkg/api"
	"Task36a.4.1Aggregator/pkg/storage"
)

// Сервер.
type server struct {
	db  storage.Interface
	api *api.API
}

// Конфигурация приложения.
type config struct {
	Period  int      `json:"request_period"`
	LinkArr []string `json:"rss"`
}
