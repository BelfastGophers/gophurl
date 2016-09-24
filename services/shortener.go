package services

import (
	"github.com/AndrewSpeed/gophurl/handlers"
	"github.com/AndrewSpeed/gophurl/models"

	"fmt"
	"net/http"
)

// Shortener provides a collection of web handlers
type WebShortener interface {
	Index(http.ResponseWriter, *http.Request)
	GetRedirectUrl(string) (string, error)
}

// Shortener is a service which provides a web accessible URL shortener.
type Shortener struct {
	Handler WebShortener
}

// Start creates an HTTP server and starts it on the given port number
func (svc *Shortener) Start(port int) error {
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}

func NewShortener(urlRepo models.ShortURLRepository) *Shortener {
	// Take a repository and pass it to the handler. We are agnostic about
	// what the underlying storage is and are only working with the
	// interface
	handler := handlers.NewWebShortener(urlRepo)

	svc := &Shortener{
		Handler: handler,
	}
	http.HandleFunc("/", svc.Handler.Index)
	return svc
}
