package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"gocv.io/x/gocv"
	"image"
	"image/color"
	"math"
	"time"
)

func main() {
	webcam, _ := gocv.VideoCaptureFile("Basler_video.ogg")
	a := app.New()
	w := a.NewWindow("Pilot Line v1.0")
	w.Resize(fyne.NewSize(800, 400))
	btn_start := widget.NewButton("Start", func() {})
	btn_stop := widget.NewButton("Stop", a.Quit)
	input := widget.NewEntry()
	input.SetPlaceHolder("Enter text...")
	imgen := gocv.NewMat()
	matCanny := gocv.NewMat()
	matLines := gocv.NewMat()
	var imgen1, imgen2 = line_det2(webcam, imgen, matCanny, matLines)
	var content1 = pic(imgen1, imgen2)
	fmt.Println("3")
	go func() {
		for range time.Tick(33 * time.Millisecond) {
			w.Resize(fyne.NewSize(800, 400))
			fmt.Println("4")
			imgen1, imgen2 = line_det2(webcam, imgen, matCanny, matLines)
			content1 = pic(imgen1, imgen2)
			btn_box := container.New(layout.NewGridLayout(2), btn_stop, btn_start)
			content2 := container.New(layout.NewGridWrapLayout(fyne.NewSize(100, 50)), btn_box, input)
			content := container.New(layout.NewGridLayoutWithRows(2), content1, content2)
			w.SetContent(content)
		}
	}()
	w.CenterOnScreen()
	w.ShowAndRun()
}

func line_det2(wc *gocv.VideoCapture, imgen gocv.Mat, matCanny gocv.Mat, matLines gocv.Mat) (gocv.Mat, gocv.Mat) {
	//defer close(ch1)
	//defer close(ch2)

	wc.Read(&imgen)
	// gocv.Resize(imgen, &imgen, image.Pt(400, 200), 0, 0, gocv.InterpolationArea)
	croppedMat := imgen.Region(image.Rect(0, 0, 600, 700))
	imgen2 := croppedMat.Clone()
	imgen1 := imgen.Clone()

	gocv.Canny(imgen2, &matCanny, 230, 250)
	gocv.HoughLinesP(matCanny, &matLines, 1, math.Pi/180, 80)
	for i := 0; i < matLines.Rows(); i++ {
		pt1 := image.Pt(0, int(matLines.GetVeciAt(i, 0)[1]))
		pt2 := image.Pt(600, int(matLines.GetVeciAt(i, 0)[3]))
		pt3 := image.Pt(0, 640)
		pt4 := image.Pt(600, 640)
		pt5 := image.Pt(0, 390)
		pt6 := image.Pt(600, 390)

		gocv.Line(&imgen2, pt1, pt2, color.RGBA{0, 255, 0, 50}, 10)
		gocv.Line(&imgen2, pt3, pt4, color.RGBA{0, 255, 0, 50}, 10)
		gocv.Line(&imgen2, pt5, pt6, color.RGBA{255, 0, 0, 50}, 5)
	}
	fmt.Println("1")

	return imgen1, imgen2

}
func pic(val1 gocv.Mat, val2 gocv.Mat) *fyne.Container {
	img2, _ := val1.ToImage()
	img3, _ := val2.ToImage()
	img := canvas.NewImageFromImage(img2)
	imgnew := canvas.NewImageFromImage(img3)
	img_box := container.New(layout.NewGridLayout(2), img, imgnew)
	content1 := container.NewMax(img_box)
	fmt.Println("2")
	return content1
}
