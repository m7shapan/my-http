package jobs

import (
	"crypto/md5"
	"fmt"

	"github.com/m7shapan/my-http/models"
	"github.com/m7shapan/my-http/repositories"
)

type hashingJob struct {
	responseRepository repositories.ResponseRepository
	url                string
}

func NewHashingJob(r repositories.ResponseRepository, url string) *hashingJob {
	return &hashingJob{
		responseRepository: r,
		url:                url,
	}
}

func (j *hashingJob) Execute() interface{} {
	resp, err := j.responseRepository.Get(j.url)
	if err != nil {
		return map[string]interface{}{
			"err": err,
		}
	}

	return map[string]interface{}{
		"response": models.Response{
			URL:  j.url,
			Hash: fmt.Sprintf("%x", md5.Sum(resp)),
		},
	}
}
