package storage

// Post Публикация, получаемая из RSS.
type Post struct {
	ID          int    // Идентификатор записи.
	Title       string // Заголовок новости.
	Content     string // Содержание новости.
	PublishedAt int64  // Время публикации новости.
	Link        string // Ссылка на источник новости.
}

// Interface Задаёт контракт на работу с БД.
type Interface interface {
	Posts(n int) ([]Post, error) // Получение последних новостей.
	AddPost(t Post) error        // Добавление новости в БД.
}
