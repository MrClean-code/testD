package parser

import (
	"fmt"
	"github.com/MrClean-code/testD/pkg/model"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"log"
	"math/rand"
	"net/http"
	"regexp"
	"strings"
	"time"
)

type ParserDealList struct {
	Deals []model.Deal
	r     *http.Request
}

func NewParserDealList(r *http.Request) *ParserDealList {
	return &ParserDealList{r: r}
}

var sl2 []model.Deal

func (p *ParserDealList) ParseData() ([]model.Deal, error) {

	p.Deals = make([]model.Deal, 0)

	var pN = p.r.URL.Query().Get("name")
	if pN == "" {
		log.Fatal("error nil pN")
	} else {
		log.Println(pN)
	}

	// Путь к исполняемому файлу ChromeDriver (необходимо скачать и установить)
	const chromeDriverPath = "C:\\Program Files\\chromedriver_win32\\chromedriver.exe"

	// Настройка ChromeDriver для использования браузера Chrome
	fmt.Println("start")
	opts := []selenium.ServiceOption{}
	service, err := selenium.NewChromeDriverService(chromeDriverPath, 9515, opts...)
	if err != nil {
		fmt.Printf("Ошибка создания сервиса ChromeDriver: %v\n", err)
		//return
	}
	defer service.Stop()

	// Опции браузера Chrome
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	chromeCaps := chrome.Capabilities{
		Args: []string{
			// Дополнительные параметры Chrome
			// Например, можно скрыть окно браузера и запустить в фоновом режиме
			//"headless",
			"window-size=1920,1080",
		},
	}
	caps.AddChrome(chromeCaps)

	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", 9515))
	if err != nil {
		fmt.Printf("Ошибка создания WebDriver: %v\n", err)
		//return
	}
	defer wd.Quit()

	url := fmt.Sprintf("https://www.avito.ru/voronezh?q=%s", pN)
	fmt.Println(url)

	fmt.Println("Открытие страницы...")
	if err := wd.Get(url); err != nil {
		fmt.Printf("Ошибка открытия страницы: %v\n", err)
		//return
	}
	fmt.Println("Страница успешно открыта")

	rand.Seed(time.Now().UnixNano())
	randomSeconds := rand.Intn(10) + 5
	time.Sleep(time.Duration(randomSeconds) * time.Second)

	fmt.Println("Получение текста элементов...")
	elements, err := wd.FindElements(selenium.ByCSSSelector, "h3[itemprop='name'], p[data-marker='item-price'], div.style-root-uufhX, span.desktop-1y4c9y0")

	time.Sleep((time.Duration(randomSeconds) - 5) * time.Second)
	if err != nil {
		fmt.Printf("Ошибка при поиске элементов: %v\n", err)
		//return
	}
	fmt.Println("Элементы успешно найдены")

	sl := make([]string, 0)
	for _, elem := range elements {
		text, _ := elem.Text()
		if text != "" {
			//fmt.Println(text)
			sl = append(sl, text)
		}

		// Проверяем, есть ли элемент <a> внутри текущего элемента
		aElement, err := elem.FindElement(selenium.ByTagName, "a")
		if err != nil {
			continue
		}

		// Получаем значение атрибута href только если элемент <a> найден
		href, _ := aElement.GetAttribute("href")
		if href != "" {
			//fmt.Println(href)
			sl = append(sl, href)
		}
	}

	sl = DeleteFirstElement(sl) // ов
	sl = DeleteDataBetweenLink(sl)
	sl = DeleteFisrtTail(sl)
	sl = DeleteDoubleLink(sl)
	sl2 = SplitData(sl)
	sl2 = EditDataFormat(sl2)

	sl2 = sl2[:len(sl2)-1] // убираем последний символ

	//for i := 0; i < len(sl2); i++ {
	//	fmt.Println(sl2[i])
	//}

	return sl2, nil
}
func EditDataFormat(sl2 []model.Deal) []model.Deal {
	for i := 0; i < len(sl2); i++ {
		str := strings.ReplaceAll(sl2[i].Price, " ", "")

		// Создаем регулярное выражение для поиска числа
		re := regexp.MustCompile(`\d+`)

		// Ищем все числа в строке
		numbers := re.FindAllString(str, -1)

		combined := strings.Join(numbers, "")

		sl2[i].Price = combined
		sl2[i].Score = strings.ReplaceAll(sl2[i].Score, ",", ".")
		str = strings.ReplaceAll(sl2[i].CountReviews, " ", "")
		numbers = re.FindAllString(str, -1)
		combined = strings.Join(numbers, "")
		sl2[i].CountReviews = combined
	}

	return sl2
}

func DeleteFisrtTail(sl []string) []string {
	for i, _ := range sl {

		if strings.HasPrefix(sl[i], "http") {
			cSl := len(sl[:i])
			if cSl > 5 {
				sl = append(sl[cSl-6:], sl[cSl:]...)
			}
			break
		}

	}
	return sl
}

func DeleteDoubleLink(sl []string) []string {
	for i := len(sl) - 2; i >= 0; i-- {
		if strings.HasPrefix(sl[i], "http") && strings.HasPrefix(sl[i+1], "http") {
			sl = append(sl[:i+1], sl[i+2:]...)
			fmt.Println("")
		}
	}
	return sl
}

func DeleteFirstElement(sl []string) []string {
	for i := 0; i < len(sl)-2; i++ {
		if strings.HasPrefix(sl[i], "http") && len(sl[i+1]) == 3 {
			sl = append(sl[:i+1], sl[i+2:]...)
		}
	}
	return sl
}

func DeleteDataBetweenLink(sl []string) []string {
	var lineCount int
	var lastIndex int
	var i2 int

	for i, _ := range sl {
		if i2 != 0 {
			i = i2 // lastIndex
		}

		if strings.HasPrefix(sl[i], "http") {

			// Если это первая ссылка, просто обновляем индекс последней ссылки и счётчик строк
			if lastIndex == 0 {
				lastIndex = i
				lineCount = 0
			} else {
				// Если между ссылками больше 5 строк, удаляем элементы между ними
				if lineCount > 3 && (lastIndex < len(sl)-2) {
					sl = append(sl[:lastIndex+1], sl[i+1:]...)
					// Обновляем индекс последней ссылки и обнуляем счётчик строк
					i2 = lastIndex
					lineCount = 0
				} else if lineCount == 3 {
					lastIndex = i2
					i2++
					lineCount = 0
				} else if lineCount < 3 {
					i2++
				}
			}
		} else {
			// Если элемент не является ссылкой, увеличиваем счётчик строк
			lineCount++

			if i == i2 {
				i2++
			}
		}
	}
	return sl
}

func SplitData(sl []string) []model.Deal {
	c := 0
	sl2 = make([]model.Deal, 0)
	i2 := 0
	for i, item := range sl {

		if strings.HasPrefix(item, "http") && c == 3 && c < len(sl)-1 {
			i2++
			comments := strings.Split(sl[i-1], "\n")
			deal := &model.Deal{
				i2,
				sl[i-3],
				comments[0],
				sl[i-2],
				comments[2],
				comments[1],
				sl[i],
			}
			sl2 = append(sl2, *deal)
			c = 0
		} else if strings.HasPrefix(item, "http") && c != 3 {
			c = 0
		} else {
			c++
		}
	}
	return sl2
}
