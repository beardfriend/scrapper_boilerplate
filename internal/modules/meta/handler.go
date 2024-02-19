package meta

import "context"

type Handler interface {
	Init(ctx context.Context)
	GetAppMeta() *Meta
}

type handler struct {
	ctx  context.Context
	Meta *Meta
}

func NewHandler(meta *Meta) Handler {
	return &handler{Meta: meta}
}

func (h *handler) Init(ctx context.Context) {
	h.ctx = ctx
}

type Meta struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	// 10000 * main
	// 100 * sub
	// 1 * minor
	VersionSum  int    `json:"versionSum"`
	VersionText string `json:"versionText"`
	CreatedAt   string `json:"createdAt"`
}

func (h *handler) GetAppMeta() *Meta {
	return h.Meta
}
