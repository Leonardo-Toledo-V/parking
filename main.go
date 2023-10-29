package main

import (
    "parking/screens"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Parking")

	w.CenterOnScreen()
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(800, 600))
	screens.NewScene(w)
	w.ShowAndRun()
}
