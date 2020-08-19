package v1

import (
	"net/http"

	"github.com/bancodobrasil/stop-analyzing-api/internal/db"

	"github.com/gin-gonic/gin"
)

//Dependencies
type TagRepository interface {
	GetAllTags() ([]db.TagModel, error)
}

//Controller .
type Controller struct {
	tagRepository TagRepository
}

//InitRoutesV1 Initialize Route V1
func InitRoutesV1(tagRepository TagRepository) Controller {
	return Controller{tagRepository: tagRepository}
}

//Index .
func (c *Controller) Index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Version 1 API")
}

//Choice .
func (c *Controller) Choice(ctx *gin.Context) {

	// Placeholder - Replace with algorithm

	payload :=
		[]byte(`{
	"completeness": "0.3",
	"choices": [
		{
			"id": "5f355a24ccb4180025ee98ab",
			"title": "Fashion Shirt",
			"subtitle": "Colored nice Shirt",
			"contentURL": "https://lorempixel.com/640/480/",
			"attributes": [
				{
					"size": "S",
					"color": "multiple"
				}
			]
		},
		{
			"id": "5f358359ccb4180025ee98ad",
			"title": "Corporate Shirt",
			"subtitle": "Corporate long sleeve Shirt",
			"contentURL": "https://lorempixel.com/640/480/",
			"attributes": [
				{
					"size": "S",
					"color": "black"
				}
			]
		}
	]
}`)
	ctx.Data(http.StatusOK, "application/json", payload)
}
