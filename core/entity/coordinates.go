package entity

type Coordinates struct {
	AxisY int
	AxisX int
}

func NewCoordinates(axisY int, axisX int) Coordinates {
	return Coordinates{AxisX: axisX, AxisY: axisY}
}
