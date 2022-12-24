package handlers_test

import (
	"os"
	"testing"
)

func TestWriteFile(t *testing.T) {
	// create on-the fly txt file and save to pwd ../uploads/test.txt
	// get working dir
	wd, err := os.Getwd()
	if err != nil {
		t.Errorf("WriteFile() error = %v", err)
		return
	}

	pwd := wd + "/../uploads/"

	// open file for writing
	file, err := os.Create(pwd + "test.txt")
	if err != nil {
		t.Errorf("WriteFile() error = %v", err)
		return
	}

	// write to file
	_, err = file.WriteString("test")
	if err != nil {
		t.Errorf("WriteFile() error = %v", err)
		return
	}

	// close file
	err = file.Close()
	if err != nil {
		t.Errorf("WriteFile() error = %v", err)
		return
	}

	// check file exists
	_, err = os.Stat(pwd + "test.txt")
	if os.IsNotExist(err) {
		t.Errorf("WriteFile() error = %v", err)
		return
	}

	// remove file after test
	err = os.Remove(pwd + "test.txt")
	if err != nil {
		t.Errorf("WriteFile() error = %v", err)
		return
	}
}

func TestReadFile(t *testing.T) {
	// create on-the fly txt file and save to pwd ../uploads/test.txt
	// get working dir
	wd, err := os.Getwd()
	if err != nil {
		t.Errorf("ReadFile() error = %v", err)
		return
	}

	pwd := wd + "/../"
	filename := "LICENCE"

	// check file exists
	_, err = os.Stat(pwd + filename)
	if os.IsNotExist(err) {
		t.Errorf("ReadFile() error = %v", err)
		return
	}

	// read file
	_, err = os.ReadFile(pwd + filename)
	if err != nil {
		t.Errorf("ReadFile() error = %v", err)
		return
	}
}
