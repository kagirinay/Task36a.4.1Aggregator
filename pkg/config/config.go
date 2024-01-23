package config

// Конфигурация приложения.
type configJson struct {
	Period  int      `json:"request_period"`
	LinkArr []string `json:"rss"`
}
