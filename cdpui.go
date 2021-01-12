package cdpui

import (
	"context"
	"strings"

	"github.com/chromedp/chromedp"
)

// UI ui
type UI struct {
	url    string
	ctx    context.Context
	cancel context.CancelFunc
}

// Config 配置参数
type Config struct {
	flags []chromedp.ExecAllocatorOption
	opts  []chromedp.ContextOption
}

// Option config option
type Option func(c *Config)

// WithExecAllocatorFlags with some exec allocator options.
func WithExecAllocatorFlags(flags ...chromedp.ExecAllocatorOption) Option {
	return func(c *Config) {
		c.flags = flags
	}
}

// WithContextOptions with some context options.
func WithContextOptions(opts ...chromedp.ContextOption) Option {
	return func(c *Config) {
		c.opts = opts
	}
}

// New new a ui with option flags
func New(url string, opts ...Option) UI {
	c := new(Config)
	for _, opt := range opts {
		opt(c)
	}
	if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
		url = "http://" + url
	}

	flags := append(chromedp.DefaultExecAllocatorOptions[:], chromedp.DisableGPU,
		chromedp.Flag("headless", false), // 不使用headless
		chromedp.Flag("incognito", true), // 使用隐身模式
		chromedp.Flag("app", url),
		chromedp.Flag("enable-automation", false), // 取消通知用户他们的浏览器是由自动测试控制的
		chromedp.Flag("disable-infobars", true),   // 不显示控制信息提示
	)
	flags = append(flags, c.flags...)

	ctx, cancel := chromedp.NewExecAllocator(context.Background(), flags...)
	ctx, _ = chromedp.NewContext(ctx, c.opts...)
	return UI{url, ctx, cancel}
}

// Run show ui 如果show失败,直接panic
func (sf UI) Run() {
	go func() {
		err := chromedp.Run(sf.ctx, chromedp.Navigate(sf.url), chromedp.WaitVisible(`body`))
		if err != nil && !strings.Contains(err.Error(), "context canceled") {
			panic(err)
		}
	}()
}

// Wait return the chan wait ui close
func (sf UI) Wait() <-chan struct{} {
	return sf.ctx.Done()
}

// Close close ui
func (sf UI) Close() error {
	sf.cancel()
	return nil
}
