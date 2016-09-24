package data

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/AndrewSpeed/gophurl/models"
)

type JSONShortURLRepository struct {
	Filepath string
	URLList  *models.URLList
}

func NewJSONShortURLRepository(filePath string) (*JSONShortURLRepository, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}

	list := &models.URLList{}

	if err := json.Unmarshal(data, &list); err != nil {
		return nil, err
	}

	repo := &JSONShortURLRepository{
		Filepath: filePath,
		URLList:  list,
	}

	return repo, nil
}

func (repo *JSONShortURLRepository) Get(code string) (*models.ShortURL, error) {
	shortUrls := repo.URLList.GetShortUrls()

	for _, shortUrl := range shortUrls {
		if code == shortUrl.Code {
			return shortUrl, nil
		}
	}
	return nil, errors.New(fmt.Sprintf("Unable to locate url for '%v'", code))
}

func (repo *JSONShortURLRepository) Add(url *models.ShortURL) error {
	return nil
}

func (repo *JSONShortURLRepository) Delete(url *models.ShortURL) error {
	return nil
}

func (repo *JSONShortURLRepository) Update(url *models.ShortURL) error {
	return nil
}

func (repo *JSONShortURLRepository) All() (map[string]*models.ShortURL, error) {
	return repo.URLList.ShortURLs, nil
}

func (repo *JSONShortURLRepository) IncrementAccessed(code string) error {
	shortUrls := repo.URLList.GetShortUrls()

	for _, shortUrl := range shortUrls {
		if code == shortUrl.Code {
			shortUrl.AccessCount++
			err := repo.write()

			if err != nil {
				return err
			}

			return nil
		}
	}
	return errors.New(fmt.Sprintf("Unable to locate ShortUrl with code '%v'", code))
}

func (repo *JSONShortURLRepository) write() error {
	list := repo.URLList

	jsonList, err := json.MarshalIndent(list, "", "    ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(repo.Filepath, jsonList, 0644)
	if err != nil {
		return err
	}

	return nil
}
