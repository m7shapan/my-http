package cmd

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/m7shapan/my-http/repositories"
	"github.com/m7shapan/my-http/services"
)

type CMD struct {
}

func (c CMD) Start() {
	c.md5ResponseHandler()
}

func (c CMD) md5ResponseHandler() {
	parallel := flag.Int("parallel", 10, "the max number of parallel requests")
	flag.Parse()

	urls := flag.Args()
	for i := 0; i < len(urls); i++ {
		if !strings.HasPrefix(urls[i], "http") {
			urls[i] = "http://" + urls[i]
		}
	}

	httpRepository := repositories.NewHTTPResponseRepository()
	md5Service := services.NewMD5HashingService(httpRepository)

	responses, err := md5Service.HashParallel(urls, *parallel)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < len(responses); i++ {
		fmt.Println(responses[i].URL, responses[i].Hash)
	}
}
