package event_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteHandler(c *gin.Context) {
	c.Redirect(http.StatusFound, "/")
}
