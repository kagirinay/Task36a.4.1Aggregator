package rss

import (
	"testing"
)

func TestRSSToStruct(t *testing.T) {
	ups, err := News("https://habr.com/ru/rss/hub/go/all/?f1=ru")
	if err != nil {
		t.Fatal(err)
	}
	if len(ups) == 0 {
		t.Fatal("Данные не раскодированны.")
	}
	t.Logf("Полученно %d новостей\n%+v", len(ups), ups)

	ups, err = News("https://habr.com/ru/rss/best/daily/?f1=ru")
	if err != nil {
		t.Fatal(err)
	}
	if len(ups) == 0 {
		t.Fatal("Данные не раскодированы.")
	}
	t.Logf("Получено %d новостей\n%+v", len(ups), ups)
}
