package main

import (
	"time"

	"image"
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"gocv.io/x/gocv"
)


func main() {
	webcam, _ := gocv.VideoCaptureFile("Basler_video.ogg")
	a := app.New()
	w := a.NewWindow("Background Image")
	w.Resize(fyne.NewSize(400, 400))
	// var fpsint int = int(fps)
	// var fpsdur time.Duration = time.Duration(fpsint)
	go func() {
		for range time.Tick(33 * time.Millisecond) {

			imgen := gocv.NewMat()
			matCanny := gocv.NewMat()
			matLines := gocv.NewMat()
			webcam.Read(&imgen)
			// gocv.Resize(imgen, &imgen, image.Pt(400, 200), 0, 0, gocv.InterpolationArea)
			croppedMat := imgen.Region(image.Rect(0, 0, 600, 700))
			imgen2 := croppedMat.Clone()

			gocv.Canny(imgen2, &matCanny, 230, 250)
			gocv.HoughLinesP(matCanny, &matLines, 1, math.Pi/180, 80)
			for i := 0; i < matLines.Rows(); i++ {
				pt1 := image.Pt(0, int(matLines.GetVeciAt(i, 0)[1]))
				pt2 := image.Pt(600, int(matLines.GetVeciAt(i, 0)[3]))
				gocv.Line(&imgen2, pt1, pt2, color.RGBA{0, 255, 0, 50}, 10)
			}
			img2, _ := imgen2.ToImage()
			img := canvas.NewImageFromImage(img2)

			//img.FillMode = canvas.ImageFillOriginal
			clock := container.New(layout.NewMaxLayout(), img)

			w.SetContent(clock)

		}
	}()
	w.CenterOnScreen()
	w.ShowAndRun()

}
