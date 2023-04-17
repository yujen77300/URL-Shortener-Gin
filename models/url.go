package models

type ShortenURL struct {
	ShortURL string `json:"shortUrl" bson:"short_url"`
	OriginalURL string `json:"originalUrl" bson:"original_url"`
}
