package s3dm

// This struct should be implemented by others that have a physical position
// and rotation

type Xform struct {
	Mat3
	Position V3
	Scale    V3
}

func NewXform() *Xform {
	t := new(Xform)
	t.ResetXform()
	return t
}

func (t *Xform) ResetXform() {
	t.SetIdentity()
	t.Position = V3{0, 0, 0}
	t.Scale = V3{1, 1, 1}
}

func (t *Xform) Copy() *Xform {
	o := new(Xform)
	o.Mat3 = *t.Mat3.Copy()
	o.Position = t.Position
	o.Scale = t.Scale
	return o
}

func (t *Xform) GetMatrix4() [4 * 4]float64 {
	result := [4 * 4]float64{}

	// Set rotation
	m := t.GetMatrix()
	result[0] = m[0] * t.Scale.X
	result[1] = m[1] * t.Scale.X
	result[2] = m[2] * t.Scale.X
	result[4] = m[3] * t.Scale.Y
	result[5] = m[4] * t.Scale.Y
	result[6] = m[5] * t.Scale.Y
	result[8] = m[6] * t.Scale.Z
	result[9] = m[7] * t.Scale.Z
	result[10] = m[8] * t.Scale.Z

	// Set Position
	result[12] = t.Position.X
	result[13] = t.Position.Y
	result[14] = t.Position.Z

	// Set Identity
	result[15] = 1

	return result
}
