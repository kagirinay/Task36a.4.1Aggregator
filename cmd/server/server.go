package main

import (
	"Task36a.4.1Aggregator/pkg/api"
	"Task36a.4.1Aggregator/pkg/rss"
	"Task36a.4.1Aggregator/pkg/storage"
	"Task36a.4.1Aggregator/pkg/storage/memdb"
	"Task36a.4.1Aggregator/pkg/storage/postgres"
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

// Сервер.
type server struct {
	db  storage.Interface
	api *api.API
}

// Конфигурация приложения.
type configJson struct {
	Period  int      `json:"request_period"`
	LinkArr []string `json:"rss"`
}

func main() {

	// Создаём объект сервера.
	var srv server

	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Создаём объекты баз данных.

	// БД заглушка
	db1 := memdb.New()
	_ = db1

	// Реляционная БД PostgresSQL.
	db2, err := postgres.New(ctx, "postgres://postgres:password@192.168.58.133:5432/news")
	if err != nil {
		log.Fatal(err)
	}
	_ = db2

	// Инициализируем хранилище выбранного сервера БД.
	srv.db = db2

	// Создаём объект API и регистрируем обработчики.
	srv.api = api.New(srv.db)

	// Создаём каналы для новостей и ошибок.
	chanPosts := make(chan []storage.Post)
	chanErr := make(chan error)
	var config configJson

	// Получение и парсинг ссылок из конфигурации.
	myLinks := receivingRSS("config.json", chanErr)
	for i := range myLinks.LinkArr {
		go parseNews(myLinks.LinkArr[i], chanErr, chanPosts, config.Period)
	}

	//Обработка потока новостей.
	go func() {
		for posts := range chanPosts {
			for i := range posts {
				err := db2.AddPost(posts[i])
				if err != nil {
					return
				}
			}
		}
	}()

	// Обработка потока ошибок.
	go func() {
		for err := range chanErr {
			log.Println("Ошибка:", err)
		}
	}()

	// Запуск веб сервера с API и приложением.
	err = http.ListenAndServe(":80", srv.api.Router())
	if err != nil {
		log.Fatal(err)
	}
}

// loadConfiguration Чтение и раскодирование файла конфигурации.
func loadConfiguration(file string) configJson {
	var config configJson
	configFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config
}

// Получает отдельные ссылки из конфигурации, ошибки направляет в поток ошибок.
func receivingRSS(fileName string, errors chan<- error) configJson {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		errors <- err
	}
	defer jsonFile.Close()

	byteValue, _ := io.ReadAll(jsonFile)

	var links configJson

	json.Unmarshal(byteValue, &links)

	return links
}

// Получает новости по ссылкам и отправляет новости и ошибки в соответствующие каналы.
func parseNews(links string, errors chan<- error, posts chan<- []storage.Post, period int) {
	for {
		newPosts, err := rss.News(links)
		if err != nil {
			errors <- err
			continue
		}
		posts <- newPosts
		time.Sleep(time.Minute * time.Duration(period))
	}
}
