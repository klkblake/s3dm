package s3dm

// This struct should be implemented by others that have a physical position
// and rotation

type Xform struct {
	Mat3
	pos *V3
	scale *V3
}

func NewXform() *Xform {
	t := new(Xform)
	t.ResetXform()
	return t
}

func (t *Xform) ResetXform() {
	t.SetIdentity()
	t.pos = NewV3(0, 0, 0)
	t.scale = NewV3(1, 1, 1)
}

func (t *Xform) Copy() *Xform {
	o := new(Xform)
	o.Mat3 = *t.Mat3.Copy()
	o.pos = t.pos.Copy()
	o.scale = t.scale.Copy()
	return o
}

func (t *Xform) Position() *V3 {
	return t.pos
}

func (t *Xform) SetPosition(v *V3) {
	t.pos.Set(v)
}

func (t *Xform) MoveGlobal(v *V3) {
	t.SetPosition(t.Position().Add(v))
}

func (t *Xform) MoveLocal(v *V3) {
	m := t.GetMatrix()
	d := NewV3(
		v.X * m[0] + v.Y * m[1] + v.Z * m[2],
		v.X * m[3] + v.Y * m[4] + v.Z * m[5],
		v.X * m[6] + v.Y * m[7] + v.Z * m[8])
	t.SetPosition(t.Position().Add(d))
}

func (t *Xform) Scale() *V3 {
	return t.scale
}

func (t *Xform) SetScale(v *V3) {
	t.scale.Set(v)
}

func (t *Xform) ScaleGlobal(v *V3) {
	t.SetScale(t.Scale().Add(v))
}

func (t *Xform) ScaleLocal(v *V3) {
	m := t.GetMatrix()
	d := NewV3(
		v.X * m[0] + v.Y * m[1] + v.Z * m[2],
		v.X * m[3] + v.Y * m[4] + v.Z * m[5],
		v.X * m[6] + v.Y * m[7] + v.Z * m[8])
	t.SetScale(t.Scale().Add(d))
}

func (t *Xform) GetMatrix4() [4*4]float64 {
	result := [4*4]float64 {}

	// Set rotation
	m := t.GetMatrix()
	result[0] = m[0] * t.Scale().X
	result[1] = m[3] * t.Scale().X
	result[2] = m[6] * t.Scale().X
	result[4] = m[1] * t.Scale().Y
	result[5] = m[4] * t.Scale().Y
	result[6] = m[7] * t.Scale().Y
	result[8] = m[2] * t.Scale().Z
	result[9] = m[5] * t.Scale().Z
	result[10] = m[8] * t.Scale().Z

	// Set Position
	result[12] = t.Position().X
	result[13] = t.Position().Y
	result[14] = t.Position().Z

	// Set Identity
	result[15] = 1

	return result
}
