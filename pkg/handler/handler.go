package handler

import (
	"encoding/json"
	"fmt"
	"github.com/MrClean-code/testD/pkg/model"
	"github.com/MrClean-code/testD/pkg/parser"
	"github.com/MrClean-code/testD/pkg/service"
	"log"
	"net/http"
)

var h23 *Handler

type Handler struct {
	Services *service.Service
}
type Response struct {
	Result string     `json:"result"`
	Deal   model.Deal `json:"deal"`
	Error  string     `json:"error"`
}

type ResponseDeals struct {
	Deals []model.Deal
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{Services: services}
}

func (h *Handler) InitRoutes(h2 *Handler) {
	h23 = h2
	http.HandleFunc("/insert/services", insertDealHandler)       // вставить данные из парсера
	http.HandleFunc("/get/all/services", getDealsHandler)        // получить данные из бд
	http.HandleFunc("/get/services/name", getDealsByNameHandler) // получить данные по названию услуги
}

func getDealsByNameHandler(w http.ResponseWriter, r *http.Request) {
	var name = r.URL.Query().Get("name")
	if name == "" {
		log.Fatal("error nil name")
	} else {
		log.Println(name)
	}

	deals, _ := h23.Services.GetDealsByName(name)

	response := &ResponseDeals{
		Deals: deals,
	}

	jsonResponse, _ := toJSON(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func insertDealHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("обработка ")
	parsData := parser.NewParserList(r)
	dealsMsg := h23.Services.InsertDeals(parsData.ParseData())

	response := &Response{
		Result: dealsMsg,
	}

	jsonResponse, _ := toJSON(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func getDealsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("обработка ")
	deals, _ := h23.Services.GetAllDeals()

	response := &ResponseDeals{
		Deals: deals,
	}

	jsonResponse, _ := toJSON(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func toJSON(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}
