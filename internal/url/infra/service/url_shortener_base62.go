package service

import (
	"strings"
)

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func NewURLShortenerBase62() *URLShortenerBase62 {
	return &URLShortenerBase62{}
}

type URLShortenerBase62 struct {
}

func (u *URLShortenerBase62) ShortenURL(id int) (string, error) {
	chars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var result strings.Builder

	mod := id % 62
	remainder := id / 62
	result.WriteByte(chars[mod])

	for remainder > 0 {
		mod = remainder % 62
		remainder = remainder / 62
		result.WriteByte(chars[mod])
	}

	return reverseString(result.String()), nil
}
