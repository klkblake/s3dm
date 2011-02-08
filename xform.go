package s3dm

// This struct should be implemented by others that have a phyical position and
// rotation

// TODO: Scale
type Transform struct {
	Mat3
	pos *V3
}

func NewTransform() *Transform {
	t := new(Transform)
	t.SetIdentity()
	t.pos = NewV3(0, 0, 0)
	return t
}

func (t *Transform) Copy() *Transform {
	o := new(Transform)
	o.Mat3 = *t.Mat3.Copy()
	o.pos = t.pos.Copy()
	return o
}

func (t *Transform) Position() *V3 {
	return t.pos
}

func (t *Transform) SetPosition(v *V3) {
	t.pos.Set(v)
}
