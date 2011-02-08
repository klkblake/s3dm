package s3dm

// This struct should be implemented by others that have a phyical position and
// rotation

// TODO: Scale
type Xform struct {
	Mat3
	pos *V3
}

func NewXform() *Xform {
	t := new(Xform)
	t.SetIdentity()
	t.pos = NewV3(0, 0, 0)
	return t
}

func (t *Xform) Copy() *Xform {
	o := new(Xform)
	o.Mat3 = *t.Mat3.Copy()
	o.pos = t.pos.Copy()
	return o
}

func (t *Xform) Position() *V3 {
	return t.pos
}

func (t *Xform) SetPosition(v *V3) {
	t.pos.Set(v)
}
