package s3dm

import "math"
import "strconv"

// ----- Matrix math helper functions ------------------------------------------

func Matrix4MultiplyVector3(m [4*4]float64, v *V3) *V3 {
	return NewV3(
		v.X * m[0] + v.Y * m[1] + v.Z * m[2],
		v.X * m[4] + v.Y * m[5] + v.Z * m[6],
		v.X * m[8] + v.Y * m[9] + v.Z * m[10])
}

func Matrix4Multiply(a, b [4*4]float64) [4*4]float64 {
	var result [4*4]float64

	for row := 0; row < 4; row++ {
		ca := 4 * row
		cb := ca + 1
		cc := ca + 2
		cd := ca + 3

		result[ca] = 
			a[ca] * b[0] +
			a[cb] * b[4] +
			a[cc] * b[8] +
			a[cd] * b[12]

		result[cb] =
			a[ca] * b[1] +
			a[cb] * b[5] +
			a[cc] * b[9] +
			a[cd] * b[13]

		result[cc] = 
			a[ca] * b[2] +
			a[cb] * b[6] +
			a[cc] * b[10] +
			a[cd] * b[14]

		result[cd] = 
			a[ca] * b[3] +
			a[cb] * b[7] +
			a[cc] * b[11] +
			a[cd] * b[15]
	}

	return result
}

func Matrix4Rotate(m [4*4]float64, angle float64, axis *V3) [4*4]float64 {
	sinAngle := math.Sin(angle * math.Pi / 180)
	cosAngle := math.Cos(angle * math.Pi / 180)
	oneMinusCos := float64(1) - cosAngle

	axis = axis.Unit()

	xx := axis.X * axis.X;
	yy := axis.Y * axis.Y;
	zz := axis.Z * axis.Z;
	xy := axis.X * axis.Y;
	yz := axis.Y * axis.Z;
	zx := axis.Z * axis.X;
	xs := axis.X * sinAngle;
	ys := axis.Y * sinAngle;
	zs := axis.Z * sinAngle;

	var rotationMatrix [4*4]float64

	rotationMatrix[0] = (oneMinusCos * xx) + cosAngle;
	rotationMatrix[1] = (oneMinusCos * xy) - zs;
	rotationMatrix[2] = (oneMinusCos * zx) + ys;
	rotationMatrix[3] = 0;

	rotationMatrix[4] = (oneMinusCos * xy) + zs;
	rotationMatrix[5] = (oneMinusCos * yy) + cosAngle;
	rotationMatrix[6] = (oneMinusCos * yz) - xs;
	rotationMatrix[7] = 0;

	rotationMatrix[8] = (oneMinusCos * zx) - ys;
	rotationMatrix[9] = (oneMinusCos * yz) + xs;
	rotationMatrix[10] = (oneMinusCos * zz) + cosAngle;
	rotationMatrix[11] = 0;

	rotationMatrix[12] = 0;
	rotationMatrix[13] = 0;
	rotationMatrix[14] = 0;
	rotationMatrix[15] = 1;

	return Matrix4Multiply(rotationMatrix, m)
}

// ----- Mat4 struct -----------------------------------------------------------

type Mat4 struct {
	matrix [4*4]float64	
}

func identityMat4() [4*4]float64 {
	return [4*4]float64 {
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1 }
}

func NewMat4() *Mat4 {
	m := new(Mat4)
	m.SetIdentity()
	return m
}

func (m *Mat4) Copy() *Mat4 {
	n := NewMat4()
	for i := 0; i < 4*4; i += 1 {
		n.matrix[i] = m.matrix[i]
	}
	return n
}

func (m *Mat4) SetIdentity() {
	m.matrix = identityMat4()
}

func (m *Mat4) GetMatrix() [4*4]float64 {
	return m.matrix
}

func (m *Mat4) RotateLocal(angle float64, axis *V3) {
	m.matrix = Matrix4Rotate(m.matrix, angle, axis)
}

func (m *Mat4) RotateGlobal(angle float64, axis *V3) {
	axis = Matrix4MultiplyVector3(m.matrix, axis)
	m.matrix = Matrix4Rotate(m.matrix, angle, axis)
}

func (m *Mat4) MoveLocal(t *V3) {
	d := NewV3 (
		t.X * m.matrix[0] + t.Y * m.matrix[4] + t.Z * m.matrix[8],
		t.X * m.matrix[1] + t.Y * m.matrix[5] + t.Z * m.matrix[9],
		t.X * m.matrix[2] + t.Y * m.matrix[6] + t.Z * m.matrix[10])
	m.SetPosition(m.Position().Add(d));
}

func (m *Mat4) MoveGlobal(t *V3) {	
	m.SetPosition(m.Position().Add(t));
}

func (m *Mat4) Right() *V3 {
	return NewV3(m.matrix[0], m.matrix[1], m.matrix[2])
}

func (m *Mat4) Up() *V3 {
	return NewV3(m.matrix[4], m.matrix[5], m.matrix[6])
}

func (m *Mat4) Forward() *V3 {
	return NewV3(m.matrix[8], m.matrix[9], m.matrix[10])
}

func (m *Mat4) Position() *V3 {
	return NewV3(m.matrix[12], m.matrix[13], m.matrix[14])
}

func (m *Mat4) SetPosition(p *V3) {
	m.matrix[12] = p.X
	m.matrix[13] = p.Y
	m.matrix[14] = p.Z
}

func (m *Mat4) String() string {
	s := "["
	for i := 0; i < 16; i += 1 {
		s += strconv.Ftoa64(m.matrix[i], 'e', 2)
		if i < 15 {
			s += ", "
		}
	}
	s += "]"
	return s
}


