package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"gocv.io/x/gocv"
)

func main() {
	a := app.New()
	w := a.NewWindow("Background Image")
	w.Resize(fyne.NewSize(600, 600))
	webcam, _ := gocv.OpenVideoCapture(0)
	imgen := gocv.NewMat()
	webcam.Read(&imgen)
	img2, _ := imgen.ToImage()

	img := canvas.NewImageFromImage(img2)
	img.FillMode = canvas.ImageFillOriginal

	w.SetContent(container.New(layout.NewVBoxLayout(), img, widget.NewButton("Quit", a.Quit)))
	w.CenterOnScreen()
	w.Show()
	w.ShowAndRun()

}
