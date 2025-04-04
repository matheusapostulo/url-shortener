package domain

func NewURL(ID int, longURL string, shortURL string) (url *URL, err error) {
	url = &URL{
		ID:       ID,
		LongURL:  longURL,
		ShortURL: shortURL,
	}

	err = url.Validate()
	if err != nil {
		return nil, err
	}

	return url, nil
}

type URL struct {
	ID       int    `json:"id"`
	LongURL  string `json:"long_url"`
	ShortURL string `json:"short_url"`
}

func (u *URL) IsEmpty() bool {
	return u == nil
}

func (u *URL) Validate() error {
	return nil
}
