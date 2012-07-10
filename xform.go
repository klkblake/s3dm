package s3dm

// This struct should be implemented by others that have a physical position
// and rotation

var XformIdentity = Xform{
	V3{0, 0, 0},
	QtrnnIdentity,
}

type Xform struct {
	Position V3
	Rotation Qtrnn
}

func (xf Xform) Matrix() (result Mat4) {
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
	result[12] = xf.Position.X
	result[13] = xf.Position.Y
	result[14] = xf.Position.Z

	// Set Identity
	result[15] = 1
	return
}
