package main

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"gocv.io/x/gocv"
)

func main() {
	webcam, _ := gocv.OpenVideoCapture(0)
	a := app.New()
	w := a.NewWindow("Background Image")
	w.Resize(fyne.NewSize(800, 600))
	go func() {
		for range time.Tick(33 * time.Millisecond) {
			imgen := gocv.NewMat()
			webcam.Read(&imgen)
			img2, _ := imgen.ToImage()
			img := canvas.NewImageFromImage(img2)
			clock := container.New(layout.NewMaxLayout(), img)
			w.SetContent(clock)
		}
	}()
	w.CenterOnScreen()
	w.ShowAndRun()

}
