package postgres

import (
	"Task36a.4.1Aggregator/pkg/storage"
	"context"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	_, err := New(ctx, "postgres://postgres:password@192.168.58.133:5432/news")
	if err != nil {
		t.Fatal(err)
	}
}

func TestStore_AddPost(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	dataBase, err := New(ctx, "postgres://postgres:password@192.168.58.133:5432/news")
	post := storage.Post{
		Title:       "тестирования",
		Content:     "Пробный текст",
		PublishedAt: 5,
		Link:        "Ссылка",
	}
	dataBase.AddPost(post)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Создана запись.")
}
