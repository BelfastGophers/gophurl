package handlers

import (
	"github.com/BelfastGophers/gophurl/handlers/html"
	"github.com/BelfastGophers/gophurl/models"
	"github.com/BelfastGophers/gophurl/usecases"

	"html/template"
	"net/http"
)

// ShortURLInteractor is an interface which the handler can use to interact
// with a collection of ShortURLs. The implmentation of the interactor is
// where all business rules can be implemented
type ShortURLInteractor interface {
	GetAll() ([]*models.ShortURL, error)
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

	// Get all of the URL's from the Repository
	_, err := handler.ShortURLInteractor.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	tmpl, err := template.New("name").Parse(html.Index)
	if err = tmpl.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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
