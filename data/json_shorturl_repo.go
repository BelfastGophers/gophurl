package data

import (
	"github.com/BelfastGophers/gophurl/models"
)

type JsonShortURLRepository struct {
	Filepath string
	URLList  *models.URLList
}

func NewJsonShortURLRepository(filePath string) (*JsonShortURLRepository, error) {
	repo := &JsonShortURLRepository{}
	return repo, nil
}

func (repo *JsonShortURLRepository) Get(code string) (*models.ShortURL, error) {
	return nil, nil
}

func (repo *JsonShortURLRepository) Add(url *models.ShortURL) error {
	return nil
}

func (repo *JsonShortURLRepository) Delete(url *models.ShortURL) error {
	return nil
}

func (repo *JsonShortURLRepository) Update(url *models.ShortURL) error {
	return nil
}

func (repo *JsonShortURLRepository) All() (map[string]*models.ShortURL, error) {
	return map[string]*models.ShortURL{}, nil
}

func (repo *JsonShortURLRepository) IncrementAccessed(code string) error {
	return nil
}

func (repo *JsonShortURLRepository) write() error {
	return nil
}
