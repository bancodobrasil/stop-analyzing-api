package db

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"gotest.tools/assert"
)

func TestListAllTags(t *testing.T) {
	os.Setenv("DATABASE_URL", "postgresql://user2020:pass2020@localhost:5432/stop-analyzing-api")

	dbCli, err := Connect()
	if err != nil {
		fmt.Errorf("Error at test List All Tags - Connect to Database: %s", err)
	}
	defer dbCli.Disconnect()
	tags, err := dbCli.GetAllTags()
	if err != nil {
		fmt.Errorf("Error at test List All Tags - Get All Tags: %s", err)
	}

	expectedResult := []TagModel{}
	assert.Equal(t, true, reflect.DeepEqual(tags, expectedResult))
}

func TestShouldCreateNewTag(t *testing.T) {
	os.Setenv("DATABASE_URL", "postgresql://user2020:pass2020@localhost:5432/stop-analyzing-api")

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

	if tag.TagName != expectedName {
		t.Errorf("Error while creating tag, expecting name: %s, got: %s", expectedName, tag.TagName)
	}

	dbCli.DeleteTag(expectedName)
}

func TestShouldFetchAndCreateTags(t *testing.T) {
	os.Setenv("DATABASE_URL", "postgresql://user2020:pass2020@localhost:5432/stop-analyzing-api")

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

		fetched, ok := tags[e.TagName]

		if !ok {
			t.Errorf("Expecting existing %s, but it was not found", e.TagName)
		}

		if fetched.ID != e.ID {
			t.Errorf("Existing tag with wrong id. Expecting %d, but it was not found %d", e.ID, fetched.ID)
		}
	}

	//clean-up
	for _, v := range tags {
		dbCli.DeleteTag(v.TagName)
	}
}

func TestShouldCreateNewItem(t *testing.T) {
	os.Setenv("DATABASE_URL", "postgresql://user2020:pass2020@localhost:5432/stop-analyzing-api")

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

	if !item.Active {
		t.Errorf("Item should be created as active and it was %v", item.Active)
	}

	if item.Title != itemTitle {
		t.Errorf("Wrong item title, expected: %s, found: %s", itemTitle, item.Title)
	}

	if item.Subtitle != itemSubtitle {
		t.Errorf("Wrong item subtitle, expected: %s, found: %s", itemSubtitle, item.Subtitle)
	}

	if item.ContentURL != itemContent {
		t.Errorf("Wrong item content, expected: %s, found: %s", itemContent, item.ContentURL)
	}

	if len(item.Tags()) != 2 {
		t.Errorf("Wrong item tags, expected size: %d, found: %d", 2, len(item.Tags()))
	}
}
