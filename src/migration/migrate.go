package migration

import (
	"encoding/json"
)

//Exec function selects the correct migration type (url or filesystem) and then execute the import
func Exec(path string, recreate bool) error {

	return nil
}

func migrateFromURL(url string) {

	//TODO: execute Get

}
func migrateFromFile(filePath string) {

}

func parseItems(body []byte) ([]Item, error) {

	items := make([]Item, 0)

	if err := json.Unmarshal(body, &items); err != nil {
		return nil, err
	}

	return items, nil
}
