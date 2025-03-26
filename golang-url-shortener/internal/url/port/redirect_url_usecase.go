package port

type RedirectURLInputDto struct {
	ShortURL string `json:"short_url"`
}

type RedirectURLOutputDto struct {
	LongURL string `json:"long_url"`
}

type RedirectURLUsecase interface {
	Execute(input RedirectURLInputDto) (RedirectURLOutputDto, error)
}
