package directory

import (
	"context"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type Handler interface {
	Init(ctx context.Context)
	SelectDirectory() (dirPath string, err error)
}

type handler struct {
	appContext context.Context
}

func NewHandler() Handler {
	return &handler{}
}

func (h *handler) Init(ctx context.Context) {
	h.appContext = ctx
}

func (h *handler) SelectDirectory() (dirPath string, err error) {
	dirPath, err = runtime.OpenDirectoryDialog(h.appContext, runtime.OpenDialogOptions{
		CanCreateDirectories: true,
	})

	if err != nil {
		return
	}

	return
}
