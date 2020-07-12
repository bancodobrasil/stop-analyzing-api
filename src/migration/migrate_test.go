package migration

import (
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
