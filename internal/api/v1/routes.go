package v1

import (
	"github.com/bancodobrasil/stop-analyzing-api/internal/db"
	"net/http"

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
