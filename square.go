package square

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func NewSquare(x, y int, a uint) Square {
	return Square{Point{x, y}, a}
}

func (s Square) End() Point {
	return Point{
		x: s.start.x + int(s.a),
		y: s.start.y + int(s.a),
	}
}

func (s Square) Area() uint {
	return s.a * s.a
}

func (s Square) Perimeter() uint {
	return s.a * 4
}
