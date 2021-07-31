package repositories

import (
	"io/ioutil"
	"net/http"
	"strings"
)

type ResponseRepository interface {
	Get(string) ([]byte, error)
}

type httpResponseRepository struct {
}

func (r httpResponseRepository) Get(url string) ([]byte, error) {
	if !strings.HasPrefix(url, "http") {
		url = "http://" + url
	}

	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	return ioutil.ReadAll(response.Body)
}

func NewHTTPResponseRepository() ResponseRepository {
	return httpResponseRepository{}
}
