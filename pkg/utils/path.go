package utils

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetPath() string {

	mode := os.Getenv("GO_ENV")
	if mode == "DEV" || mode == "DEBUG" || mode == "TEST" {
		return "."
	}

	exePath, _ := os.Executable()
	absPath, _ := filepath.EvalSymlinks(exePath)

	path := filepath.Join(filepath.Dir(absPath))
	if runtime.GOOS == "darwin" {
		path = filepath.Join((filepath.Dir(path)))
	}

	return path

}
