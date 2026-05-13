package main

import (
	"image/color"

	"github.com/Martim03/go-game/game/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth  = 650
	screenHeight = 450
	windowTitle  = "Playtest"
)

type Game struct {
	// TODO: Implement loaders (such as sprites, music, etc...)
	balls       []game.BallActor
	pressedKeys []ebiten.Key
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

/***************
* Update Logic *
***************/

func (g *Game) killBall(i int, b game.BallActor) {
	g.balls = append(g.balls[:i], g.balls[i+1:]...)
	b.Destroy()
}

func (g *Game) readInput() {
	// TODO: Should the read be buffered?
	g.pressedKeys = inpututil.AppendPressedKeys(g.pressedKeys[:0])

	// TODO: Optimize this with MAP
	for _, k := range g.pressedKeys {
		for i, b := range g.balls {
			if b.VerifyKey(k) {
				g.killBall(i, b)
			}
		}
	}
}

func (g *Game) Update() error {
	g.readInput()
	return nil
}

/****************
* Drawing Logic *
****************/

func drawBall(screen *ebiten.Image, b game.BallActor) {
	// TODO: generate random color
	lightNavyBlue := color.RGBA{45, 89, 135, 255}

	x, y := b.GetPos()
	r := b.GetRadius()
	vector.FillCircle(screen, x, y, r, lightNavyBlue, false)
	ebitenutil.DebugPrintAt(screen, b.GetKey().String(), int(x), int(y))
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, b := range g.balls {
		drawBall(screen, b)
	}
}

/*************
* Start Game *
*************/

func NewGame() *Game {
	b := make([]game.BallActor, 0)
	b = append(b, game.NewBall())
	pk := make([]ebiten.Key, 0)

	return &Game{
		balls:       b,
		pressedKeys: pk,
	}
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle(windowTitle)

	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
