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

func updateTime(clock *fyne.Container, webcam *gocv.VideoCapture) {
	clock.RemoveAll()
	// clock.Refresh()

	imgen := gocv.NewMat()
	webcam.Read(&imgen)
	img2, _ := imgen.ToImage()

	img := canvas.NewImageFromImage(img2)
	img.FillMode = canvas.ImageFillOriginal
	formatted := container.New(layout.NewMaxLayout(), img)
	clock.Add(formatted)

}

func main() {
	webcam, _ := gocv.OpenVideoCapture(0)
	a := app.New()
	w := a.NewWindow("Background Image")
	w.Resize(fyne.NewSize(300, 300))
	// webcam, _ := gocv.OpenVideoCapture(0)
	// imgen := gocv.NewMat()
	// webcam.Read(&imgen)
	// img2, _ := imgen.ToImage()
	// img := canvas.NewImageFromImage(img2)
	// img.FillMode = canvas.ImageFillOriginal
	clock := container.New(layout.NewMaxLayout())
	updateTime(clock, webcam)
	w.SetContent(clock)

	// var fpsint int = int(fps)
	// var fpsdur time.Duration = time.Duration(fpsint)

	go func() {
		for range time.Tick(138 * time.Millisecond) {
			updateTime(clock, webcam)
			clock.Refresh()
			// clock.Remove(img)

		}
	}()
	w.CenterOnScreen()
	w.Show()
	w.ShowAndRun()

}
