package s3dm

import "math"

type Sphere struct {
	Xform
	Radius float64
}

func NewSphere(pos *V3, radius float64) *Sphere {
	s := new(Sphere)
	s.ResetXform()
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
	ro, rd := r.O(), r.D()
    A := rd.Dot(rd)
    B := float64(2) * (
		rd.X * (ro.X - pos.X) +
		rd.Y * (ro.Y - pos.Y) +
		rd.Z * (ro.Z - pos.Z))
	C := ((ro.X - pos.X) * (ro.X - pos.X) +
		(ro.Y - pos.Y) * (ro.Y - pos.Y) +
		(ro.Z - pos.Z) * (ro.Z - pos.Z)) - 
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

		intersection := NewV3(ro.X+rd.X*t, ro.Y+rd.Y*t, ro.Z+rd.Z*t)		
		normal := s.Normal(intersection)
		return intersection, normal
	}
	return nil, nil
}
