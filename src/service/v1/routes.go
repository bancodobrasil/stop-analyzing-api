package v1

import (
	"fmt"
	"net/http"

	"github.com/bancodobrasil/stop-analyzing-api/db"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//Controller .
type Controller struct {
	database *db.DatabasePrisma
}

//InitRoutesV1 Initialize Route V1
func InitRoutesV1(database db.DatabasePrisma) Controller {
	return Controller{database: &database}
}

//Index .
func (c *Controller) Index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Version 1 API")
}

//ListAllTags .
func (c *Controller) ListAllTags(ctx *gin.Context) {
	tags, err := c.database.GetAllTags()
	if err != nil {
		logrus.Errorf("%s", err)
		ctx.String(http.StatusBadRequest, fmt.Sprintf("Error on list all tags: %s", err))
		return
	}
	ctx.JSON(http.StatusOK, tags)
}