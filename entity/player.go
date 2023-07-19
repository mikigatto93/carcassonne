package entity

const (
	MaxNumOfFollowers = 6
)

type Player struct {
	Name           string
	FollowersCount int
	Id             string
	Order          uint8
	Ready          bool
}

func NewPlayer(name string, id string, order uint8) *Player {
	p := Player{name, MaxNumOfFollowers, id, order, false}
	return &p
}
