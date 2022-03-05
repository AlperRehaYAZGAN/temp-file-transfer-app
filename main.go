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

// App current working directory
var pwd string

var store = persistence.NewInMemoryStore(time.Minute * 5)

func main() {
	// get current working directory
	pwd, _ = os.Getwd()

	r := gin.New()
	r.MaxMultipartMemory = 8 << 20 // 8 MiB

	r.Static("/static", pwd+"/static")
	r.LoadHTMLFiles(pwd+"/views/index.html", pwd+"/views/file.html")

	// Group ui routes
	ui := r.Group("/")
	{
		ui.GET("/", HomePageHandler)                  // home page
		ui.GET("/f/:filename", GetFileByKeyUIHandler) // try to get file by normal
	}

	// Group api routes
	api := r.Group("/api")
	{
		api.GET("/:key/:filename", GetFileByKeyHandler) // try to get file with key
		api.POST("/:filename", UploadOneFileHandler)    // upload one files
		api.PUT("/:filename", UploadOneFileHandler)     // upload one files
	}

	// start server
	APP_PORT := os.Getenv("PORT")
	if APP_PORT == "" {
		APP_PORT = "9090"
	}
	if err := r.Run(":" + APP_PORT); err != nil {
		log.Fatal(err)
	}
}

// HomePageHandler is a home page handler
func HomePageHandler(ctx *gin.Context) {
	ctx.File(pwd + "/views/index.html")
}

// GetFileByKeyUIHandler is a key protected file display ui handler
func GetFileByKeyUIHandler(ctx *gin.Context) {
	// 1 - get filename from param, key from query
	filename := ctx.Param("filename")
	key := ctx.Query("key")

	// 5 - render template with data
	ctx.HTML(http.StatusOK, "file.html", gin.H{
		"filename": filename,
		"key":      key,
	})
}

func GetFileByKeyHandler(ctx *gin.Context) {
	key := ctx.Param("key")
	if key == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "key not found",
		})
		return
	}

	// get file from cache
	filename, _ := store.Cache.Get(key)

	// cast filename to string if not nil
	filenameStr := ""
	if filename != nil {
		filenameStr = filename.(string)
	}

	// check filename len > 3
	if len(filenameStr) < 3 {
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
	// get current requested hostname and scheme (http/https)
	hostname := ctx.Request.Host
	scheme := "http"
	if ctx.Request.TLS != nil {
		scheme = "https"
	}
	// get file from request
	filename := ctx.Param("filename")
	// Multipart form
	file, _ := ctx.FormFile("myfile")

	// check file
	if file == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "file not found",
		})
		return
	}

	err := ctx.SaveUploadedFile(file, "uploads/"+filename)
	if err != nil {
		log.Fatal(err)
	}

	// cache file with random key
	key := "f-" + GenerateRandomString(6)
	store.Set(key, file.Filename, time.Minute*1)

	// return json response
	ctx.JSON(http.StatusOK, gin.H{
		"status":         true,
		"type":           "upload-one",
		"file_real_name": file.Filename,
		"file_save_name": filename,
		"key":            key,
		"api_url":        "" + scheme + "://" + hostname + "/api/" + key + "/" + filename,
		"ui_url":         "" + scheme + "://" + hostname + "/f/" + filename + "?key=" + key,
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
