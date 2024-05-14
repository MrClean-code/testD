package handler

import (
	"fmt"
	"github.com/MrClean-code/testD/pkg/parser"
	"github.com/MrClean-code/testD/pkg/service"
	"net/http"
	"time"
)

type Handler struct {
	Services *service.Service
}

var h23 *Handler

func NewHandler(services *service.Service) *Handler {
	return &Handler{Services: services}
}

func (h *Handler) InitRoutes(h2 *Handler) {
	h23 = h2
	http.HandleFunc("/search_services", dealHandler)        // получить по названию услуги
	http.HandleFunc("/search_services_int", dealIntHandler) // вставить данные
}

func dealHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("обработка ")
	time.Sleep(2 * time.Second)
	parsData := parser.NewParserList(r)
	h23.Services.InsertDeals(parsData.ParseData())

}

func dealIntHandler(writer http.ResponseWriter, request *http.Request) {

}
