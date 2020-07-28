package migration

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestShouldParseSimpleItemPayload(t *testing.T) {

	resp := `[{
		"title": "Item 1",
		"subtitle": "First item to store",
		"contentURL": "https://via.placeholder.com/300",
		"tags": ["tag-1", "tag-2","tag-3"]
	},
	{
		"title": "Item 2",
		"subtitle": "Second item to store",
		"contentURL": "https://via.placeholder.com/400",
		"tags": ["tag-2", "tag-3", "tag-4"]
    }]`

	parsed, err := parseItems([]byte(resp))

	if err != nil {
		t.Errorf("Error while parsing items: %s", err)
		return
	}

	if len(parsed) != 2 {
		t.Errorf("Error while parsing items, expected size: 2, received: %d", len(parsed))
		return
	}

	if len(parsed[0].Tags) != 3 {
		t.Errorf("Error while parsing tags, expected size: 3, received: %d", len(parsed))
		return
	}
}

func TestShouldReturnTrueForValidURL(t *testing.T) {

	var validURLs = [...]string{
		"http://teste.com",
		"http://teste.com:9000",
		"http://teste.com:9000/",
		"http://teste.com:9000/API",
		"https://teste.com:9000",
		"https://teste.com:9000/",
		"https://teste.com:9000/api",
		"https://www.teste.com/api?query=true",
	}

	for _, url := range validURLs {

		t.Run(fmt.Sprintf("Test for url %s", url), func(t *testing.T) {

			if !isURL(url) {
				t.Errorf("Should return true for valid url %s", url)
			}
		})
	}
}

func TestShouldReturnFalseForFileSystemOrMalformedUrl(t *testing.T) {

	var invalidURLs = [...]string{
		"http/teste.com",
		"/home/test.json",
		"~/teste.json",
		"teste.json",
	}

	for _, url := range invalidURLs {

		t.Run(fmt.Sprintf("Test for url %s", url), func(t *testing.T) {

			if isURL(url) {
				t.Errorf("Should return false for invalid url or filesystem %s", url)
			}
		})
	}
}

func TestShouldReturnErrorForNonOKStatusCode(t *testing.T) {

	var commonNonOkStatusCode = [...]int{
		201,
		202,
		203,
		400,
		401,
		403,
		404,
		500,
		501,
	}

	for _, code := range commonNonOkStatusCode {

		t.Run(fmt.Sprintf("Test for code %d", code), func(t *testing.T) {

			ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Error(w, "Any error", code)
			}))

			_, err := executeGET(ts.URL)

			if err == nil {
				t.Error("Should return error for non 200 status code")
			}
		})
	}
}
