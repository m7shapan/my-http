package services

import (
	"crypto/md5"
	"fmt"

	"github.com/m7shapan/my-http/models"
	"github.com/m7shapan/my-http/pkg/dispatchers"
	"github.com/m7shapan/my-http/pkg/jobs"
	"github.com/m7shapan/my-http/repositories"
)

type HashingService interface {
	Hash([]string) ([]models.Response, error)
	HashParallel([]string, int) ([]models.Response, error)
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

func (h md5HashingService) HashParallel(urls []string, maxParallel int) ([]models.Response, error) {
	var responses []models.Response
	resultChannel := make(chan interface{}, maxParallel)
	dispatcher := dispatchers.NewDispatcher(resultChannel, maxParallel)

	dispatcher.Start()
	for i := 0; i < len(urls); i++ {
		dispatcher.Push(jobs.NewHashingJob(h.responseRepository, urls[i]))
	}

	dispatcher.Close()

	for r := range resultChannel {

		r := r.(map[string]interface{})
		if err, found := r["err"]; found {
			return nil, err.(error)
		}

		responses = append(responses, r["response"].(models.Response))
	}

	return responses, nil
}
