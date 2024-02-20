package api

import (
	"Task36a.4.1Aggregator/pkg/storage"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// API Программный интерфейс сервера.
type API struct {
	db     storage.Interface // БД.
	router *mux.Router       // Маршрутизатор.
}

// New Конструктор объекта API.
func New(db storage.Interface) *API {
	api := API{db: db}
	api.router = mux.NewRouter()
	api.endpoints()

	return &api
}

// Регистрирует объекты API.
func (api *API) endpoints() {
	// Получить n последних новостей
	api.router.HandleFunc("/news/{n}", api.postsHandler).Methods(http.MethodGet, http.MethodOptions)
	// Веб-приложение
	api.router.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir("./webapp"))))
}

// Router Получение маршрутизатора запросов.
// Требуется для передачи маршрутизатора веб-серверу.
func (api *API) Router() *mux.Router {
	
	return api.Router()
}

// Возвращает все новости.
func (api *API) postsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		return
	}
	s := mux.Vars(r)["n"]
	n, _ := strconv.Atoi(s)
	news, err := api.db.Posts(n)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(news)
}
