package s3dm

type Plane struct {
	Xform
}

func NewPlane(origin, normal *V3) *Plane {
	p := new(Plane)
	p.SetIdentity()
	p.SetPosition(origin)

	up := normal
	forward := up.Perp()
	right := up.Cross(forward)

	p.SetRightUpForward(right, up, forward)
	return p
}

func (p *Plane) Intersect(r *Ray) (*V3, *V3) {
	normal := p.Up()
	o := p.Position()

	denom := normal.Dot(r.D)
	if denom == 0 {
		return nil, nil
	}
	c := normal.Dot(o)
	t := (c - normal.Dot(r.O)) / denom
	if t <= 0 {
		return nil, nil
	}
	return r.O.Add(r.D.Muls(t)), normal	
}
