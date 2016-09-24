package handlers

import (
	"github.com/AndrewSpeed/gophurl/handlers/html"
	"github.com/AndrewSpeed/gophurl/models"
	"github.com/AndrewSpeed/gophurl/usecases"

	"html/template"
	"net/http"
)

// ShortURLInteractor is an interface which the handler can use to interact
// with a collection of ShortURLs. The implmentation of the interactor is
// where all business rules can be implemented
type ShortURLInteractor interface {
	GetAll() (map[string]*models.ShortURL, error)
	GetUrl(string) (string, error)
	IncrementAccessCount(string) error
}

// WebShortener provides a number of HTTP Handlers.
type WebShortener struct {
	ShortURLInteractor ShortURLInteractor
}

// Index is the generic handler that will capture any route that has not been assigned
// a handler. It has 2 purposes:
//   * If no path is provided, display the list of stored Short URLs
//   * If a path is provided, try to redirect to it
func (handler *WebShortener) Index(w http.ResponseWriter, req *http.Request) {
	// TODO: Check if a path is provided. If it is follow it
	shortUrlCode := req.URL.Path[1:]

	// Get all of the URL's from the Repository
	urls, err := handler.ShortURLInteractor.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// if a shortUrlCode is present in the URL, redirect to it
	if shortUrlCode != "" {
		redirectUrl, err := handler.GetRedirectUrl(shortUrlCode)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		err = handler.ShortURLInteractor.IncrementAccessCount(shortUrlCode)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		http.Redirect(w, req, redirectUrl, http.StatusFound)
	}

	tmpl, err := template.New("name").Parse(html.Index)
	if err = tmpl.Execute(w, urls); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (handler *WebShortener) GetRedirectUrl(urlCode string) (string, error) {
	return handler.ShortURLInteractor.GetUrl(urlCode)
}

// NewShortenerWeb returns a struct which implements the WebShortener interface. It
// populates fields to usable values
func NewWebShortener(urlRepo models.ShortURLRepository) *WebShortener {
	return &WebShortener{
		ShortURLInteractor: &usecases.ShortURLInteractor{
			URLRepository: urlRepo,
		},
	}
}
