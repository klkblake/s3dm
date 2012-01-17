package s3dm

import "math"
import "strconv"

type Mat3 struct {
	matrix [3*3]float64
}

func NewMat3() *Mat3 {
	m := new(Mat3)
	m.SetIdentity()
	return m
}

func (m *Mat3) Copy() *Mat3 {
	n := NewMat3()
	for i := 0; i < 3*3; i += 1 {
		n.matrix[i] = m.matrix[i]
	}
	return n
}

func (m *Mat3) SetIdentity() {
	m.matrix = [3*3]float64 {
		1, 0, 0,
		0, 1, 0,
		0, 0, 1 }
}

func (m *Mat3) RotateLocal(angle float64, axis *V3) {
	*m = *m.rotate(angle, axis)
}

func (m *Mat3) RotateGlobal(angle float64, axis *V3) {
	axis = m.Mulv(axis)
	*m = *m.rotate(angle, axis)
}

func (m *Mat3) GetMatrix() [3*3]float64 {
	return m.matrix
}

func (m *Mat3) Right() *V3 {
	return NewV3(m.matrix[0], m.matrix[1], m.matrix[2])
}

func (m *Mat3) Up() *V3 {
	return NewV3(m.matrix[3], m.matrix[4], m.matrix[5])
}

func (m *Mat3) Forward() *V3 {
	return NewV3(m.matrix[6], m.matrix[7], m.matrix[8])
}

func (m *Mat3) SetRightUpForward(right, up, forward *V3) {
	m.matrix[0] = right.X; m.matrix[1] = right.Y; m.matrix[2] = right.Z
	m.matrix[3] = up.X; m.matrix[4] = up.Y; m.matrix[5] = up.Z
	m.matrix[6] = forward.X; m.matrix[7] = forward.Y; m.matrix[8] = forward.Z
}

// Get matrix rotation as Euler angles in degrees
func (m *Mat3) GetEuler() *V3 {
	x := math.Atan((-m.matrix[5]) / m.matrix[8])
    y := math.Asin(m.matrix[2])
    z := math.Atan((-m.matrix[1]) / m.matrix[0])

	// Convert to Degrees
	x *= 180 / math.Pi
	y *= 180 / math.Pi
	z *= 180 / math.Pi

	return NewV3(x, y, z)
}

// Set matrix rotation to Euler angles in degrees
func (m *Mat3) SetEuler(r *V3) {
	// Convert to Radians
	r.X *= math.Pi / 180
	r.Y *= math.Pi / 180
	r.Z *= math.Pi / 180

	m.matrix[0] = math.Cos(r.Y) * math.Cos(r.Z)
	m.matrix[1] = -math.Cos(r.Y) * math.Sin(r.Z)
	m.matrix[2] = math.Sin(r.Y)

	m.matrix[3] = math.Sin(r.X) * math.Sin(r.Y) * math.Cos(r.Z) +
		math.Cos(r.X)*math.Sin(r.Z)
	m.matrix[4] = -math.Sin(r.X) * math.Sin(r.Y) * math.Sin(r.Z) +
		math.Cos(r.X) * math.Cos(r.Z)
	m.matrix[5] = -math.Sin(r.X) * math.Cos(r.Y)

	m.matrix[6] = -math.Cos(r.X) * math.Sin(r.Y) * math.Cos(r.Z) +
		math.Sin(r.X) * math.Sin(r.Z)
	m.matrix[7] = math.Cos(r.X) * math.Sin(r.Y) * math.Sin(r.Z) +
		math.Sin(r.X) * math.Cos(r.Z)
	m.matrix[8] = math.Cos(r.X) * math.Cos(r.Y)
}

// Set matrix rotation to quateronion 'q'
func (m *Mat3) SetQuaternion(q *Qtrnn) {
	xx, xy, xz, xw := q.X*q.X, q.X*q.Y, q.X*q.Z, q.X*q.W
	yy, yz, yw := q.Y*q.Y, q.Y*q.Z, q.Y*q.W
	zz, zw := q.Z*q.Z, q.Z*q.W

	m.matrix[0] = 1.0 - 2.0 * (yy + zz)
	m.matrix[1] = 2.0 * (xy - zw)
	m.matrix[2] = 2.0 * (xz + yw)
	m.matrix[3] = 2.0 * (xy + zw)
	m.matrix[4] = 1.0 - 2.0 * (xx + zz)
	m.matrix[5] = 2.0 * (yz - xw)
	m.matrix[6] = 2.0 * (xz - yw)
	m.matrix[7] = 2.0 * (yz + xw)
	m.matrix[8] = 1.0 - 2.0 * (xx + yy)
}

// Multiply 'm' by 'o' and return result
func (m *Mat3) Mul(o *Mat3) *Mat3 {
	result := NewMat3()

	for row := 0; row < 3; row++ {
		ca := 3 * row
		cb := ca + 1
		cc := ca + 2

		result.matrix[ca] = 
			m.matrix[ca] * o.matrix[0] +
			m.matrix[cb] * o.matrix[3] +
			m.matrix[cc] * o.matrix[6]

		result.matrix[cb] =
			m.matrix[ca] * o.matrix[1] +
			m.matrix[cb] * o.matrix[4] +
			m.matrix[cc] * o.matrix[7]

		result.matrix[cc] = 
			m.matrix[ca] * o.matrix[2] +
			m.matrix[cb] * o.matrix[5] +
			m.matrix[cc] * o.matrix[8]
	}

	return result
}

// Multiply 'm' by 'v' and return result
func (m *Mat3) Mulv(v *V3) *V3 {
	return NewV3(
		v.X * m.matrix[0] + v.Y * m.matrix[1] + v.Z * m.matrix[2],
		v.X * m.matrix[3] + v.Y * m.matrix[4] + v.Z * m.matrix[5],
		v.X * m.matrix[6] + v.Y * m.matrix[7] + v.Z * m.matrix[8])
}

// Unexported rotate used by RotateLocal & RotateGlobal
func (m *Mat3) rotate(angle float64, axis *V3) *Mat3 {
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

	rotation := NewMat3()

	rotation.matrix[0] = (oneMinusCos * xx) + cosAngle;
	rotation.matrix[1] = (oneMinusCos * xy) - zs;
	rotation.matrix[2] = (oneMinusCos * zx) + ys;

	rotation.matrix[3] = (oneMinusCos * xy) + zs;
	rotation.matrix[4] = (oneMinusCos * yy) + cosAngle;
	rotation.matrix[5] = (oneMinusCos * yz) - xs;

	rotation.matrix[6] = (oneMinusCos * zx) - ys;
	rotation.matrix[7] = (oneMinusCos * yz) + xs;
	rotation.matrix[8] = (oneMinusCos * zz) + cosAngle;

	return rotation.Mul(m)
}

func (m *Mat3) String() string {
	s := "["
	for i := 0; i < 9; i += 1 {
		s += strconv.FormatFloat(m.matrix[i], 'e', 2, 64)
		if i < 8 {
			s += ", "
		}
	}
	s += "]"
	return s
}

