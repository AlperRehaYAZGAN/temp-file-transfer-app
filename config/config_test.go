package config_test

import (
	"os"
	"testing"

	cfg "github.com/AlperRehaYAZGAN/temp-file-transfer-app/config"
)

func TestReadConfig(t *testing.T) {
	// get working dir
	wd, err := os.Getwd()
	if err != nil {
		t.Errorf("ReadConfig() error = %v", err)
		return
	}

	pwd := wd + "/../"

	// call service method
	cfg.ReadConfig(pwd, "sample-test")

	// get config values
	values := cfg.C

	if values.App.Env != "test" {
		t.Errorf("ReadConfig() cfg.C.App.Env = %v, want %v", values.App.Env, "test")
		return
	}

	if values.App.Port != "9090test" {
		t.Errorf("ReadConfig() cfg.C.App.Port = %v, want %v", values.App.Port, "9090test")
		return
	}

	if values.App.Version != "1.0.0test" {
		t.Errorf("ReadConfig() cfg.C.App.Version = %v, want %v", values.App.Version, "1.0.0test")
		return
	}

	if values.App.UploadsDir != "uploadstest" {
		t.Errorf("ReadConfig() cfg.C.App.UploadsDir = %v, want %v", values.App.UploadsDir, "uploadstest")
		return
	}
}
