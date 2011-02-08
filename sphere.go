package s3dm

import "math"

type Sphere struct {
	Transform
	Radius float64
}

func NewSphere(pos *V3, radius float64) *Sphere {
	s := new(Sphere)
	s.SetIdentity()
	s.SetPosition(pos)
	s.Radius = radius
	return s
}

// Returns the normal vector for a point 'p' on sphere 's'
func (s *Sphere) Normal(p *V3) *V3 {
	delta := p.Sub(s.Position())
	return delta.Unit()
}

func (s *Sphere) Intersect(r *Ray) (*V3, *V3) {	
	pos := s.Position()
    A := r.D.Dot(r.D)
    B := float64(2) * (
		r.D.X * (r.O.X - pos.X) +
		r.D.Y * (r.O.Y - pos.Y) +
		r.D.Z * (r.O.Z - pos.Z))
	C := ((r.O.X - pos.X) * (r.O.X - pos.X) +
		(r.O.Y - pos.Y) * (r.O.Y - pos.Y) +
		(r.O.Z - pos.Z) * (r.O.Z - pos.Z)) - 
		s.Radius * s.Radius

	delta := B * B - 4 * A * C;
	if delta > 0 {
		t0 := (-B - math.Sqrt(delta)) / 2;
		t1 := (-B + math.Sqrt(delta)) / 2;

		t := float64(0)

		// t0 must be smaller than t1
		if t0 > t1 {
			t0, t1 = t1, t0
		}

		// Sphere behind ray
		if t1 < 0 {
			return nil, nil
		}
	
		if t0 < 0 {
			t = t1
		} else {
			t = t0
		}

		intersection := NewV3(r.O.X+r.D.X*t, r.O.Y+r.D.Y*t, r.O.Z+r.D.Z*t)		
		normal := s.Normal(intersection)
		return intersection, normal
	}
	return nil, nil
}
