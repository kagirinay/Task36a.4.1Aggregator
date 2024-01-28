package main

import (
	"Task36a.4.1Aggregator/pkg/conf"
	"fmt"
)

/*
// Сервер.
type server struct {
	db  storage.Interface
	api *api.API
}

func main() {

	// Создаём объект сервера.
	var srv server
	// Создаём каналы для новостей и ошибок.
	chanPosts := make(chan []storage.Post)
	chanErr := make(chan error)
	// Создаём объекты баз данных.
	// БД заглушка
	db1 := memdb.New()
	_ = db1
	// Реляционная БД PostgresSQL.
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	db2, err := postgres.New(ctx, "postgres://postgres:password@192.168.58.133:5432/news")
	if err != nil {
		log.Fatal(err)
	}
	_ = db2
	// Инициализируем хранилище выбранного сервера БД.
	srv.db = db2
	// Создаём объект API и регистрируем обработчики.
	config := conf.ConJson{}
	srv.api = api.New(srv.db)
	myLinks := receivingRSS("C:/Users/kagir/GolandProjects/Task36a.4.1Aggregator/cmd/server/conf.json", chanErr)
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
	err = http.ListenAndServe("localhost:8087", srv.api.Router())
	if err != nil {
		panic(err)
	}
}

// Получает отдельные ссылки из конфигурации, ошибки направляет в поток ошибок.
func receivingRSS(fileName string, errors chan<- error) *conf.ConJson {
	jsonFile, err := os.Open(fileName)
	if err != nil {
		errors <- err
	}
	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			panic(err)
		}
	}(jsonFile)
	byteValue, _ := io.ReadAll(jsonFile)
	var links conf.ConJson
	json.Unmarshal(byteValue, &links)

	return &links
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
*/

func main() {
	fmt.Println("Считывание конфигурации")
	fmt.Println(conf.NewConfig())
	conf.NewConfig()
}
