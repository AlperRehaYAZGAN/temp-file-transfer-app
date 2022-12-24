package handlers

import (
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/AlperRehaYAZGAN/temp-file-transfer-app/config"
	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

const (
	MAX_FILE_SIZE = 1024 * 1024 * 10 // 10MB
	CACHE_NS      = "ykt"
)

// UploadService is a struct for auth core
type UploadService struct {
	inAppCache *persistence.InMemoryStore
}

func NewUploadService(
	inAppCache *persistence.InMemoryStore,
) *UploadService {
	return &UploadService{
		inAppCache: inAppCache,
	}
}

type RespondJson struct {
	Status  bool        `json:"status"`
	Intent  string      `json:"intent"`
	Message interface{} `json:"message"`
}

func RespondJsonCall(ctx *gin.Context, code int, intent string, message interface{}, err error) {
	respondJson(ctx, code, intent, message, err)
}

func respondJson(ctx *gin.Context, code int, intent string, message interface{}, err error) {
	if err == nil {
		ctx.JSON(code, RespondJson{
			Status:  true,
			Intent:  intent,
			Message: message,
		})
	} else {
		ctx.JSON(code, RespondJson{
			Status:  false,
			Intent:  intent,
			Message: err.Error(),
		})
	}
}

func (us *UploadService) InitRouter(r *gin.Engine) *gin.Engine {
	// -- upload-file routes (group)

	// set public folder
	r.Static("/public", "./public")

	// ping for r
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// ui routes
	ui := r.Group("/")
	ui.GET("/", us.HandleUiIndex)
	// ui.GET("/f/:filename", us.HandleUiGetFileByName)

	// api routes

	// get uploaded file
	api := r.Group("/api/v1/")
	api.GET("/f/:code", us.HandleGetFileByCode)
	// upload file (one per request)
	api.POST("/upload", func(ctx *gin.Context) {
		code, data, err := us.HandleUploadFile(ctx)
		respondJson(ctx, code, "ykt:::tempfileupload:::/upload", data, err)
	})

	return r
}

func (us *UploadService) InitUploadsOldFileCleaner() {
	// remove old files with time.Ticker
	ticker := time.NewTicker(1 * time.Minute)
	go func() {
		for range ticker.C {
			// remove old files
			files, err := ioutil.ReadDir(config.Pwd + "/" + config.C.App.UploadsDir)
			if err != nil {
				log.Println("Error while reading uploads dir: ", err)
			}

			cleaned := 0
			for _, f := range files {
				if time.Since(f.ModTime()).Minutes() > 1 {
					err = os.Remove(config.Pwd + "/" + config.C.App.UploadsDir + "/" + f.Name())
					if err != nil {
						log.Println("Error while removing old file: ", err)
					}

					cleaned++
				}
			}

			if cleaned > 0 {
				// current time like: 13:59
				now := time.Now().Format("15:04")

				// log cleanup
				log.Println("Cleaned old files: ", now, " - cleaned: ", cleaned)
			}
		}
	}()

}
