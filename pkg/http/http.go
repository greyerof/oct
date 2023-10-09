package http

import (
	"fmt"
	"io"
	gohttp "net/http"
)

// Returns the body content from a url query response.
func GetHTTPBody(url string) ([]uint8, error) {
	//nolint:gosec
	resp, err := gohttp.Get(url)
	if err != nil {
		return nil, fmt.Errorf("http request %s failed with error: %w", url, err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading body from %s: %w, body: %s", url, err, string(body))
	}
	return body, nil
}
