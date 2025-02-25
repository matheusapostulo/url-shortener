package http

import (
	"encoding/json"
	"net/http"

	"github.com/matheusapostulo/url-shortener/internal/url/usecase"
)

func NewURLHandler(createUC usecase.CreateURLUsecase) *URLHandler {
	return &URLHandler{
		createUrlUsecase: createUC,
	}
}

type URLHandler struct {
	createUrlUsecase usecase.CreateURLUsecase
}

func (h *URLHandler) CreateURL(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateURLInputDto

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.createUrlUsecase.Execute(input)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(output)
}
