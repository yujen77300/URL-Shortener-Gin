package models

type ShortenURL struct {
	ShortString string `json:"shortUrl" bson:"short_string"`
	OriginalURL string `json:"originalUrl" bson:"original_url"`
}
