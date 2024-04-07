package parser

import (
	"fmt"
	"github.com/MrClean-code/testD/pkg/model"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type ParserDealList struct {
	Deals []model.Deal
	r     *http.Request
}

func NewParserDealList(r *http.Request) *ParserDealList {
	return &ParserDealList{r: r}
}

func (p *ParserDealList) ParseData() ([]model.Deal, error) {

	//var pN = p.r.URL.Query().Get("name")
	//if pN == "" {
	//	log.Fatal("error nil pN")
	//} else {
	//	log.Println(pN)
	//}
	//
	//requestURL := fmt.Sprintf("https://www.avito.ru/voronezh?q=%s", pN)
	//fmt.Println(requestURL)
	//
	//time.Sleep(time.Second)
	//// Создаем новый HTTP клиент
	//client := &http.Client{}
	//
	//// Создаем новый HTTP запрос
	//req, err := http.NewRequest("GET", requestURL, nil)
	//if err != nil {
	//	log.Fatal("Ошибка при создании запроса:", err)
	//}
	//
	//// Добавляем заголовок User-Agent к запросу
	//req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/58.0.3029.110 Safari/537.3")
	//req.Header.Set("Content-Type", "text/html")
	//
	//// Выполняем запрос
	//response, err := client.Do(req)
	//time.Sleep(2 * time.Second)
	//if err != nil {
	//	log.Fatal("Ошибка при отправке запроса:", err)
	//}
	//defer response.Body.Close()
	//
	//body, err := ioutil.ReadAll(response.Body)
	//if err != nil {
	//	log.Fatal("Ошибка при чтении ответа:", err)
	//}
	//
	//fmt.Println("Статус код ответа:", response.Status)
	//fmt.Println("Тело ответа:", string(body)) // преобразование []byte в строку для вывода

	var pN = p.r.URL.Query().Get("name")
	if pN == "" {
		log.Fatal("error nil pN")
	} else {
		log.Println(pN)
	}

	requestURL := fmt.Sprintf("https://uslugi.yandex.ru/193-voronezh/category?from=suggest&text=%s", pN)
	fmt.Println(requestURL)

	err := downloadWebsite(requestURL)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return nil, nil
}

func downloadWebsite(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Failed to fetch the website. Status code: %d", resp.StatusCode)
	}

	// Extract a valid directory name from the URL
	baseURL := extractValidDirName(url)
	err = os.MkdirAll(baseURL, os.ModePerm)
	if err != nil {
		return err
	}

	return parseHTML(resp.Body, baseURL)
}

func extractValidDirName(url string) string {
	// Replace invalid characters with underscores
	validChars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789-_"
	result := strings.Map(func(r rune) rune {
		if strings.ContainsRune(validChars, r) {
			return r
		}
		return '_'
	}, url)

	// Ensure the result is not empty
	if result == "" {
		return "invalid_dir"
	}

	return result
}

func parseHTML(body io.Reader, baseURL string) error {
	tokenizer := html.NewTokenizer(body)

	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			return nil
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			if token.Data == "a" {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						link := attr.Val
						if strings.HasPrefix(link, "http") || strings.HasPrefix(link, "www") {
							downloadFile(link, baseURL)
						}
						fmt.Println("Found link:", link)
					}
				}
			}
		}
	}
}

func downloadFile(link, baseURL string) {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	filename := filepath.Join(baseURL, filepath.Base(link))
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Downloaded: %s\n", filename)
}
