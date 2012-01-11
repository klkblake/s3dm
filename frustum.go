package s3dm

import "math"

type Frustum struct {
	Planes [6]*Plane
	Xform
	Near float64
	Far float64
	Fovy float64
	Aspect float64
}

func NewFrustum(near float64, far float64, fovy float64, aspect float64) *Frustum {
	frustum := new(Frustum)
	frustum.ResetXform()
	frustum.Near = near
	frustum.Far = far
	frustum.Fovy = fovy
	frustum.Aspect = aspect
	frustum.Update()
	return frustum
}

func (frustum *Frustum) Update() {
	lookAt := frustum.Forward().Unit()
	angleY := frustum.Fovy * 0.5
	angleX := angleY * frustum.Aspect
	// Near
	frustum.Planes[0] = NewPlane(frustum.Position().Add(lookAt.Muls(frustum.Near)), lookAt)
	// Far
	frustum.Planes[1] = NewPlane(frustum.Position().Add(lookAt.Muls(frustum.Far)), lookAt.Muls(-1))
	// Top
	frustum.Planes[2] = NewPlane(frustum.Position(), frustum.Mulv(NewV3(0, -math.Cos(angleY), -math.Sin(angleY))))
	// Bottom
	frustum.Planes[3] = NewPlane(frustum.Position(), frustum.Mulv(NewV3(0, math.Cos(angleY), -math.Sin(angleY))))
	// Left
	frustum.Planes[4] = NewPlane(frustum.Position(), frustum.Mulv(NewV3(math.Cos(angleX), 0, -math.Sin(angleX))))
	// Right
	frustum.Planes[5] = NewPlane(frustum.Position(), frustum.Mulv(NewV3(-math.Cos(angleX), 0, -math.Sin(angleX))))
}
