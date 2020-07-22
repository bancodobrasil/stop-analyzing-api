package migration

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/bancodobrasil/stop-analyzing-api/db"
	"github.com/sirupsen/logrus"
)

//Do function selects the correct migration type (url or filesystem) and then execute the import
func Do(path string, recreate bool) error {

	db, err := db.Connect()

	if err != nil {
		panic(err)
	}

	defer db.Disconnect()

	importer := databaseImporter{db: &db}

	if recreate {
		if err := importer.Drop(); err != nil {
			return err
		}
	}

	if isURL(path) {
		urlMigrator := urlMigrator{&importer}
		return urlMigrator.Migrate(path)
	}

	fsMigrator := filesystemMigrator{&importer}
	return fsMigrator.Migrate(path)
}

//Migrator migrates from a path source to a specific destination.
type Migrator interface {
	Migrate(source string) error
}

//ItemImporter uses a destination source instance to drop and import items
type ItemImporter interface {
	Import(items []Item) error
	Drop() error
}

type urlMigrator struct {
	importer ItemImporter
}

func (um *urlMigrator) Migrate(url string) error {

	logrus.Infof("Migrating database from URL: %s", url)

	body, err := executeGET(url)

	if err != nil {
		return err
	}

	items, err := parseItems(body)

	if err != nil {
		return err
	}

	return um.importer.Import(items)
}

//filesystemMigrator is a struct to migration from json file.
type filesystemMigrator struct {
	importer ItemImporter
}

/*
Migrate migrate from json file.

-filePath: json file path. Example: ./testdata/migration/test-migration.json

*/
func (fm *filesystemMigrator) Migrate(filePath string) error {
	logrus.Infof("Migrating database from file: %s", filePath)

	jsonFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	items, err := parseItems(jsonFile)
	if err != nil {
		return err
	}

	return fm.importer.Import(items)
}

type databaseImporter struct {
	db *db.DatabasePrisma
}

func (di *databaseImporter) Import(items []Item) error {

	for _, item := range items {
		if _, err := di.db.CreateItem(item.Title, item.Subtitle, item.ContentURL, item.Tags); err != nil {
			return err
		}
	}

	return nil
}

func (di *databaseImporter) Drop() error {
	_, err := di.db.DropAllTags()
	_, err = di.db.DropAllItems()
	return err
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
