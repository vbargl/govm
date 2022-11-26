package testutil

import (
	"net/http"
	"testing"
)

type HttpClient struct {
	T *testing.T

	ExpectedUrl string
	Content     string
}

func (client *HttpClient) Get(url string) (*http.Response, error) {
	if client.ExpectedUrl != url {
		client.T.Errorf("expected url = %s, got %s", client.ExpectedUrl, url)
	}

	return &http.Response{
		Status:     "200 OK",
		StatusCode: 200,
		Body:       NewBufferString(client.Content),
	}, nil
}
