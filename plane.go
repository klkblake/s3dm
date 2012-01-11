package s3dm

type Plane struct {
	Xform
	n *V3
}

func NewPlane(o, n *V3) *Plane {
	p := new(Plane)
	p.ResetXform()
	p.n = NewV3(0,0,0)

	p.SetO(o)
	p.SetN(n)
	return p
}

func (p *Plane) O() *V3 {
	return p.Position()
}

func (p *Plane) N() *V3 {
	return p.Mulv(p.n)
}

func (p *Plane) SetO(o *V3) {
	p.SetPosition(o)
}

func (p *Plane) SetN(n *V3) {
	p.SetIdentity() // Clear rotations
	p.n.Set(n.Unit())
}

func (p *Plane) Intersect(r *Ray) (*V3, *V3) {
	po, pn := p.O(), p.N()
	ro, rd := r.O(), r.D()

	denom := pn.Dot(rd)
	if denom == 0 {
		return nil, nil
	}
	c := pn.Dot(po)
	t := (c - p.n.Dot(ro)) / denom
	if t <= 0 {
		return nil, nil
	}
	// If hitting underside; flip plane normal
	if denom > 0 {
		return ro.Add(rd.Muls(t)), pn.Muls(-1)
	}
	return ro.Add(rd.Muls(t)), pn
}

func (p *Plane) Side(point *V3) int {
	dot := p.N().Dot(point)
	if dot > 0 {
		return 1
	} else if dot < 0 {
		return -1
	}
	return 0
}
