package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	//"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	const ClientHeight, ClientWidth float32 = 400, 300
	
	//colors
	RED := color.NRGBA{R: 255, G: 0, B: 0, A: 255}
	GREEN := color.NRGBA{R: 0, G: 255, B: 0, A: 255}
	BLUE := color.NRGBA{R: 0, G: 0, B: 255, A: 255}

	App := app.New()
	Window := App.NewWindow("Upr 01")
	
	Window.Resize(fyne.NewSize(ClientWidth, ClientHeight))

	face := canvas.NewCircle(BLUE)
	face.Move(fyne.Position{X: 50, Y: 50})
	face.Resize(fyne.Size{Width: 200, Height: 300})
	face.StrokeColor = color.Black
	face.StrokeWidth = 1.5
	
	eye_1 := canvas.NewCircle(RED)
	eye_1.Move(fyne.Position{X: 100, Y: 100})
	eye_1.Resize(fyne.Size{Width: 25, Height: 100})
	eye_1.StrokeColor = color.Black
	eye_1.StrokeWidth = 1.5
	pupil_1 := canvas.NewCircle(color.Black)
	pupil_1.Move(fyne.Position{X: 105, Y: 150})
	pupil_1.Resize(fyne.Size{Width: 15, Height: 25})

	eye_2 := canvas.NewCircle(RED)
	eye_2.Move(fyne.Position{X: 175, Y: 100})
	eye_2.Resize(fyne.Size{Width: 25, Height: 100})
	eye_2.StrokeColor = color.Black
	eye_2.StrokeWidth = 1.5
	pupil_2 := canvas.NewCircle(color.Black)
	pupil_2.Move(fyne.Position{X: 180, Y: 150})
	pupil_2.Resize(fyne.Size{Width: 15, Height: 25})

	mouth := canvas.NewRectangle(GREEN)
	mouth.Move(fyne.Position{X: 100, Y: 275})	
	mouth.Resize(fyne.Size{Width:100, Height: 10})
	mouth.StrokeColor = color.Black
	mouth.StrokeWidth = 1.5
	
	nose := canvas.NewRectangle(GREEN)
	nose.Move(fyne.Position{X: 145, Y: 100})	
	nose.Resize(fyne.Size{Width:10, Height: 165})
	nose.StrokeColor = color.Black
	nose.StrokeWidth = 1.5
	Window.SetContent(fyne.NewContainer(face, eye_1, pupil_1, eye_2, pupil_2, nose, mouth))
	
	Window.Show()

	App.Run()
}
