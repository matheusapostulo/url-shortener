package port

type CreateURLInputDto struct {
	LongURL string `json:"long_url"`
}

type CreateURLOutputDto struct {
	ShortURL string `json:"short_url"`
}

type CreateURLUsecase interface {
	Execute(input CreateURLInputDto) (CreateURLOutputDto, error)
}
