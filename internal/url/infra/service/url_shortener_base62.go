package service

func NewURLShortener() *URLShortener {
	return &URLShortener{}
}

type URLShortener struct {
}

func (u *URLShortener) ShortenURL(url string) (string, error) {
	return "", nil
}
