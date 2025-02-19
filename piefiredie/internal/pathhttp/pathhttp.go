package pathhttp

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func Get(url string) (string, error) {
	r, err := myClient.Get(url)
	if err != nil {
		return "", err
	}
	defer r.Body.Close()

	if r.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			return "", err
		}
		bodyString := string(bodyBytes)
		return bodyString, nil
	}

	return "", errors.New(fmt.Sprintf("net/http: %d", r.StatusCode))
}
