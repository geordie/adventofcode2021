package navigation

import (
	"log"
	"strconv"
	"strings"
)

type Direction int

const (
	None Direction = iota
	Up
	Down
	Forward
)

type Vector struct {
	Direction Direction
	Magnitude int
}

type Position struct {
	Aim     int
	Forward Vector
	Down    Vector
}

func ParseDirection(s string) Vector {

	magnitude, err := strconv.Atoi(strings.Split(s, " ")[1])
	if err != nil {
		log.Fatal(err)
	}
	direction := None
	if strings.HasPrefix(s, "f") {
		direction = Forward
	} else if strings.HasPrefix(s, "u") {
		direction = Up
	} else if strings.HasPrefix(s, "d") {
		direction = Down
	}

	return Vector{direction, magnitude}
}

func (pos *Position) Add(v Vector) {
	if v.Direction == Forward {
		pos.Forward.Magnitude += v.Magnitude
	} else if v.Direction == Down {
		pos.Down.Magnitude += v.Magnitude
	} else if v.Direction == Up {
		pos.Down.Magnitude -= v.Magnitude
	}
}

func (pos *Position) Move(v Vector) {
	if v.Direction == Forward {
		pos.Forward.Magnitude += v.Magnitude
		pos.Down.Magnitude += pos.Aim * v.Magnitude
	} else if v.Direction == Down {
		pos.Aim += v.Magnitude
	} else if v.Direction == Up {
		pos.Aim -= v.Magnitude
	}
}

func (twoD *Position) Product() int {
	return twoD.Forward.Magnitude * twoD.Down.Magnitude
}
