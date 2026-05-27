package main

import (
	"image/color"
	"time"

	"github.com/Martim03/go-game/game/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth      = 650
	screenHeight     = 450
	windowTitle      = "Playtest"
	spawnTimeSeconds = 1 * time.Second
)

type Game struct {
	// TODO: Implement loaders (such as sprites, music, etc...)
	balls       map[ebiten.Key]*game.BallActor
	pressedKeys []ebiten.Key
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

/***************
* Update Logic *
***************/

func spawnBall(g *Game) {
	for {
		time.Sleep(spawnTimeSeconds)
		ball := game.SpawnRandomBall(screenWidth, screenHeight)
		g.balls[ball.GetKey()] = ball
	}
}

func (g *Game) killBall(b *game.BallActor) {
	b.Destroy()
	delete(g.balls, b.GetKey())
}

func (g *Game) readInput() {
	// TODO: Should the read be buffered?
	g.pressedKeys = inpututil.AppendPressedKeys(g.pressedKeys[:0])

	for _, key := range g.pressedKeys {
		ball, exists := g.balls[key]
		if exists {
			g.killBall(ball)
			// TODO: reward player
		}
		// TODO: else take damage?
	}
}

func isOutOfBounds(ball *game.BallActor) bool {
	x, y := ball.GetPos()
	r := ball.GetRadius()
	return x-r > screenWidth || x+r < 0 || y-r > screenWidth || y+r < 0
}

func (g *Game) moveBalls() {
	for _, ball := range g.balls {
		ball.Move()
		if isOutOfBounds(ball) {
			g.killBall(ball)
			// TODO: Take damage?
		}
	}
}

func (g *Game) Update() error {
	g.readInput()
	g.moveBalls()
	return nil
}

/****************
* Drawing Logic *
****************/

func drawBall(screen *ebiten.Image, ball *game.BallActor) {
	// TODO: generate random color
	lightNavyBlue := color.RGBA{45, 89, 135, 255}

	x, y := ball.GetPos()
	r := ball.GetRadius()
	vector.FillCircle(screen, x, y, r, lightNavyBlue, false)
	ebitenutil.DebugPrintAt(screen, ball.GetKey().String(), int(x), int(y))
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, ball := range g.balls {
		drawBall(screen, ball)
	}
}

/*************
* Start Game *
*************/

func NewGame() *Game {
	bMap := make(map[ebiten.Key]*game.BallActor)
	pk := make([]ebiten.Key, 0)

	game := &Game{
		balls:       bMap,
		pressedKeys: pk,
	}

	go spawnBall(game)

	return game
}

func main() {
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle(windowTitle)

	game := NewGame()
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
