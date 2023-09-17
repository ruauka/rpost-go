package ui

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"r-post/parse"
)

const (
	mLeft  = 200
	length = 500
)

func UrlDesc() *canvas.Text {
	text := canvas.NewText("Url", color.White)
	text.Alignment = fyne.TextAlignTrailing
	//text.TextStyle = fyne.TextStyle{Italic: true}
	text.Resize(fyne.NewSize(40, 50))
	text.Move(fyne.NewPos(50, 43))

	return text
}

func Url() *widget.Entry {
	url := widget.NewEntry()
	url.SetText("localhost:8080/execute")
	url.Resize(fyne.NewSize(600, 40))
	url.Move(fyne.NewPos(105, 50))

	return url
}

func FileBtn(filePath *widget.Entry, w fyne.Window) *widget.Button {
	file := widget.NewButton("Browse", func() {
		dialog.ShowFileOpen(func(f fyne.URIReadCloser, _ error) {
			filePath.SetText(f.URI().Path())
		}, w)
	})
	file.Resize(fyne.NewSize(70, 40))
	file.Move(fyne.NewPos(30, 100))

	return file
}

func FilePath() *widget.Entry {
	pathFile := widget.NewEntry()
	pathFile.SetPlaceHolder("path to json...")
	pathFile.Resize(fyne.NewSize(600, 40))
	pathFile.Move(fyne.NewPos(105, 100))

	return pathFile
}

func SendBtn(response, url, filePath *widget.Entry) *widget.Button {
	btn := widget.NewButton("Send", func() {
		if url.Text == "" {
			response.SetText("Не указан URL сервиса!")
			return
		}
		if filePath.Text == "" {
			response.SetText("Не указан путь к json!")
			return
		}

		res, err := parse.GetResp(url.Text, filePath.Text)
		if err != nil {
			response.SetText(err.Error())
			return
		}

		response.SetText(res)
	})
	btn.Resize(fyne.NewSize(200, 40))
	btn.Move(fyne.NewPos(300, 150))

	return btn
}

func Response() *widget.Entry {
	response := widget.NewEntry()
	response.Resize(fyne.NewSize(600, 700))
	response.Move(fyne.NewPos(105, 200))

	return response
}
