package handlers_test

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"testing"

	"github.com/AlperRehaYAZGAN/temp-file-transfer-app/config"
	"github.com/AlperRehaYAZGAN/temp-file-transfer-app/handlers"
)

const (
	ENDPOINT_UPLOAD = "http://localhost:9090/api/v1/upload"
	ENDPOINT_GET    = "http://localhost:9090/api/v1/f"
)

func onServerInit() {
	// parse application envs
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("INIT: Cannot get current working directory os.Getwd()")
	}
	config.Pwd = dir
	config.ReadConfig(dir+"/../", "")
}

var Code string

func TestUploadFileToServer(t *testing.T) {
	// Open the file to read its contents
	onServerInit()
	wd, _ := os.Getwd()
	file, err := os.Open(wd + "/../LICENCE")
	if err != nil {
		t.Errorf("TestUploadFileToServer: Cannot open file: %s", err.Error())
	}
	defer file.Close()

	// Create a buffer to hold the request body
	var body bytes.Buffer

	// Create a new multipart writer
	writer := multipart.NewWriter(&body)

	// Create a new form-data header with the name "file"
	part, err := writer.CreateFormFile("file", file.Name())
	if err != nil {
		t.Errorf("TestUploadFileToServer: Cannot create form file: %s", err.Error())
		return
	}

	// Use io.Copy to copy the file contents to the form-data part
	_, err = io.Copy(part, file)
	if err != nil {
		t.Errorf("TestUploadFileToServer: Cannot copy file contents to form file: %s", err.Error())
		return
	}

	// Close the writer to write the ending boundary
	err = writer.Close()
	if err != nil {
		t.Errorf("TestUploadFileToServer: Cannot close writer: %s", err.Error())
		return
	}

	// Use the http.Post function to post the request
	resp, err := http.Post(ENDPOINT_UPLOAD, writer.FormDataContentType(), &body)
	if err != nil {
		t.Errorf("TestUploadFileToServer: Cannot post request: %s", err.Error())
		return
	}
	defer resp.Body.Close()

	// stringify body
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Error("TestUploadFileToServer() could not read response body: ", err)
		return
	}

	// response.Body should be a json string from struct handlers.RespondJson
	var respondJson handlers.RespondJson
	err = json.Unmarshal(bodyBytes, &respondJson)
	if err != nil {
		t.Error("TestUploadFileToServer() could not unmarshal response body: ", err)
		return
	}

	// Check the response
	if resp.StatusCode != http.StatusOK {
		t.Errorf("TestUploadFileToServer() resp.StatusCode = %v, want %v", resp.StatusCode, http.StatusOK)
		return
	}

	if respondJson.Status == false {
		t.Errorf("TestUploadFileToServer() respondJson.Status = %v, want %v", respondJson.Status, true)
		return
	}

	// RespondJson.Message should be a code as string not interface
	Code = respondJson.Message.(string)
	t.Logf("TestUploadFileToServer() Code = %v", Code)
}

func TestGetUploadedFile(t *testing.T) {
	t.Logf("TestGetUploadedFile() Code = %v", Code)
	// Create a new request using http to get the file
	req, err := http.NewRequest("GET", ENDPOINT_GET+"/"+Code, nil)
	if err != nil {
		t.Errorf("TestGetUploadedFile: Cannot create new request: %s", err.Error())
		return
	}

	// Send the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("TestGetUploadedFile: Cannot send request: %s", err.Error())
		return
	}

	// Check the response
	if resp.StatusCode != http.StatusOK {
		t.Errorf("TestGetUploadedFile() resp.StatusCode = %v, want %v", resp.StatusCode, http.StatusOK)
		return
	}
}

func TestUploadLargeFileToServer(t *testing.T) {
	onServerInit()
	// Open the file to read its contents
	wd, _ := os.Getwd()
	file, err := os.Open(wd + "/../tests/big-size.mp4")
	if err != nil {
		t.Errorf("TestUploadLargeFileToServer: Cannot open file: %s", err.Error())
	}
	defer file.Close()

	// Create a buffer to hold the request body
	var body bytes.Buffer

	// Create a new multipart writer
	writer := multipart.NewWriter(&body)

	// Create a new form-data header with the name "file"
	part, err := writer.CreateFormFile("file", file.Name())
	if err != nil {
		t.Errorf("TestUploadLargeFileToServer: Cannot create form file: %s", err.Error())
		return
	}

	// Use io.Copy to copy the file contents to the form-data part
	_, err = io.Copy(part, file)
	if err != nil {
		t.Errorf("TestUploadLargeFileToServer: Cannot copy file contents to form file: %s", err.Error())
		return
	}

	// Close the writer to write the ending boundary
	err = writer.Close()
	if err != nil {
		t.Errorf("TestUploadLargeFileToServer: Cannot close writer: %s", err.Error())
		return
	}

	// Use the http.Post function to post the request
	resp, err := http.Post(ENDPOINT_UPLOAD, writer.FormDataContentType(), &body)
	if err != nil {
		t.Errorf("TestUploadLargeFileToServer: Cannot post request: %s", err.Error())
		return
	}
	defer resp.Body.Close()

	// Check the response
	if resp.StatusCode != http.StatusUnprocessableEntity {
		t.Errorf("TestUploadLargeFileToServer() resp.StatusCode = %v, want %v", resp.StatusCode, http.StatusUnprocessableEntity)
		return
	}
}

func TestUploadExeFileToServer(t *testing.T) {
	onServerInit()
	// Open the file to read its contents
	wd, _ := os.Getwd()
	file, err := os.Open(wd + "/../tests/not-allowed.exe")
	if err != nil {
		t.Errorf("TestUploadExeFileToServer: Cannot open file: %s", err.Error())
	}
	defer file.Close()

	// Create a buffer to hold the request body
	var body bytes.Buffer

	// Create a new multipart writer
	writer := multipart.NewWriter(&body)

	// Create a new form-data header with the name "file"
	part, err := writer.CreateFormFile("file", file.Name())
	if err != nil {
		t.Errorf("TestUploadExeFileToServer: Cannot create form file: %s", err.Error())
		return
	}

	// Use io.Copy to copy the file contents to the form-data part
	_, err = io.Copy(part, file)
	if err != nil {
		t.Errorf("TestUploadExeFileToServer: Cannot copy file contents to form file: %s", err.Error())
		return
	}

	// Close the writer to write the ending boundary
	err = writer.Close()
	if err != nil {
		t.Errorf("TestUploadExeFileToServer: Cannot close writer: %s", err.Error())
		return
	}

	// Use the http.Post function to post the request
	resp, err := http.Post(ENDPOINT_UPLOAD, writer.FormDataContentType(), &body)
	if err != nil {
		t.Errorf("TestUploadExeFileToServer: Cannot post request: %s", err.Error())
		return
	}
	defer resp.Body.Close()

	// Check the response
	if resp.StatusCode != http.StatusTeapot {
		t.Errorf("TestUploadExeFileToServer() resp.StatusCode = %v, want %v", resp.StatusCode, http.StatusTeapot)
		return
	}
}
