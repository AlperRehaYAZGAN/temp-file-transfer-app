package handlers

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/AlperRehaYAZGAN/temp-file-transfer-app/config"
	"github.com/AlperRehaYAZGAN/temp-file-transfer-app/utils"
	"github.com/gin-gonic/gin"
)

// UploadFile godoc
// @Summary upload temporary file
// @Schemes
// @Description upload temporary file
// @Tags Upload
// @Accept multipart/form-data
// @Param file formData file true "file"
// @Produce */*
// @Success 200 {object} string "uploaded file temp code"
// @Failure 401 {object} handlers.RespondJson "Unauthorized"
// @Failure 403 {object} handlers.RespondJson "Forbidden"
// @Failure 500 {object} handlers.RespondJson "Internal Server Error"
// @Router /upload [post]
func (us *UploadService) HandleUploadFile(ctx *gin.Context) (int, interface{}, error) {
	// Multipart form
	file, err := ctx.FormFile("file")
	if err != nil {
		return http.StatusBadRequest, nil, errors.New("please provide a file")
	}

	// check file size < MAX_FILE_SIZE
	max := int64(MAX_FILE_SIZE)
	if file.Size > max {
		return http.StatusUnprocessableEntity, nil, errors.New("file size must be less than 10MB")
	}

	// generate random filename and save
	random := utils.GenerateRandomDigitString(8)
	to := config.Pwd + "/" + config.C.App.UploadsDir + "/" + random
	err = ctx.SaveUploadedFile(file, to)
	if err != nil {
		log.Println("Error while saving file: ", err)
		return http.StatusInternalServerError, nil, errors.New("error while saving file")
	}

	// generate key and save to cache
	key := CACHE_NS + ":" + random
	// cache file with random key
	err = us.inAppCache.Set(key, file.Filename, time.Minute)
	if err != nil {
		log.Println("Error while caching file: ", err)
		return http.StatusInternalServerError, nil, errors.New("error while caching file")
	}

	return http.StatusOK, random, nil
}
