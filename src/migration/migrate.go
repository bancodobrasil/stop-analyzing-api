package migration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/sirupsen/logrus"
)

//Exec function selects the correct migration type (url or filesystem) and then execute the import
func Exec(path string, recreate bool) error {

	//TODO: drop db on if recreate is true

	if isURL(path) {
		return migrateFromURL(path)
	}

	return migrateFromFile(path)
}

func migrateFromURL(url string) error {

	logrus.Infof("Migrating database from URL: %s", url)

	body, err := executeGET(url)

	if err != nil {
		return err
	}

	items, err := parseItems(body)

	if err != nil {
		return err
	}

	return importItems(items)
}

func migrateFromFile(filePath string) error {
	logrus.Infof("Migrating database from file: %s", filePath)
	return nil
}

func parseItems(body []byte) ([]Item, error) {

	items := make([]Item, 0)

	if err := json.Unmarshal(body, &items); err != nil {
		return nil, err
	}

	return items, nil
}

func importItems([]Item) error {
	//TODO: Import items
	return nil
}

func isURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}

func executeGET(url string) ([]byte, error) {

	cli := http.Client{Timeout: 30 * time.Second}
	resp, err := cli.Get(url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Unable to get value, unexpected response status code %d from %s", resp.StatusCode, url)
	}

	return ioutil.ReadAll(resp.Body)
}
