package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cache/persistence"
	"github.com/gin-gonic/gin"
)

var store = persistence.NewInMemoryStore(time.Minute * 5)

func main() {
	r := gin.Default()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	r.Static("/file", "uploads")
	r.GET("/", HelloWorldIndexHandler)
	r.GET("/get/:key", GetFileHandler)
	r.POST("/upload-one", UploadOneFileHandler)
	// r.POST("/upload-many", UploadMultipleFileHandler)

	// start server
	APP_PORT := os.Getenv("PORT")
	if APP_PORT == "" {
		APP_PORT = "9090"
	}
	if err := r.Run(":" + APP_PORT); err != nil {
		log.Fatal(err)
	}
}

// HelloWorldIndexHandler is a simple health check endpoint
func HelloWorldIndexHandler(ctx *gin.Context) {
	ctx.File("views/index.html")
}

func GetFileHandler(ctx *gin.Context) {
	key := ctx.Param("key")
	if key == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "key not found",
		})
		return
	}

	// get file from cache
	filename, _ := store.Cache.Get(key)
	// cast filename to string
	filenameStr := filename.(string)

	// check filename len > 2
	if len(filenameStr) < 2 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "key not found",
		})
		return
	}

	// check file exists
	if _, err := os.Stat("uploads/" + filenameStr); os.IsNotExist(err) {
		ctx.JSON(http.StatusNotFound, gin.H{
			"message": "key not found",
		})
		return
	}

	// read and return file from /uploads folder
	ctx.File("uploads/" + filenameStr)
}

func UploadOneFileHandler(ctx *gin.Context) {
	// Multipart form
	file, _ := ctx.FormFile("myfile")

	// check file
	if file == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "file not found",
		})
		return
	}

	err := ctx.SaveUploadedFile(file, "uploads/"+file.Filename)
	if err != nil {
		log.Fatal(err)
	}

	// cache file with random key
	key := "f-" + GenerateRandomString(6)
	store.Set(key, file.Filename, time.Minute*1)

	// return json response
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"type":   "upload-one",
		"file":   file.Filename,
		"key":    key,
	})
	return // return
}

func UploadMultipleFileHandler(ctx *gin.Context) {
	// Multipart form
	form, _ := ctx.MultipartForm()
	files := form.File["myfile"]

	// check file
	if len(files) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "file not found",
		})
		return
	}

	// uploaded_filenames
	uploaded_filenames := make([]string, len(files))

	for _, file := range files {
		log.Println(file.Filename)
		err := ctx.SaveUploadedFile(file, "uploads/"+file.Filename)
		if err != nil {
			log.Fatal(err)
		}
		uploaded_filenames = append(uploaded_filenames, file.Filename)
	}

	// return json response
	ctx.JSON(http.StatusOK, gin.H{
		"status": true,
		"type":   "",
		"files":  uploaded_filenames,
	})
	return // return
}

func GenerateRandomString(n int) string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
