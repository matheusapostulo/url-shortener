package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/matheusapostulo/url-shortener/internal/url/domain"
	"github.com/matheusapostulo/url-shortener/internal/url/usecase"
)

func NewURLHandler(createUC usecase.CreateURLUsecase, redirectUC usecase.RedirectURLUsecase) *URLHandler {
	return &URLHandler{
		createUrlUsecase:   createUC,
		redirectUrlUsecase: redirectUC,
	}
}

type URLHandler struct {
	createUrlUsecase   usecase.CreateURLUsecase
	redirectUrlUsecase usecase.RedirectURLUsecase
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

func (h *URLHandler) RedirectURL(w http.ResponseWriter, r *http.Request) {
	shortUrl := chi.URLParam(r, "short-url")

	input := usecase.RedirectURLInputDto{
		ShortURL: shortUrl,
	}

	output, err := h.redirectUrlUsecase.Execute(input)
	if err != nil {
		if errors.Is(err, domain.ErrURLNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, output.LongURL, http.StatusMovedPermanently)
}
