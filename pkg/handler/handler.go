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

type Handler struct {
	Services *service.Service
	Mux      *http.ServeMux
}

type Response struct {
	Result string `json:"result"`
}

type ResponseDeals struct {
	Deals []model.Deal
}

func NewHandler(services *service.Service) *Handler {
	h := &Handler{
		Services: services,
		Mux:      http.NewServeMux(),
	}
	h.initRoutes()
	return h
}

func (h *Handler) initRoutes() {
	h.Mux.HandleFunc("/insert/services", h.insertDealHandler)
	h.Mux.HandleFunc("/get/all/services", h.getDealsHandler)
	h.Mux.HandleFunc("/get/services", h.getDealsByNameHandler)
}

func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	corsMiddleware(h.Mux).ServeHTTP(w, r)
}

func (h *Handler) getDealsByNameHandler(w http.ResponseWriter, r *http.Request) {
	var name = r.URL.Query().Get("name")
	if name == "" {
		log.Fatal("error nil name")
	} else {
		log.Println(name)
	}

	deals, err := h.Services.GetDealsByName(name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := &ResponseDeals{
		Deals: deals,
	}

	jsonResponse, err := toJSON(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (h *Handler) insertDealHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("обработка ")
	parsData := parser.NewParserList(r)
	//data := make([]model.Deal, 0)
	data, err := parsData.ParseData()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println("Data ", len(data))

	_ = h.Services.InsertDeals(data, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	for _, item := range data {
		fmt.Println(item)
	}

	//response := &ResponseDeals{
	//	Deals: data,
	//}

	jsonResponse, err := toJSON(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func (h *Handler) getDealsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("обработка ")
	deals, err := h.Services.GetAllDeals()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := &ResponseDeals{
		Deals: deals,
	}

	jsonResponse, err := toJSON(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func toJSON(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if r.Method == http.MethodOptions {
			return
		}
		next.ServeHTTP(w, r)
	})
}
