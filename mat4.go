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

func (m *Mat4) SetRightUpForward(right, up, forward *V3) {
	m.matrix[0] = right.X; m.matrix[1] = right.Y; m.matrix[2] = right.Z
	m.matrix[4] = up.X; m.matrix[5] = up.Y; m.matrix[6] = up.Z
	m.matrix[8] = forward.X; m.matrix[9] = forward.Y; m.matrix[10] = forward.Z
}

// Get matrix rotation as Euler angles in degrees
func (m *Mat4) GetEuler() *V3 {
	x := math.Atan((-m.matrix[6]) / m.matrix[10])
    y := math.Asin(m.matrix[2])
    z := math.Atan((-m.matrix[1]) / m.matrix[0])

	// Convert to Degrees
	x *= 180 / math.Pi
	y *= 180 / math.Pi
	z *= 180 / math.Pi

	return NewV3(x, y, z)
}

// Set matrix rotation to Euler angles in degrees
func (m *Mat4) SetEuler(r *V3) {
	// Convert to Radians
	r.X *= math.Pi / 180
	r.Y *= math.Pi / 180
	r.Z *= math.Pi / 180

	m.matrix[0] = math.Cos(r.Y) * math.Cos(r.Z)
	m.matrix[1] = -math.Cos(r.Y) * math.Sin(r.Z)
	m.matrix[2] = math.Sin(r.Y)

	m.matrix[4] = math.Sin(r.X) * math.Sin(r.Y) * math.Cos(r.Z) + 
		math.Cos(r.X)*math.Sin(r.Z)
	m.matrix[5] = -math.Sin(r.X) * math.Sin(r.Y) * math.Sin(r.Z) + 
		math.Cos(r.X) * math.Cos(r.Z)
	m.matrix[6] = -math.Sin(r.X) * math.Cos(r.Y)

	m.matrix[8] = -math.Cos(r.X) * math.Sin(r.Y) * math.Cos(r.Z) + 
		math.Sin(r.X) * math.Sin(r.Z)
	m.matrix[9] = math.Cos(r.X) * math.Sin(r.Y) * math.Sin(r.Z) + 
		math.Sin(r.X) * math.Cos(r.Z)
	m.matrix[10] = math.Cos(r.X) * math.Cos(r.Y)
}

// Set matrix rotation to quateronion 'q'
func (m *Mat4) SetQuaternion(q *Qtrnn) {
	xx, xy, xz, xw := q.X*q.X, q.X*q.Y, q.X*q.Z, q.X*q.W
	yy, yz, yw := q.Y*q.Y, q.Y*q.Z, q.Y*q.W
	zz, zw := q.Z*q.Z, q.Z*q.W
	
	m.matrix[0] = 1.0 - 2.0 * (yy + zz)
	m.matrix[1] = 2.0 * (xy - zw)
	m.matrix[2] = 2.0 * (xz + yw)
	m.matrix[4] = 2.0 * (xy + zw)
	m.matrix[5] = 1.0 - 2.0 * (xx + zz)
	m.matrix[6] = 2.0 * (yz - xw)
	m.matrix[8] = 2.0 * (xz - yw)
	m.matrix[9] = 2.0 * (yz + xw)
	m.matrix[10] = 1.0 - 2.0 * (xx + yy)
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


