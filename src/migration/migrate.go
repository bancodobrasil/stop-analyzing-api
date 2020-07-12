package migration

import (
	"encoding/json"
	"net/url"

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
	//TODO: execute Get
	return nil

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

func isURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
