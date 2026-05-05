package game

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

type Entity struct {
	xPos, yPos float32
	xVel, yVel float32
	radius     float32
	key        ebiten.Key
}

type Ball interface {
	GetPos() (float32, float32)
	GetRadius() float32
	GetKey() ebiten.Key
	Move(x, y float32)
	Accelerate(dtX, dtY float32)
	VerifyKey(k ebiten.Key) bool
	Destroy()
}

func (self *Entity) GetPos() (float32, float32) {
	return self.xPos, self.yPos
}

func (self *Entity) GetRadius() float32 {
	return self.radius
}

func (self *Entity) GetKey() ebiten.Key {
	return self.key
}

func (self *Entity) Move(x, y float32) {
	// TODO: add delta time
	self.xPos += self.xVel
	self.yPos += self.yVel
}

func (self *Entity) Accelerate(dtX, dtY float32) {
	self.xVel += dtX
	self.yVel += dtY
}

func (self *Entity) Destroy() {
	fmt.Println("Destroyed :(")
	self = nil
}

func (self *Entity) VerifyKey(k ebiten.Key) bool {
	return k == self.key
}

func NewBall() Ball {
	// TODO: generate random key
	// TODO: spawn at random position
	// TODO: generate random radius
	var ball Ball = &Entity{xPos: 100, yPos: 100, xVel: 0, yVel: 0, radius: 50, key: ebiten.KeyW}
	return ball
}
