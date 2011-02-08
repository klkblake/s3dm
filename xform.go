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
	t.ResetXform()
	return t
}

func (t *Xform) ResetXform() {
	t.SetIdentity()
	t.pos = NewV3(0, 0, 0)
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

func (t *Xform) GetMatrix4() [4*4]float64 {
	result := [4*4]float64 {}
	
	// Set rotation
	m := t.GetMatrix()
	result[0] = m[0]
	result[1] = m[1]
	result[2] = m[2]
	result[4] = m[3]
	result[5] = m[4]
	result[6] = m[5]
	result[8] = m[6]
	result[9] = m[7]
	result[10] = m[8]

	// Set Position
	result[12] = t.Position().X
	result[13] = t.Position().Y
	result[14] = t.Position().Z

	// Set Identity
	result[15] = 1

	return result
	
}
