package main

import (
	"fmt"
	"image/color"
	"strconv"
	"time"

	"github.com/Martim03/go-game/game/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	screenWidth        = 650
	screenHeight       = 450
	windowTitle        = "Playtest"
	spawnTimeSeconds   = 1 * time.Second
	startingHP         = 99
	startingGamePoints = 0
	baseDamage         = 1
	pointsPerBall      = 100
)

type Game struct {
	// TODO: Implement loaders (such as sprites, music, etc...)
	// TODO: Extend to multiplayer
	balls       map[ebiten.Key]*game.BallActor
	pressedKeys []ebiten.Key
	player      game.Player
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

/***************
* Update Logic *
***************/

func isPlayerDead(player game.Player) bool {
	return player.GetHealthPoints() <= 0
}

func spawnBall(g *Game) {
	for {
		time.Sleep(spawnTimeSeconds)
		ball := game.SpawnRandomBall(screenWidth, screenHeight)
		_, exists := g.balls[ball.GetKey()]
		if !exists {
			// Only spawn ball if the letter is unique
			g.balls[ball.GetKey()] = ball
		}
	}
}

func (g *Game) killBall(b *game.BallActor) {
	b.Destroy()
	delete(g.balls, b.GetKey())
}

func (g *Game) readInput() {
	// TODO: Should the read be buffered?
	g.pressedKeys = inpututil.AppendJustPressedKeys(g.pressedKeys[:0])

	for _, key := range g.pressedKeys {
		ball, exists := g.balls[key]
		if exists {
			g.killBall(ball)
			g.player.IncrementGamePoints(pointsPerBall)
		} else {
			g.player.LoseHealthPoints(baseDamage)
		}
	}
}

func isOutOfBounds(ball *game.BallActor) bool {
	x, y := ball.GetPos()
	r := ball.GetRadius()
	return x-r > screenWidth || x+r < 0 || y-r > screenHeight || y+r < 0
}

func (g *Game) moveBalls() {
	for _, ball := range g.balls {
		ball.Move()
		if isOutOfBounds(ball) {
			g.killBall(ball)
			g.player.LoseHealthPoints(baseDamage)
		}
	}
}

func (g *Game) Update() error {
	if isPlayerDead(g.player) {
		fmt.Println("Game Over!")
		return ebiten.Termination
	}

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

func drawHUD(screen *ebiten.Image, player game.Player) {
	// TODO: enhance this
	ebitenutil.DebugPrintAt(screen, strconv.Itoa(player.GetHealthPoints()), int(10), int(10))
	ebitenutil.DebugPrintAt(screen, strconv.Itoa(player.GetGamePoints()), int(50), int(10))
}

func (g *Game) Draw(screen *ebiten.Image) {
	drawHUD(screen, g.player)

	for _, ball := range g.balls {
		drawBall(screen, ball)
	}
}

/*************
* Start Game *
*************/

func NewGame() *Game {
	ballMap := make(map[ebiten.Key]*game.BallActor)
	pressedKeys := make([]ebiten.Key, 0)
	player := game.NewPlayer(startingHP, startingGamePoints)

	game := &Game{
		balls:       ballMap,
		pressedKeys: pressedKeys,
		player:      player,
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
