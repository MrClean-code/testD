package main

import (
	"github.com/MrClean-code/testD/pkg/handler"
	"github.com/MrClean-code/testD/pkg/repository"
	"github.com/MrClean-code/testD/pkg/service"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"net/url"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("ошибка инициализации конфигурации: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("ошибка загрузки переменных окружения: %s", err.Error())
	}

	db, err := repository.NewPostgresDB()
	if err != nil {
		log.Fatalf("не удалось инициализировать базу данных: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	port := viper.GetString("PORT")
	addr := ":" + port
	if port == "" {
		log.Fatal("Порт не указан в конфигурации")
	}

	log.Printf("Сервер запущен на порту %s\n", addr)

	// Создание HTTP клиента в режиме инкогнито
	client, err := createIncognitoHTTPClient()
	if err != nil {
		log.Fatalf("Не удалось создать HTTP клиент: %v", err)
	}

	// Пример выполнения запроса с использованием инкогнито HTTP клиента
	resp, err := client.Get("http://example.com")
	if err != nil {
		log.Fatalf("Тестовая проверка http://example.com: %v", err)
	}
	defer resp.Body.Close()

	// log.Printf("Статус ответа: %s", resp.Status)

	if err := http.ListenAndServe(addr, handlers); err != nil {
		log.Fatalf("Не удалось запустить сервер: %v", err)
	}
}

func initConfig() error {
	viper.SetConfigName("config")
	viper.AddConfigPath("configs")
	return viper.ReadInConfig()
}

func createIncognitoHTTPClient() (*http.Client, error) {
	proxyURL, err := url.Parse("http://31.43.179.54:80")
	if err != nil {
		return nil, err
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}

	client := &http.Client{
		Transport: &transportWithCustomHeaders{
			Transport: transport,
			headers: map[string]string{
				"User-Agent":      "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, как Gecko) Chrome/91.0.4472.124 Safari/537.36",
				"Accept-Language": "en-US,en;q=0.9",
			},
		},
		Jar: nil, // Отключаем использование куки
	}

	return client, nil
}

type transportWithCustomHeaders struct {
	Transport http.RoundTripper
	headers   map[string]string
}

func (t *transportWithCustomHeaders) RoundTrip(req *http.Request) (*http.Response, error) {
	for key, value := range t.headers {
		req.Header.Set(key, value)
	}
	req.Header.Del("Cookie") // Удаляем куки из заголовков
	return t.Transport.RoundTrip(req)
}
