package home_controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/schweigert/inscreveai/view/home_view"
)

func IndexHandler(c *gin.Context) {
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(home_view.Index()))
}
