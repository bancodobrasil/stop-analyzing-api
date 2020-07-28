package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

//ListAllTags .
func (c *Controller) ListAllTags(ctx *gin.Context) {
	tags, err := c.tagRepository.GetAllTags()
	if err != nil {
		logrus.Errorf("%s", err)
		ctx.String(http.StatusBadRequest, fmt.Sprintf("Error on list all tags: %s", err))
		return
	}
	ctx.JSON(http.StatusOK, tags)
}
