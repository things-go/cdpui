package cdpui

import (
	"fmt"

	"github.com/chromedp/chromedp"
)

// WindowSize 设置初始窗口大小: 800,600
func WindowSize(width, height int) chromedp.ExecAllocatorOption {
	return chromedp.WindowSize(width, height)
}

// WindowPosition 指定初始窗口位置:-窗口位置=x,y
func WindowPosition(x, y int) chromedp.ExecAllocatorOption {
	return chromedp.Flag("window-size", fmt.Sprintf("%d,%d", x, y))
}

// Maximized 无论先前的设置如何,都启动浏览器最大化
func Maximized() chromedp.ExecAllocatorOption {
	return chromedp.Flag("start-maximized", true)
}

// Fullscreen 全屏显示
func Fullscreen() chromedp.ExecAllocatorOption {
	return chromedp.Flag("start-fullscreen", true)
}
