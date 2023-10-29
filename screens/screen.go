package screens

import (

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
)

type GameScene struct {
	window fyne.Window
	content        *fyne.Container
}

func NewScene(window fyne.Window) *GameScene {
	scene := &GameScene{window: window}
    scene.Render()
    return scene
}

func (s *GameScene) Render() {
	backgroundImage := canvas.NewImageFromURI( storage.NewFileURI("./assets/bg.png") )
    backgroundImage.Resize(fyne.NewSize(800,600))
	backgroundImage.Move( fyne.NewPos(0,0) )

	s.content = container.NewWithoutLayout(
        backgroundImage, // Fondo
    )
    s.window.SetContent(s.content) 
    /* models.StartGame() */
}
