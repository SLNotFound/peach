package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type config struct {
	EditWidget    *widget.Entry
	PreviewWidget *widget.RichText
	CurrentFile   fyne.URI
	SaveMenuItem  *fyne.MenuItem
}

var cfg config

func main() {
	// 创建一个应用
	a := app.New()

	// 为这个应用创建一个窗口
	w := a.NewWindow("Markdown")

	// 获取用户接口
	edit, preview := cfg.makeUI()

	// 设置窗口的内容
	w.SetContent(container.NewHSplit(edit, preview))

	// 展示窗口&运行应用
	w.Resize(fyne.Size{Width: 800, Height: 600})
	//居中显示
	w.CenterOnScreen()
	w.ShowAndRun()
}

func (app *config) makeUI() (*widget.Entry, *widget.RichText) {
	edit := widget.NewMultiLineEntry()
	preview := widget.NewRichTextFromMarkdown("")
	app.EditWidget = edit
	app.PreviewWidget = preview

	edit.OnChanged = preview.ParseMarkdown

	return edit, preview
}
