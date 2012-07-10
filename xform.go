package s3dm

// This struct should be implemented by others that have a physical position
// and rotation

var XformIdentity = Xform{
	Position{},
	QtrnnIdentity,
}

type Xform struct {
	Position Position
	Rotation Qtrnn
}

func (xf Xform) Matrix(origin Position) (result Mat4) {
	// Set rotation
	m := xf.Rotation.Matrix()
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
	p := xf.Position.Sub(origin)
	result[12] = p.X
	result[13] = p.Y
	result[14] = p.Z

	// Set Identity
	result[15] = 1
	return
}
