package handlers

import (
	"context"
	"errors"
	"strings"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type DirectoryHandler struct {
}

func NewDirectoryHandler() *DirectoryHandler {
	return &DirectoryHandler{}
}

func (h *DirectoryHandler) SelectExcelFile(ctx context.Context) (filePath string, err error) {
	filePath, err = runtime.OpenFileDialog(ctx, runtime.OpenDialogOptions{})
	if err != nil {
		return
	}

	if filePath == "" {
		return
	}

	if !strings.Contains(filePath, ".xlsx") {
		return "", errors.New("엑셀 파일이 아닙니다")
	}

	return
}
