package migration

//Item represents a migration (file and url) entry point for an Item entity
type Item struct {
	Title      string   `json:"title"`
	Subtitle   string   `json:"subtitle"`
	ContentURL string   `json:"contentURL"`
	Tags       []string `json:"tags"`
}
