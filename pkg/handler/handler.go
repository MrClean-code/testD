package handler

import (
	"fmt"
	"github.com/MrClean-code/testD/pkg/parser"
	"github.com/MrClean-code/testD/pkg/service"
	"net/http"
	"time"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() {
	http.HandleFunc("/search_services", dealHandler)

}

func dealHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("обработка ")
	time.Sleep(2 * time.Second)
	parsData := parser.NewParserList(r)
	parsData.ParseData()

	//user_id, date, err := parseEventParamsGet(r)

}
