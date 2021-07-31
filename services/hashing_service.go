package services

import (
	"crypto/md5"
	"fmt"

	"github.com/m7shapan/my-http/models"
	"github.com/m7shapan/my-http/repositories"
)

type HashingService interface {
	Hash([]string) ([]models.Response, error)
}

type md5HashingService struct {
	responseRepository repositories.ResponseRepository
}

func NewMD5HashingService(r repositories.ResponseRepository) HashingService {
	return md5HashingService{
		responseRepository: r,
	}
}

func (h md5HashingService) Hash(urls []string) ([]models.Response, error) {
	var responses []models.Response
	for i := 0; i < len(urls); i++ {
		resp, err := h.responseRepository.Get(urls[i])
		if err != nil {
			return nil, err
		}

		responses = append(responses, models.Response{
			URL:  urls[i],
			Hash: fmt.Sprintf("%x", md5.Sum(resp)),
		})
	}

	return responses, nil
}
