package point

type Point struct {
	x float64
	y float64
}

func (p *Point) NewPoint(x, y float64) {
	p.x = x
	p.y = y
}

func (p *Point) X() float64 {
	return p.x
}

func (p *Point) Y() float64 {
	return p.y
}
