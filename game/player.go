package game

import "github.com/google/uuid"

const (
	MaxNumOfFollowers = 6
)

type Player struct {
	Name           string
	FollowersCount int
	id             string
}

func NewPlayer(name string) *Player {
	p := Player{name, MaxNumOfFollowers, uuid.NewString()}
	return &p
}
