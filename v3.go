package s3dm

import "math"
import "fmt"

type V3 struct {
	X, Y, Z float64
}

func NewV3(x, y, z float64) *V3 {
	v := new(V3)
	v.X = x
	v.Y = y
	v.Z = z
	return v
}

// Returns a new vector equal to 'v'
func (v *V3) Copy() *V3 {
	return NewV3(v.X, v.Y, v.Z)
}

// Returns true if 'v' and 'o' are equal
func (v *V3) Equals(o *V3) bool {
	return v.X == o.X && v.Y == o.Y && v.Z == o.Z
}

// Returns the magnitude of 'v'
func (v *V3) Length() float64 {
	return math.Sqrt((v.X * v.X) + (v.Y * v.Y) + (v.Z * v.Z))
}

// Returns the Dot product of 'v' and 'o'
func (v *V3) Dot(o *V3) float64 {
	return (v.X * o.X) + (v.Y * o.Y) + (v.Z * o.Z)
}

// Returns the cross product of 'v' and 'o'
func (v *V3) Cross(o *V3) *V3 {
	return NewV3(
		(v.Y * o.Z) - (v.Z * o.Y),
		(v.Z * o.X) - (v.X * o.Z),
		(v.X * o.Y) - (v.Y * o.X))
}

// Returns a vector reflected by 'norm'
func (v *V3) Reflect(norm *V3) *V3 {
	distance := float64(2) * v.Dot(norm)
	return NewV3(v.X - distance * norm.X, 
		v.Y - distance * norm.Y, 
		v.Z - distance * norm.Z)
}

// Returns a normalized vector perpendicular to 'v'
func (v *V3) Perp() *V3 {
	perp := v.Cross(NewV3(-1, 0, 0))
	if perp.Length() == 0 {
		// If v is too close to -x try -y
		perp = v.Cross(NewV3(0, -1, 0))
	}
	return perp.Unit()
}

// Returns a new vector equal to 'v' but with a magnitude of 'l'
func (v * V3) SetLength(l float64) *V3 {
	u := v.Unit()
	return u.Muls(l)
}

// Returns a new vector equal to 'v' normalized
func (v *V3) Unit() *V3 {
	l := v.Length()	
	return v.Divs(l)
}

func (v *V3) Add(o *V3) *V3 {
	return NewV3(
		v.X + o.X,
		v.Y + o.Y,
		v.Z + o.Z)
}

func (v *V3) Adds(o float64) *V3 {
	return NewV3(
		v.X + o,
		v.Y + o,
		v.Z + o)
}

func (v *V3) Sub(o *V3) *V3 {
	return NewV3(
		v.X - o.X,
		v.Y - o.Y,
		v.Z - o.Z)
}

func (v *V3) Subs(o float64) *V3 {
	return NewV3(
		v.X - o,
		v.Y - o,
		v.Z - o)
}

func (v *V3) Mul(o *V3) *V3 {
	return NewV3(
		v.X * o.X,
		v.Y * o.Y,
		v.Z * o.Z)
}

func (v *V3) Muls(o float64) *V3 {
	return NewV3(
		v.X * o,
		v.Y * o,
		v.Z * o)
}

func (v *V3) Div(o *V3) *V3 {
	return NewV3(
		v.X / o.X,
		v.Y / o.Y,
		v.Z / o.Z)
}

func (v *V3) Divs(o float64) *V3 {
	return NewV3(
		v.X / o,
		v.Y / o,
		v.Z / o)
}

func (v *V3) String() string {
	return fmt.Sprint(v.X, ", ", v.Y, ", ", v.Z)
}

