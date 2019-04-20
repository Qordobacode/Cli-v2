package general

import (
	"io"
	"net/http"
	"time"
)

const (
	// ApplicationJsonType used in Http header 'Content-Type'
	ApplicationJsonType = "application/json"
)

var (
	// HTTPClient - custom one with a delay set
	HTTPClient = http.Client{
		Timeout: time.Minute * 1,
	}
)

func PostToServer(qordoba *Config, pushFileURL string, reader io.Reader) (*http.Response, error) {
	request, err := http.NewRequest("POST", pushFileURL, reader)

	if err != nil {
		return nil, err
	}
	request.Header.Add("x-auth-token", qordoba.Qordoba.AccessToken)
	request.Header.Add("Content-Type", ApplicationJsonType)
	return HTTPClient.Do(request)
}