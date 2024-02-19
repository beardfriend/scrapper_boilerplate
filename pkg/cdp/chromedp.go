package cdp

import (
	"context"

	"github.com/chromedp/chromedp"
)

type ChromeDP struct {
	context context.Context
	cancel  context.CancelFunc
	opts    []func(*chromedp.ExecAllocator)
}

func NewChromeDP() *ChromeDP {
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.DisableGPU,
		chromedp.WindowSize(1980, 1080),
		// chromedp.UserDataDir("requests"),
	)

	return &ChromeDP{
		opts: opts,
	}
}

func (c *ChromeDP) HeadlessFalse() *ChromeDP {
	c.opts = append(c.opts,
		chromedp.Flag("headless", false),
	)

	return c
}

func (c *ChromeDP) HeadlessTrue() *ChromeDP {
	c.opts = append(c.opts,

		chromedp.Flag("headless", true),
	)

	return c
}

func (c *ChromeDP) Init(ctx context.Context) {
	allocCtx, cancel := chromedp.NewExecAllocator(ctx, c.opts...)
	c.cancel = cancel

	c.context, _ = chromedp.NewContext(allocCtx)
}

func (c *ChromeDP) GetContext() context.Context {
	return c.context
}

func (c *ChromeDP) Close() {
	c.cancel()
}
