package port

type URLShortener interface {
	ShortenURL(id int) (string, error)
}
