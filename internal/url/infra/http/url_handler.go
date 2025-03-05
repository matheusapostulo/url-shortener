package http

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
	"github.com/matheusapostulo/url-shortener/internal/url/domain"
	"github.com/matheusapostulo/url-shortener/internal/url/port"
)

func NewURLHandler(createUC port.CreateURLUsecase, redirectUC port.RedirectURLUsecase) *URLHandler {
	return &URLHandler{
		createUrlUsecase:   createUC,
		redirectUrlUsecase: redirectUC,
	}
}

type URLHandler struct {
	createUrlUsecase   port.CreateURLUsecase
	redirectUrlUsecase port.RedirectURLUsecase
}

func (h *URLHandler) CreateURL(w http.ResponseWriter, r *http.Request) {
	var input port.CreateURLInputDto

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	output, err := h.createUrlUsecase.Execute(input)
	if err != nil {
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	response.JSON(w, http.StatusCreated, output)
}

func (h *URLHandler) RedirectURL(w http.ResponseWriter, r *http.Request) {
	shortUrl := chi.URLParam(r, "short-url")

	input := port.RedirectURLInputDto{
		ShortURL: shortUrl,
	}

	output, err := h.redirectUrlUsecase.Execute(input)
	if err != nil {
		if errors.Is(err, domain.ErrURLNotFound) {
			response.Error(w, http.StatusNotFound, err.Error())
			return
		}
		response.Error(w, http.StatusInternalServerError, err.Error())
		return
	}

	http.Redirect(w, r, output.LongURL, http.StatusMovedPermanently)
}
