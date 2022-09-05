package controller

import (
	"URLshortener/service"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"net/http"
)

type Handlers struct {
	urlService service.URLService
}

func NewHandlers(urlService *service.URLService) *Handlers {
	return &Handlers{*urlService}
}

func (h *Handlers) AddUrl(w http.ResponseWriter, r *http.Request) {
	newUrl, _ := io.ReadAll(r.Body)
	shortUrl := h.urlService.AddNewUrlPair(string(newUrl))
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, shortUrl)
	return
}

func (h *Handlers) GetUrl(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	code := vars["code"]
	originalUrl, urlExists := h.urlService.GetOriginalUrl(code)
	if !urlExists {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "URL not found")
	} else {
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, originalUrl)
	}
	return
}
