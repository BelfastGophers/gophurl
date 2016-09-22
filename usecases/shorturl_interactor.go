package usecases

import (
	"github.com/BelfastGophers/gophurl/models"
)

type ShortURLInteractor struct {
	URLRepository models.ShortURLRepository
}

func (interactor *ShortURLInteractor) GetAll() ([]*models.ShortURL, error) {
	return []*models.ShortURL{}, nil
}
