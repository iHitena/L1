package TP24

type Point struct {
	Distance int
	x        int
	y        int
}

// попытка реализовать конструктор
func NewPoint(x, y int) *Point {

	return &Point{x: x, y: y}
}

func (p *Point) DistanceXY() int {
	p.Distance = p.y - p.x
	return p.Distance
}
