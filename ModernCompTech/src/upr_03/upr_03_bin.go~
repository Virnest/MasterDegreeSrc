package main

import (
	"image/color"
	"math/rand"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	//"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/canvas"
)

const (
	CLIENT_HEIGHT, CLIENT_WIDTH float32 = 1080, 1920
	MASK_HEIGHT, MASK_WIDTH float32 = 300, 200
)

func main() {	
	App := app.New()
	Window := App.NewWindow("Upr 03")
	
	Window.Resize(fyne.NewSize(CLIENT_WIDTH, CLIENT_HEIGHT))
	//Window.SetFullScreen(true)
	myContainer := fyne.NewContainer()
	
	randX := rand.Perm(int(CLIENT_WIDTH - MASK_WIDTH))
	randY := rand.Perm(int(CLIENT_HEIGHT - MASK_HEIGHT))

	for i := range 10 {
		x := float32(randX[i])
		y := float32(randY[i])
		draw_mask(myContainer, x, y)
		//draw_mask(myContainer, 300, 50)
	}
	Window.SetContent(myContainer)
	Window.Show()

	App.Run()
}
func draw_mask(myContainer *fyne.Container, x float32, y float32) {
	clrs := []color.NRGBA{
	 	color.NRGBA{R: 255, G: 0, B: 0, A: 255},
		color.NRGBA{R: 0, G: 255, B: 0, A: 255},
		color.NRGBA{R: 0, G: 0, B: 255, A: 255},
	}
	
	smth := rand.Perm(3)
	i1, i2, i3 := smth[0], smth[1], smth[2]
	RED, GREEN, BLUE := clrs[i1], clrs[i2], clrs[i3]

	draw_face(myContainer, x, y, BLUE)
	draw_eye(myContainer, x+50, y+50, RED)
	draw_eye(myContainer, x+125, y+50, RED)
	draw_nose(myContainer, x+95, y+50, GREEN)
	draw_mouth(myContainer, x+50, y+225, GREEN)
}

func draw_face(content *fyne.Container, x float32, y float32, clr color.NRGBA) {
	face := canvas.NewCircle(clr)
	face.Move(fyne.Position{X: x, Y: y})
	face.Resize(fyne.Size{Width: MASK_WIDTH, Height: MASK_HEIGHT})
	face.StrokeColor = color.Black
	face.StrokeWidth = 1.5
	
	content.Add(face)
}

func draw_eye(content *fyne.Container, x float32, y float32, clr color.NRGBA) {
	eye := canvas.NewCircle(clr)
	eye.Move(fyne.Position{X: x, Y: y})
	eye.Resize(fyne.Size{Width: 25, Height: 100})
	eye.StrokeColor = color.Black
	eye.StrokeWidth = 1.5
	pupil := canvas.NewCircle(color.Black)
	pupil.Move(fyne.Position{X: x+5, Y: y+50})
	pupil.Resize(fyne.Size{Width: 15, Height: 25})

	content.Add(eye)
	content.Add(pupil)
} 

func draw_nose(content *fyne.Container, x float32, y float32, clr color.NRGBA) {	
	nose := canvas.NewRectangle(clr)
	nose.Move(fyne.Position{X: x, Y: y})	
	nose.Resize(fyne.Size{Width:10, Height: 165})
	nose.StrokeColor = color.Black
	nose.StrokeWidth = 1.5

	content.Add(nose)
}

func draw_mouth(content *fyne.Container, x float32, y float32, clr color.NRGBA) {
	mouth := canvas.NewRectangle(clr)
	mouth.Move(fyne.Position{X: x, Y: y})	
	mouth.Resize(fyne.Size{Width:100, Height: 10})
	mouth.StrokeColor = color.Black
	mouth.StrokeWidth = 1.5

	content.Add(mouth)
}
