package migration

import "encoding/json"

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
