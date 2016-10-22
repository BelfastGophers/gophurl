package usecases

import "github.com/BelfastGophers/gophurl/models"

type ShortURLInteractor struct {
	URLRepository models.ShortURLRepository
}

func (interactor *ShortURLInteractor) GetAll() (map[string]*models.ShortURL, error) {
	return interactor.URLRepository.All()
}

func (interactor *ShortURLInteractor) GetUrl(code string) (string, error) {
	result, err := interactor.URLRepository.Get(code)

	if err != nil {
		return "", err
	}
	return result.URL, nil
}

func (interactor *ShortURLInteractor) IncrementAccessCount(code string) error {
	err := interactor.URLRepository.IncrementAccessed(code)
	return err
}
