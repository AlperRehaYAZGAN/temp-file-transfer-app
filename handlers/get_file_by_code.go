package handlers

import (
	"net/http"
	"os"

	"github.com/AlperRehaYAZGAN/temp-file-transfer-app/config"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

// RetrieveFile godoc
// @Summary retrieve uploaded file by temp code
// @Schemes
// @Description retrieve uploaded file by temp code
// @Tags Upload
// @Accept */*
// @Param code path string true "code"
// @Produce */*
// @Success 200 {object} string "file itself"
// @Failure 404 {object} handlers.RespondJson "File expired or not found"
// @Failure 500 {object} handlers.RespondJson "Internal Server Error"
// @Router /f/{code} [get]
func (us *UploadService) HandleGetFileByCode(c *gin.Context) {
	endpointUrn := "ykt:::tempfileupload:::/f/{code}"
	code := c.Param("code")

	// cache key
	key := CACHE_NS + ":" + code

	// get filename from from cache if key is valid
	var filename string
	err := CacheGet(us.inAppCache, key, &filename)
	if err != nil {
		respondJson(c, http.StatusNotFound, endpointUrn, "File expired or not found", nil)
		return
	}

	// check file exists
	path := config.Pwd + "/" + config.C.App.UploadsDir + "/" + code

	if _, err := os.Stat(path); os.IsNotExist(err) {
		respondJson(c, http.StatusNotFound, endpointUrn, "file not found", nil)
		return
	}

	// read and return file from /uploads folder; set filename as header
	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.File(path)
}

func CacheGet(inAppCache *persistence.InMemoryStore, key string, dest interface{}) error {
	return inAppCache.Get(key, dest)
}
