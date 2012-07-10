package s3dm

var XformScaleIdentity = XformScale{
	XformIdentity,
	V3{1, 1, 1},
}

type XformScale struct {
	Xform
	Scale V3
}

func (xfs XformScale) Matrix() (result Mat4) {
	result = xfs.Xform.Matrix()
	result[0] *= xfs.Scale.X
	result[1] *= xfs.Scale.X
	result[2] *= xfs.Scale.X
	result[4] *= xfs.Scale.Y
	result[5] *= xfs.Scale.Y
	result[6] *= xfs.Scale.Y
	result[8] *= xfs.Scale.Z
	result[9] *= xfs.Scale.Z
	result[10] *= xfs.Scale.Z
	return
}
