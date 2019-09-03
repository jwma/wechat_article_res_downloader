package main

import (
	"fmt"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/dialog"
	"fyne.io/fyne/widget"
	"github.com/jwma/wechat_article_res_downloader/pkg/downloader"
)

var (
	downloaderApp    fyne.App
	downloaderWindow fyne.Window
)

func makeProcessingStatus(b *widget.Button) {
	b.SetText("Downloading...")
	b.Disable()
}

func makeFreeStatus(b *widget.Button) {
	b.SetText("Download")
	b.Enable()
}

func showMessage(m string) {
	dialog.ShowInformation("Tips", m, downloaderWindow)
}

func init() {
	downloaderApp = app.New()
	downloaderWindow = downloaderApp.NewWindow("MJ Studio")
	downloaderWindow.Resize(fyne.Size{
		Width:  440,
		Height: 129,
	})
	downloaderWindow.SetFixedSize(true)
}

func main() {
	dirInput := widget.NewEntry()
	dirInput.SetPlaceHolder("Resources will be saved here")
	urlInput := widget.NewEntry()
	urlInput.SetPlaceHolder("Wechat article URL")

	downloadButton := widget.NewButton("Download", func() {})
	downloadButton.OnTapped = func() {
		if dirInput.Text == "" {
			showMessage("Please enter the Save folder")
			makeFreeStatus(downloadButton)
			return
		}
		if urlInput.Text == "" {
			showMessage("Please enter the wechat article URL")
			makeFreeStatus(downloadButton)
			return
		}
		fmt.Println(dirInput.Text)
		fmt.Println(urlInput.Text)

		makeProcessingStatus(downloadButton)
		err := downloader.Download(dirInput.Text, urlInput.Text)
		if err != nil {
			dialog.ShowError(err, downloaderWindow)
			makeFreeStatus(downloadButton)
			return
		}
		showMessage("Successfully saved")
		makeFreeStatus(downloadButton)
	}

	downloaderWindow.SetContent(widget.NewVBox(
		widget.NewForm(&widget.FormItem{
			Text:   "Save folder",
			Widget: dirInput,
		}, &widget.FormItem{
			Text:   "Article URL",
			Widget: urlInput,
		}),
		downloadButton,
	))

	downloaderWindow.ShowAndRun()
}
