package storage

// Post Публикация, получаемая из RSS
type Post struct {
	ID      int    // Номер публикации.
	Title   string // Заголовок публикации.
	Content string // Содержание публикации.
	PubTime int64  // Время публикации.
	Link    string // Ссылка на источник.
}

// Interface Методы работы с БД.
type Interface interface {
	News(limit int) ([]Post, error) // Получаем все публикации.
	AddPosts(posts []Post) error    // Добавление постов в БД.
}
