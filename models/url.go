package models

import (
	"time"
)

// UrlList is a container for a list of URLs and their shortened forms
type URLList struct {
	LastModified *time.Time           `json:"updated"`
	ShortURLs    map[string]*ShortURL `json:"urls"`
}

// A ShortUrl is a structure which captures the
type ShortURL struct {
	URL         string     `json:"url"`
	Code        string     `json:"short_code"`
	AccessCount int        `json:"access_count"`
	Created     *time.Time `json:"created"`
}

// ShortURLRepository is the interface by which all usecases should interact
// ShortURL's from storage. The implementation detail may differ by storage
// mechanism
type ShortURLRepository interface {
	Get(string) (*ShortURL, error)
	Add(*ShortURL) error
	Delete(*ShortURL) error
	Update(*ShortURL) error
	All() (map[string]*ShortURL, error)
	IncrementAccessed(string) error
}
