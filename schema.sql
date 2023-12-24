DROP TABLE IF EXISTS news;

CREATE TABLE news (
    id SERIAL PRIMARY KEY,
    title TEXT, -- Заголовок новости.
    content TEXT NOT NULL UNIQUE, --Содержание новости.
    publishedAt BIGINT DEFAULT 0, --Время публикации новости.
    link TEXT --Ссылка на опубликованную новость.
);