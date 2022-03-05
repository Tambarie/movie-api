package api

import (
	"encoding/json"
	"fmt"
	domain "github.com/Tambarie/movie-api/internal/core/domain/movie"
	"github.com/pkg/errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const swapiURl = "https://swapi.dev/api"

func GetAllMovies() (*[]domain.Movie, error) {
	url := fmt.Sprintf("%s/films/", swapiURl)
	resp, err := http.Get(url)

	if err != nil {
		return nil, errors.Wrap(err, "unable to get swapiURL")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "unable to read response from swapiAPI")
	}

	movies := &domain.Movies{}
	if err := json.Unmarshal(body, movies); err != nil {
		return nil, errors.Wrap(err, "unable to unmarshal response body from swapiAPI")
	}
	return &movies.Results, nil
}

func GetAllCharactersByMovieID(movieId int) (*[]string, error) {
	url := fmt.Sprintf("%s/films/%d/", swapiURl, movieId)

	resp, err := http.Get(url)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to get character from swapi api"))
		return nil, errors.Wrap(err, "failed to get character from swapi api")
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println(err)
		}
	}(resp.Body)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body from swapi api")
	}

	var data domain.CharacterLinks

	if err := json.Unmarshal(body, &data); err != nil {
		log.Println(errors.Wrap(err, "failed to unmarshal response body from swapi api"))
		return nil, errors.Wrap(err, "failed to unmarshal response body from swapi api")
	}
	return &data.Characters, nil
}

func GetCharacterInfo(link string) (*domain.Character, error) {
	resp, err := http.Get(link)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to get character from swapi api"))
		return nil, errors.Wrap(err, "failed to get character from swapi api")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(errors.Wrap(err, "failed to read response body from swapi api"))
		return nil, errors.Wrap(err, "failed to read response body from swapi api")
	}

	var data domain.Character

	if err := json.Unmarshal(body, &data); err != nil {
		log.Println(errors.Wrap(err, "failed to unmarshal response body from swapi api"))
		return nil, errors.Wrap(err, "failed to unmarshal response body from swapi api")
	}
	return &data, nil
}
