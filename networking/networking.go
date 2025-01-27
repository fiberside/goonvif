package networking

import (
	"bytes"
	"net/http"
)

func SendSoap(endpoint, message string) (*http.Response, error) {
	httpClient := new(http.Client)

	resp, err := httpClient.Post(endpoint, "application/soap+xml; charset=utf-8", bytes.NewBufferString(message))
	if err != nil {
		return resp, err
	}

	return resp, nil
}
