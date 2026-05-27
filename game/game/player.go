package game

type PlayerSystem struct {
	healthPoints int
	gamePoints   int
}

type Player interface {
	IncrementGamePoints(points int)
	LoseHealthPoints(points int)
	GetHealthPoints() int
	GetGamePoints() int
}

func (player *PlayerSystem) IncrementGamePoints(points int) {
	player.gamePoints += points
}

func (player *PlayerSystem) LoseHealthPoints(points int) {
	player.healthPoints -= points
}

func (player *PlayerSystem) GetHealthPoints() int {
	return player.healthPoints
}

func (player *PlayerSystem) GetGamePoints() int {
	return player.gamePoints
}

func NewPlayer(hp, startingPoints int) Player {
	return &PlayerSystem{healthPoints: hp, gamePoints: startingPoints}
}
