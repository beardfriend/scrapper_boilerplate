//go:build wireinject
// +build wireinject

package wire

import (
	"boilerplate/internal/apps"
	"boilerplate/internal/modules/directory"
	"boilerplate/internal/modules/meta"

	"github.com/google/wire"
)

func InitializeSampleApp(m *meta.Meta) *apps.SampleApp {
	wire.Build(
		// handler
		directory.NewHandler,
		meta.NewHandler,

		// app
		apps.NewSampleApp,
	)
	return &apps.SampleApp{}
}
