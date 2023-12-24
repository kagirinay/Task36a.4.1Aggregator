package postgres

import (
	"Task36a.4.1Aggregator/pkg/storage"
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
)

// Store Хранилище данных.
type Store struct {
	db *pgxpool.Pool
}

// New Конструктор объекта хранилища.
func New(ctx context.Context, constr string) (*Store, error) {
	db, err := pgxpool.Connect(ctx, constr)
	if err != nil {
		return nil, err
	}
	s := Store{db: db}
	return &s, err
}

// Posts Выводит все существующие новости.
func (s *Store) Posts(n int) ([]storage.Post, error) {
	rows, err := s.db.Query(context.Background(), `
		SELECT
			id,
			title,
			content,
			publishedAt,
			link
		FROM 
		ORDER BY published DESC
		;
		`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []storage.Post
	// Итерирование по результату выполнения запроса и сканирование каждой строки в переменную.
	for rows.Next() {
		var t storage.Post
		err = rows.Scan(
			&t.ID,
			&t.Title,
			&t.Content,
			&t.PublishedAt,
			&t.Link,
		)
		if err != nil {
			return nil, err
		}
		// Добавление переменной в массив результатов.
		posts = append(posts, t)
	}
	//Важно не забыть проверить rows.Err()
	return posts, rows.Err()
}
