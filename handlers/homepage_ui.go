package handlers

import (
	"github.com/gin-gonic/gin"
)

func (us *UploadService) HandleUiIndex(c *gin.Context) {
	// redirect to /docs/index.html
	c.Redirect(301, "/docs/index.html")
	/*
		// send PWD/templates/index.html
		path := config.Pwd + "/templates/index.html"
		c.File(path)
	*/
}
