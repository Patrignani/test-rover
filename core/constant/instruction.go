package constant

import "errors"

type Instruction string

const (
	Left  Instruction = "L"
	Right Instruction = "R"
	Move  Instruction = "M"
)

func ToInstruction(str string) (Instruction, error) {
	switch str {
	case "L":
		return Left, nil
	case "R":
		return Right, nil
	case "M":
		return Move, nil
	default:
		return "", errors.New("Instrução não reconhecida " + str)
	}
}
