package mathy

import "math"

type Vec struct {
	X, Y int
}

func NewVec(x, y int) Vec {
	return Vec{x, y}
}

func Cardinals() []Vec {
	return []Vec{{-1, 0}, {0, 1}, {0, -1}, {1, 0}}
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

func (v Vec) Minus(other Vec) Vec {
	return Vec{
		v.X - other.X,
		v.Y - other.Y,
	}
}

func (v Vec) Distance(other Vec) int {
	return Abs(v.X-other.X) + Abs(v.Y-other.Y)
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

type Bounds struct {
	Vec
}

func (b Bounds) Contains(v Vec) bool {
	return v.X >= 0 && v.X <= b.X && v.Y >= 0 && v.Y <= b.Y
}

// Grow b as needed in order to bound v.
func (b *Bounds) Bound(v Vec) {
	if v.X >= 0 && v.Y >= 0 {
		b.X = Max(b.X, v.X)
		b.Y = Max(b.Y, v.Y)
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

type Vec3 struct {
	X, Y, Z int
}

func (v Vec3) Equals(other Vec3) bool {
	return v.X == other.X && v.Y == other.Y && v.Z == other.Z
}

func (v Vec3) Distance(other Vec3) float64 {
	dX := other.X - v.X
	dY := other.Y - v.Y
	dZ := other.Z - v.Z
	return math.Sqrt(float64(dX*dX + dY*dY + dZ*dZ))
}
