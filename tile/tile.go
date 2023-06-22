package tile

import (
	"errors"
	"fmt"
	"strings"
)

type Tile struct {
	x      int
	y      int
	id     int
	layout TileLayout
}

func New(x int, y int, id int, layout TileLayout) *Tile {
	t := Tile{
		x:      x,
		y:      y,
		id:     id,
		layout: layout,
	}

	return &t
}

func (t Tile) String() string {
	var sb strings.Builder = strings.Builder{}

	for i := 0; i < len(t.layout); i++ {
		for j := 0; j < len(t.layout); j++ {
			sb.WriteString(fmt.Sprintf("%d ", t.layout[i][j]))
		}
		sb.WriteString("\n")
	}

	return sb.String()
}

func (t *Tile) Rotate(n int) {
	rot := n % 4

	for i := 0; i < rot; i++ {
		t.layout = rotateBy90DegRight(t.layout)
	}

}

func (t *Tile) CanAttach(other *Tile) (bool, error) {
	side, otherSide := -1, -1
	if t.x+1 == other.x { //other tile left side
		side = 1
		otherSide = 3
	} else if t.x-1 == other.x { //other tile right side
		side = 3
		otherSide = 1
	} else if t.y+1 == other.y { //other tile down side
		side = 0
		otherSide = 2
	} else if t.y-1 == other.y { //other tile up side
		side = 2
		otherSide = 0
	}

	// TODO: handle errors
	sideList, _ := t.getSide(side)
	otherSideList, _ := other.getSide(otherSide)

	//fmt.Println(sideList)
	//fmt.Println(otherSideList)

	ok := true
	for i := 0; ok && i < len(sideList); i++ {
		if sideList[i] != otherSideList[i] {
			ok = false
		}
	}

	// TODO: handle errors
	return ok, nil

}

func rotateBy90DegRight(matrix TileLayout) TileLayout {

	// reverse the matrix
	for i, j := 0, len(matrix)-1; i < j; i, j = i+1, j-1 {
		matrix[i], matrix[j] = matrix[j], matrix[i]
	}

	// transpose it
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < i; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	return matrix
}

func (t *Tile) getSide(side int) ([]Terrain, error) {
	maxLen := len(t.layout)
	arr := make([]Terrain, maxLen)

	if side == 0 { // up
		for i := 0; i < maxLen; i++ {
			arr[i] = t.layout[0][i]
		}
		return arr, nil

	} else if side == 2 { // down
		for i := 0; i < maxLen; i++ {
			arr[i] = t.layout[maxLen-1][i]
		}
		return arr, nil

	} else if side == 1 { // right
		for i := 0; i < maxLen; i++ {
			arr[i] = t.layout[i][maxLen-1]
		}
		return arr, nil

	} else if side == 3 { // left
		for i := 0; i < maxLen; i++ {
			arr[i] = t.layout[i][0]
		}
		return arr, nil

	} else {
		return arr, errors.New("InvalidArgument: side must be between 0 and 3")
	}
}
