package mathy

type Vec struct {
	X, Y int
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
