package s3dm

type Plane struct {
	Xform
	n V3
	d float64
}

func NewPlane(o, n V3) *Plane {
	p := new(Plane)
	p.ResetXform()

	p.SetO(o)
	p.SetN(n)
	return p
}

func (p *Plane) O() V3 {
	return p.Position
}

func (p *Plane) N() V3 {
	return p.n
}

func (p *Plane) SetO(o V3) {
	p.Position = o
	p.d = -p.n.Dot(o)
}

func (p *Plane) SetN(n V3) {
	p.Xform.Mat3 = Mat3Identity // Clear rotations
	p.n = n.Unit()
	p.d = -p.n.Dot(p.Position)
}

func (p *Plane) RotateLocal(angle float64, axis V3) {
	p.Xform.RotateLocal(angle, axis)
	p.SetN(p.Mulv(p.n))
}

func (p *Plane) SetEuler(r V3) {
	p.Xform.SetEuler(r)
	p.SetN(p.Mulv(p.n))
}

func (p *Plane) SetQuaternion(q *Qtrnn) {
	p.Xform.SetQuaternion(q)
	p.SetN(p.Mulv(p.n))
}

func (p *Plane) SetRightUpForward(right, up, forward V3) {
	p.Xform.SetRightUpForward(right, up, forward)
	p.SetN(p.Mulv(p.n))
}

func (p *Plane) Intersect(r *Ray) (V3, V3) {
	po, pn := p.O(), p.N()
	ro, rd := r.O(), r.D()

	denom := pn.Dot(rd)
	if denom == 0 {
		return V3{}, V3{}
	}
	c := pn.Dot(po)
	t := (c - p.n.Dot(ro)) / denom
	if t <= 0 {
		return V3{}, V3{}
	}
	// If hitting underside; flip plane normal
	if denom > 0 {
		return ro.Add(rd.Muls(t)), pn.Muls(-1)
	}
	return ro.Add(rd.Muls(t)), pn
}

func (p *Plane) Side(point V3) int {
	dot := p.n.Dot(point) + p.d
	if dot > 0 {
		return 1
	} else if dot < 0 {
		return -1
	}
	return 0
}
