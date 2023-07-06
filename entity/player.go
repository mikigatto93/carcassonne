package entity

const (
	MaxNumOfFollowers = 6
)

type Player struct {
	Name           string
	FollowersCount int
	Id             string
}

func NewPlayer(name string, id string) *Player {
	p := Player{name, MaxNumOfFollowers, id}
	return &p
}
