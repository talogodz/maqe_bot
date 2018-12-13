package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
)

type Robot struct {
	angle    float64
	position []int
}

func main() {
	params := os.Args[1]

	robot := Robot{
		angle:    90.0, // Facing north
		position: []int{0, 0},
	}
	robot.execute(params)
	robot.PrintPosition()
}

func (r *Robot) execute(command string) {

	//\d+|\D+ regex

	re := regexp.MustCompile(`\d+|\D`)
	cmds := re.FindAllString(command, -1)
	for i := 0; i < len(cmds); i++ {
		switch cmds[i] {
		case "R":
			r.turn(-90)
		case "L":
			r.turn(90)
		case "W":
			n, err := strconv.Atoi(cmds[i+1])
			i++ /// skip number param
			if err != nil {
				break
			}
			r.move(n)
		}
	}
}

func (r *Robot) turn(degree float64) {
	r.angle += degree
	if r.angle < 0 {
		r.angle = 270
	} else if r.angle > 360 {
		r.angle = 90
	} else if r.angle == 360 {
		r.angle = 0
	}
}

func (r *Robot) move(n int) {
	rad := r.angle * math.Pi / 180
	sin := int(math.Round(math.Sin(rad)))
	cos := int(math.Round(math.Cos(rad)))
	r.position = []int{r.position[0] + (cos * n), r.position[1] + (sin * n)}
}

func (r *Robot) PrintPosition() {
	strDirection := toStringDirection(r.angle)
	fmt.Printf("X : %d Y: %d Direction: %s\n", r.position[0], r.position[1], strDirection)

}

func toStringDirection(angle float64) string {
	nAngle := int(angle)
	switch nAngle {
	case 0:
		return "East"
	case 90:
		return "North"
	case 180:
		return "West"
	case 270:
		return "South"
	}
	return ""
}
