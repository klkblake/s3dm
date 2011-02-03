package s3dm

type Plane struct {
	O, N *V3
}

func NewPlane(origin, normal *V3) *Plane {
	p := new(Plane)
	p.O = origin
	p.N = normal
	return p
}

func (p *Plane) Intersect(r *Ray) (*V3, *V3) {
	denom := p.N.Dot(r.D)
	if denom == 0 {
		return nil, nil
	}
	c := p.N.Dot(p.O)
	t := (c - p.N.Dot(r.O)) / denom
	if t <= 0 {
		return nil, nil
	}
	return r.O.Add(r.D.Muls(t)), p.N	
}
