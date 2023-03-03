package constant

import "errors"

type Direction string

const (
	North Direction = "N"
	East  Direction = "E"
	South Direction = "S"
	West  Direction = "W"
)

func ToDirection(str string) (Direction, error) {
	switch str {
	case "N":
		return North, nil
	case "E":
		return East, nil
	case "S":
		return South, nil
	case "W":
		return West, nil
	default:
		return "", errors.New("Direção não reconhecida " + str)
	}
}
