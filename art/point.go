package art

import "math"

type Point struct {
	x, y float64
}

func (p *Point) Scale(factor float64) {
	p.x *= factor
	p.y *= factor
}
func (p *Point) Translate(x, y float64) {
	p.x += x
	p.y += y
}
func (p *Point) Rotate(radians float64) {
	s := math.Sin(radians)
	c := math.Cos(radians)
	px := p.x
	py := p.y
	p.x = px*c - py*s
	p.y = px*s + py*c
}
