package game

import (
	"math"
	"math/rand/v2"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	MinRadius = 25
	MaxRadius = 50
)

// TODO: Create spawner struct

func computeTrajecory(x, y, targetX, targetY float32) (float32, float32) {
	// TODO: migrate this to physics module
	// map to origin
	xRel := targetX - x
	yRel := targetY - y

	sumSquares := float64(xRel*xRel + yRel*yRel)
	R := math.Sqrt(sumSquares)

	if R == 0 {
		// Prevent division by zero
		return 0, 0
	}

	return xRel / float32(R), yRel / float32(R)
}

func getRandomLetterKey() ebiten.Key {
	// KeyA to KeyZ are contiguous integers
	return ebiten.Key(rand.IntN(int(ebiten.KeyZ-ebiten.KeyA+1)) + int(ebiten.KeyA))
}

func SpawnRandomBall(maxWidth, maxHeight float32) *BallActor {
	// Generate random radius
	radius := rand.IntN(51-25) + 25

	// Establish spawn limits
	xMin := 0 - (float32(radius) - 10)
	yMin := xMin
	xMax := maxWidth + (float32(radius) - 10)
	yMax := maxHeight + (float32(radius) - 10)

	// random float: rand.Float32() * (max - min) + min

	// Generate random spawn point
	var xPoint float32
	var yPoint float32

	origin := rand.IntN(4) // left, up, right, down
	switch origin {
	case 0:
		xPoint = xMin
		yPoint = rand.Float32()*(yMax-yMin) + yMin
	case 1:
		xPoint = rand.Float32()*(xMax-xMin) + xMin
		yPoint = yMin
	case 2:
		xPoint = xMax
		yPoint = rand.Float32()*(yMax-yMin) + yMin
	case 3:
		xPoint = rand.Float32()*(xMax-xMin) + xMin
		yPoint = yMax
	}

	// Generate trajectory
	// TODO: randomize point, dont use just the middle
	dtX, dtY := computeTrajecory(xPoint, yPoint, maxWidth/2, maxHeight/2)

	// Infer velocity
	dtX *= 3
	dtY *= 3

	// Generate random key
	key := getRandomLetterKey()

	ball := NewBall(xPoint, yPoint, float32(radius), key)
	ball.Accelerate(dtX, dtY)
	return ball
}
