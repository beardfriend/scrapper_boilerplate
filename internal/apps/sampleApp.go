package apps

import (
	"boilerplate/internal/modules/directory"
	"boilerplate/internal/modules/meta"
	"context"
)

type SampleApp struct {
	MetaHandler      meta.Handler
	DirectoryHandler directory.Handler
}

func NewSampleApp(meta meta.Handler, directory directory.Handler) *SampleApp {
	return &SampleApp{
		MetaHandler:      meta,
		DirectoryHandler: directory,
	}
}

func (a *SampleApp) GetHandlers() []interface{} {
	return []interface{}{
		a.MetaHandler,
		a.DirectoryHandler,
	}
}

func (a *SampleApp) StartUp(ctx context.Context) {
	a.MetaHandler.Init(ctx)
	a.DirectoryHandler.Init(ctx)
}
