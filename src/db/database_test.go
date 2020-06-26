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
