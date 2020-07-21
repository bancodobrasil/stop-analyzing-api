package db

import (
	"fmt"
	"os"
	"testing"
)

func TestListAllTags(t *testing.T) {
	configEnvVars()

	dbCli, err := Connect()
	if err != nil {
		fmt.Errorf("Error at test List All Tags - Connect to Database: %s", err)
	}

	defer dbCli.Disconnect()

	dbCli.DropAllTags()

	tags, err := dbCli.GetAllTags()
	if err != nil {
		fmt.Errorf("Error at test List All Tags - Get All Tags: %s", err)
	}

	if len(tags) != 0 {
		t.Errorf("Error while listing all tags, expected len: %d, got: %d", 0, len(tags))
	}

	_, err = dbCli.CreateTag("name1")
	_, err = dbCli.CreateTag("name2")
	_, err = dbCli.CreateTag("name3")

	tags, err = dbCli.GetAllTags()

	if len(tags) != 3 {
		t.Errorf("Error while listing all tags, expected len: %d, got: %d", 3, len(tags))
	}

	dbCli.DropAllTags()
}

func TestShouldCreateNewTag(t *testing.T) {
	configEnvVars()

	dbCli, err := Connect()
	if err != nil {
		t.Error(err)
	}
	defer dbCli.Disconnect()

	expectedName := "TEST_NewTagName"

	dbCli.DeleteTag(expectedName)

	tag, err := dbCli.CreateTag(expectedName)
	if err != nil {
		t.Error(err)
	}

	if tag.Text != expectedName {
		t.Errorf("Error while creating tag, expecting name: %s, got: %s", expectedName, tag.Text)
	}

	dbCli.DeleteTag(expectedName)
}

func TestShouldFetchAndCreateTags(t *testing.T) {
	configEnvVars()

	dbCli, err := Connect()
	if err != nil {
		t.Error(err)
	}
	defer dbCli.Disconnect()

	//Given a scenario of two existing tags and two new ones
	existingNames := []string{"existingOne", "existingTwo"}
	newNames := []string{"newOne", "newTwo"}
	allNames := append(existingNames, newNames...)

	//clean-up tags
	for _, name := range allNames {
		dbCli.DeleteTag(name)
	}

	//create existing ones
	existingTags := make([]TagModel, 0)
	for _, existingName := range existingNames {
		tag, err := dbCli.CreateTag(existingName)
		if err != nil {
			t.Error(err)
		}

		existingTags = append(existingTags, tag)
	}

	//When we execute the fetch or create
	tags, err := dbCli.FetchOrCreateTags(allNames)
	if err != nil {
		t.Error(err)
	}

	//then we should see all valid tags
	for _, name := range allNames {
		_, ok := tags[name]

		if !ok {
			t.Errorf("Expecting existing %s, but it was not found", name)
		}
	}

	//existing with same id

	for _, e := range existingTags {

		fetched, ok := tags[e.Text]

		if !ok {
			t.Errorf("Expecting existing %s, but it was not found", e.Text)
		}

		if fetched.ID != e.ID {
			t.Errorf("Existing tag with wrong id. Expecting %d, but it was not found %d", e.ID, fetched.ID)
		}
	}

	//clean-up
	for _, v := range tags {
		dbCli.DeleteTag(v.Text)
	}
}

func TestShouldCreateNewItem(t *testing.T) {
	configEnvVars()

	dbCli, err := Connect()
	if err != nil {
		t.Error(err)
	}
	defer dbCli.Disconnect()

	itemTitle := "TEST_Title"
	itemSubtitle := "TEST_SubTitle"
	itemContent := "https://content.url_TEST"
	itemTags := []string{"tag1", "tag2"}

	item, err := dbCli.CreateItem(itemTitle, itemSubtitle, itemContent, itemTags)

	if err != nil {
		t.Error(err)
	}

	defer dbCli.DeleteItem(item.ID)
	defer dbCli.DeleteTag("tag1")
	defer dbCli.DeleteTag("tag2")

	fetchedItem, err := dbCli.FetchItem(item.ID)

	if !fetchedItem.Active {
		t.Errorf("Item should be created as active and it was %v", item.Active)
	}

	if fetchedItem.Title != itemTitle {
		t.Errorf("Wrong item title, expected: %s, found: %s", itemTitle, item.Title)
	}

	if fetchedItem.Subtitle != itemSubtitle {
		t.Errorf("Wrong item subtitle, expected: %s, found: %s", itemSubtitle, item.Subtitle)
	}

	if fetchedItem.ContentURL != itemContent {
		t.Errorf("Wrong item content, expected: %s, found: %s", itemContent, item.ContentURL)
	}

	fmt.Println(fetchedItem.Tags())

	if len(fetchedItem.Tags()) != 2 {
		t.Errorf("Wrong item tags, expected size: %d, found: %d", 2, len(fetchedItem.Tags()))
	}
}

func configEnvVars() {
	os.Setenv("DATABASE_URL", "postgresql://user2020:pass2020@localhost:5432/stop-analyzing-api")
}
