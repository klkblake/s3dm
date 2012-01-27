package s3dm

type Ray struct {
	Xform
	d V3
}

func NewRay(o, d V3) *Ray {
	r := new(Ray)
	r.ResetXform()

	r.SetO(o)
	r.SetD(d)
	return r
}

func (r *Ray) O() V3 {
	return r.Position()
}

func (r *Ray) D() V3 {
	return r.Mulv(r.d)
}

func (r *Ray) SetO(o V3) {
	r.SetPosition(o)
}

func (r *Ray) SetD(d V3) {
	r.SetIdentity() // Reset rotations
	r.d = d.Unit()
}

func (r *Ray) Advance(a float64) {
	r.MoveLocal(r.D().Muls(a))
}
