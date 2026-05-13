package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

type GameObject struct {
	xPos, yPos float32
	xVel, yVel float32
}

type Entity interface {
	GetPos() (float32, float32)
	Move(x, y float32)
	Accelerate(dtX, dtY float32)
	Destroy()
}

func (self *GameObject) GetPos() (float32, float32) {
	return self.xPos, self.yPos
}

func (self *GameObject) Move(x, y float32) {
	// TODO: add delta time
	self.xPos += self.xVel
	self.yPos += self.yVel
}

func (self *GameObject) Accelerate(dtX, dtY float32) {
	self.xVel += dtX
	self.yVel += dtY
}

func (self *GameObject) Destroy() {
	fmt.Println("Destroyed", self)
}

type BallActor struct {
	GameObject
	radius float32
	key    ebiten.Key
}

type Ball interface {
	GetRadius() float32
	GetKey() ebiten.Key
	VerifyKey(k ebiten.Key) bool
}

func (self *BallActor) GetRadius() float32 {
	return self.radius
}

func (self *BallActor) GetKey() ebiten.Key {
	return self.key
}

func (self *BallActor) VerifyKey(k ebiten.Key) bool {
	return k == self.key
}

// TODO: generate random key
// TODO: spawn at random position
// TODO: generate random radius
func NewBall() BallActor {
	return BallActor{
		GameObject: GameObject{
			xPos: 100,
			yPos: 100,
			xVel: 0,
			yVel: 0,
		},
		radius: 50,
		key:    ebiten.KeyW,
	}
}
