package s3dm

import "math"
import "fmt"

type V3 struct {
	X, Y, Z float64
}

// Returns true if 'v' and 'o' are equal
func (v V3) Equals(o V3) bool {
	return v.X == o.X && v.Y == o.Y && v.Z == o.Z
}

func (v V3) LengthSquared() float64 {
	return v.X*v.X + v.Y*v.Y + v.Z*v.Z
}

// Returns the magnitude of 'v'
func (v V3) Length() float64 {
	return math.Sqrt(v.LengthSquared())
}

// Returns the distance from 'v' to 'o'
func (v V3) Distance(o V3) float64 {
	return v.Sub(o).Length()
}

// Returns the Dot product of 'v' and 'o'
func (v V3) Dot(o V3) float64 {
	return v.X*o.X + v.Y*o.Y + v.Z*o.Z
}

// Returns the cross product of 'v' and 'o'
func (v V3) Cross(o V3) V3 {
	return V3{
		v.Y*o.Z - v.Z*o.Y,
		v.Z*o.X - v.X*o.Z,
		v.X*o.Y - v.Y*o.X}
}

// Returns a vector reflected by 'norm'
func (v V3) Reflect(norm V3) V3 {
	distance := float64(2) * v.Dot(norm)
	return V3{v.X - distance*norm.X,
		v.Y - distance*norm.Y,
		v.Z - distance*norm.Z}
}

// Returns a normalized vector perpendicular to 'v'
func (v V3) Perp() V3 {
	perp := v.Cross(V3{-1, 0, 0})
	if perp.Length() == 0 {
		// If v is too close to -x try -y
		perp = v.Cross(V3{0, -1, 0})
	}
	return perp.Unit()
}

// Returns a new vector equal to 'v' but with a magnitude of 'l'
func (v V3) SetLength(l float64) V3 {
	return v.Unit().Muls(l)
}

// Returns a new vector equal to 'v' normalized
func (v V3) Unit() V3 {
	return v.Divs(v.Length())
}

func (v V3) Add(o V3) V3 {
	return V3{
		v.X + o.X,
		v.Y + o.Y,
		v.Z + o.Z}
}

func (v *V3) AddLocal(o V3) {
	v.X += o.X
	v.Y += o.Y
	v.Z += o.Z
}

func (v V3) Adds(o float64) V3 {
	return v.Add(V3{o, o, o})
}

func (v *V3) AddsLocal(o float64) {
	v.AddLocal(V3{o, o, o})
}

func (v V3) Sub(o V3) V3 {
	return V3{
		v.X - o.X,
		v.Y - o.Y,
		v.Z - o.Z}
}

func (v *V3) SubLocal(o V3) {
	v.X -= o.X
	v.Y -= o.Y
	v.Z -= o.Z
}

func (v V3) Subs(o float64) V3 {
	return v.Sub(V3{o, o, o})
}

func (v *V3) SubsLocal(o float64) {
	v.SubLocal(V3{o, o, o})
}

func (v V3) Mul(o V3) V3 {
	return V3{
		v.X * o.X,
		v.Y * o.Y,
		v.Z * o.Z}
}

func (v *V3) MulLocal(o V3) {
	v.X *= o.X
	v.Y *= o.Y
	v.Z *= o.Z
}

func (v V3) Muls(o float64) V3 {
	return v.Mul(V3{o, o, o})
}

func (v *V3) MulsLocal(o float64) {
	v.MulLocal(V3{o, o, o})
}

func (v V3) Div(o V3) V3 {
	return V3{
		v.X / o.X,
		v.Y / o.Y,
		v.Z / o.Z}
}

func (v *V3) DivLocal(o V3) {
	v.X /= o.X
	v.Y /= o.Y
	v.Z /= o.Z
}

func (v V3) Divs(o float64) V3 {
	return v.Div(V3{o, o, o})
}

func (v *V3) DivsLocal(o float64) {
	v.DivLocal(V3{o, o, o})
}

func (v V3) String() string {
	return fmt.Sprint(v.X, ", ", v.Y, ", ", v.Z)
}
