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
	return r.Position
}

func (r *Ray) D() V3 {
	return r.Mulv(r.d)
}

func (r *Ray) SetO(o V3) {
	r.Position = o
}

func (r *Ray) SetD(d V3) {
	r.Xform.Mat3 = Mat3Identity // Reset rotations
	r.d = d.Unit()
}

func (r *Ray) Advance(a float64) {
	r.Position.AddLocal(r.Mulv(r.D().Muls(a)))
}
