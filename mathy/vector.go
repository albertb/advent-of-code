package mathy

type Vec struct {
	X, Y int
}

func NewVec(x, y int) Vec {
	return Vec{x, y}
}

type Rect struct {
	Vec
	Width, Height int
}

func (v Vec) Plus(other Vec) Vec {
	return Vec{
		v.X + other.X,
		v.Y + other.Y,
	}
}

func (v Vec) Equals(other Vec) bool {
	return v.X == other.X && v.Y == other.Y
}

// Returns [v] rotated 90 degress clockwise [n] times.
func (v Vec) Rotate90(n int) Vec {
	if n < 0 {
		n = (n%4 + 4) % 4
	}

	for range n {
		v = Vec{
			-v.Y,
			v.X,
		}
	}

	return v
}

func (v Vec) Dot(other Vec) int {
	return v.X*other.X + v.Y*other.Y
}

func (v Vec) Cross(other Vec) int {
	return v.X*other.Y - v.Y*other.X
}

func (v *Vec) Add(other Vec) {
	tmp := v.Plus(other)
	v.X, v.Y = tmp.X, tmp.Y
}

func (r Rect) Translate(v Vec) Rect {
	return Rect{
		Vec{
			r.X + v.X,
			r.Y + v.Y,
		},
		r.Width,
		r.Height}
}

func (r Rect) Intersects(other Rect) bool {
	return r.X <= other.X+other.Width &&
		other.X <= r.X+r.Width &&
		r.Y <= other.Y+other.Height &&
		other.Y <= r.Y+r.Height
}
