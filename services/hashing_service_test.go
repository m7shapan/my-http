package services

import (
	"crypto/md5"
	"fmt"
	"testing"
)

func TestTrue(t *testing.T) {
	if !true {
		t.Error("true is not true")
	}
}

type mockRepository struct {
	Body []byte
	Err  error
}

func (r mockRepository) Get(url string) ([]byte, error) {
	return r.Body, r.Err
}

func TestMD5HashingServiceHash(t *testing.T) {
	body := []byte("request body")
	repository := mockRepository{
		Body: body,
	}

	md5Service := NewMD5HashingService(repository)
	hashes, err := md5Service.Hash([]string{"http://google.com"})
	if err != nil {
		t.Error(err)
	}

	if len(hashes) != 1 {
		t.Errorf("hashes length = %d; expected 1\n", len(hashes))
	}

	if hashes[0].Hash != fmt.Sprintf("%x", md5.Sum(body)) {
		t.Errorf("hash value = %s; expected %s\n", hashes[0].Hash, fmt.Sprintf("%x", md5.Sum(body)))
	}
}

func TestMD5HashingServiceHashParallel(t *testing.T) {
	body := []byte("request body")
	repository := mockRepository{
		Body: body,
	}

	md5Service := NewMD5HashingService(repository)
	hashes, err := md5Service.HashParallel([]string{"http://google.com"}, 10)
	if err != nil {
		t.Error(err)
	}

	if len(hashes) != 1 {
		t.Errorf("hashes length = %d; expected 1\n", len(hashes))
	}

	if hashes[0].Hash != fmt.Sprintf("%x", md5.Sum(body)) {
		t.Errorf("hash value = %s; expected %s\n", hashes[0].Hash, fmt.Sprintf("%x", md5.Sum(body)))
	}
}

func TestMD5HashingServiceHashParallelWithErr(t *testing.T) {
	err := fmt.Errorf("unexpected error")
	repository := mockRepository{
		Err: err,
	}

	md5Service := NewMD5HashingService(repository)
	_, got := md5Service.HashParallel([]string{"http://google.com"}, 10)
	if got == nil {
		t.Error(err)
		t.Errorf("no error; expected error %s\n", err)
	}
}

type mockRepositoryForParallel struct {
	Response map[string][]byte
	Err      error
}

func (r mockRepositoryForParallel) Get(url string) ([]byte, error) {
	return r.Response[url], r.Err
}

func TestMD5HashingServiceHashParallelWithMultibleValues(t *testing.T) {
	var values = map[string][]byte{
		"a": []byte("a Response"),
		"b": []byte("b Response"),
		"c": []byte("c Response"),
		"d": []byte("d Response"),
		"e": []byte("e Response"),
		"f": []byte("f Response"),
		"g": []byte("g Response"),
		"h": []byte("h Response"),
		"i": []byte("i Response"),
	}

	repository := mockRepositoryForParallel{
		Response: values,
	}

	var urls []string
	for v := range values {
		urls = append(urls, v)
	}

	md5Service := NewMD5HashingService(repository)

	hashes, err := md5Service.HashParallel(urls, 10)
	if err != nil {
		t.Error(err)
	}

	if len(hashes) != len(values) {
		t.Errorf("hashes length = %d; expected %d\n", len(hashes), len(values))
	}

	for _, h := range hashes {
		if h.Hash != fmt.Sprintf("%x", md5.Sum([]byte(values[h.URL]))) {
			t.Errorf("hash value = %s; expected %s\n", hashes[0].Hash, fmt.Sprintf("%x", md5.Sum([]byte(values[h.URL]))))
		}
	}
}
