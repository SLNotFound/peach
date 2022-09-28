package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"io/ioutil"
	"os"
)

type config struct {
	EditWidget    *widget.Entry
	PreviewWidget *widget.RichText
	CurrentFile   fyne.URI
	SaveMenuItem  *fyne.MenuItem
}

var cfg config

func main() {
	os.Setenv("FYNE_FONT", "msyhl.ttc")
	// 创建一个应用
	a := app.New()

	// 为这个应用创建一个窗口
	w := a.NewWindow("Markdown")

	// 获取用户接口
	edit, preview := cfg.makeUI()
	cfg.createMenuItems(w)

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

func (app *config) createMenuItems(win fyne.Window) {
	openMenuItem := fyne.NewMenuItem("打开", app.openFunc(win))

	saveMenuItem := fyne.NewMenuItem("保存", func() {

	})
	app.SaveMenuItem = saveMenuItem
	app.SaveMenuItem.Disabled = true

	saveAsMenuItem := fyne.NewMenuItem("另存为", app.saveAsFunc(win))

	fileMenu := fyne.NewMenu("文件", openMenuItem, saveMenuItem, saveAsMenuItem)

	menu := fyne.NewMainMenu(fileMenu)

	win.SetMainMenu(menu)
}

func (app *config) openFunc(win fyne.Window) func() {
	return func() {
		openDialog := dialog.NewFileOpen(func(read fyne.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			if read == nil {
				return
			}

			defer read.Close()

			data, err := ioutil.ReadAll(read)
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			app.EditWidget.SetText(string(data))
			app.CurrentFile = read.URI()
			win.SetTitle(win.Title() + " - " + read.URI().Name())
			app.SaveMenuItem.Disabled = false
		}, win)

		openDialog.Show()
	}
}

func (app *config) saveAsFunc(win fyne.Window) func() {
	return func() {
		savaDialog := dialog.NewFileSave(func(write fyne.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, win)
				return
			}

			if write == nil {
				// user cancelled
				return
			}

			// save file
			write.Write([]byte(app.EditWidget.Text))
			app.CurrentFile = write.URI()

			defer write.Close()

			win.SetTitle(win.Title() + " - " + write.URI().Name())
			app.SaveMenuItem.Disabled = false

		}, win)
		savaDialog.Show()
	}
}
