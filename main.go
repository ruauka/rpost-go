package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	"r-post/ui"
)

func main() {
	a := app.New()

	w := a.NewWindow("R-Post")
	w.Resize(fyne.NewSize(800, 950))

	urlDesc := ui.UrlDesc()
	url := ui.Url()

	filePath := ui.FilePath()
	fileBtn := ui.FileBtn(filePath, w)

	response := ui.Response()
	sendBtn := ui.SendBtn(response, url, filePath)

	cont := container.NewWithoutLayout(url, sendBtn, response, filePath, fileBtn, urlDesc)

	w.SetContent(cont)

	w.ShowAndRun()
}

// 15 - размеры
//16 - меню, file, save
//19 - split
//21 - прокрутка
//28 вкладки
//26 многооконностть
//32, 33 менеджер файлрв

// https://github.com/gotk3/gotk3-examples/blob/master/gtk-examples/titlemenu/title-menu.go
// https://github.com/gotk3/gotk3-examples/blob/master/gtk-examples/treeselection/treeselection.go
