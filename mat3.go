package s3dm

import "math"
import "strconv"

type Mat3 [3 * 3]float64

var Mat3Identity = Mat3{
	1, 0, 0,
	0, 1, 0,
	0, 0, 1,
}

func (m Mat3) Abs() Mat3 {
	var ret Mat3
	for i := 0; i < 3*3; i++ {
		if m[i] >= 0 {
			ret[i] = m[i]
		} else {
			ret[i] = -m[i]
		}
	}
	return ret
}

func (m Mat3) Rotate(angle float64, axis V3) Mat3 {
	return m.rotate(angle, axis)
}

func (m *Mat3) RotateLocal(angle float64, axis V3) {
	*m = m.rotate(angle, axis)
}

func (m Mat3) GetMatrix() [3 * 3]float64 {
	return [3*3]float64(m)
}

func (m Mat3) Right() V3 {
	return V3{m[0], m[3], m[6]}
}

func (m Mat3) Up() V3 {
	return V3{m[1], m[4], m[7]}
}

func (m Mat3) Forward() V3 {
	return V3{m[2], m[5], m[8]}
}

func (m *Mat3) SetRightUpForward(right, up, forward V3) {
	m[0] = right.X
	m[3] = right.Y
	m[6] = right.Z
	m[1] = up.X
	m[4] = up.Y
	m[7] = up.Z
	m[2] = forward.X
	m[5] = forward.Y
	m[8] = forward.Z
}

// Get matrix rotation as Euler angles in degrees
func (m Mat3) GetEuler() V3 {
	x := math.Atan((-m[7]) / m[8])
	y := math.Asin(m[6])
	z := math.Atan((-m[3]) / m[0])

	// Convert to Degrees
	x *= 180 / math.Pi
	y *= 180 / math.Pi
	z *= 180 / math.Pi

	return V3{x, y, z}
}

// Set matrix rotation to Euler angles in degrees
func (m *Mat3) SetEuler(r V3) {
	// Convert to Radians
	r.X *= math.Pi / 180
	r.Y *= math.Pi / 180
	r.Z *= math.Pi / 180

	m[0] = math.Cos(r.Y) * math.Cos(r.Z)
	m[3] = -math.Cos(r.Y) * math.Sin(r.Z)
	m[6] = math.Sin(r.Y)

	m[1] = math.Sin(r.X)*math.Sin(r.Y)*math.Cos(r.Z) +
		math.Cos(r.X)*math.Sin(r.Z)
	m[4] = -math.Sin(r.X)*math.Sin(r.Y)*math.Sin(r.Z) +
		math.Cos(r.X)*math.Cos(r.Z)
	m[7] = -math.Sin(r.X) * math.Cos(r.Y)

	m[2] = -math.Cos(r.X)*math.Sin(r.Y)*math.Cos(r.Z) +
		math.Sin(r.X)*math.Sin(r.Z)
	m[5] = math.Cos(r.X)*math.Sin(r.Y)*math.Sin(r.Z) +
		math.Sin(r.X)*math.Cos(r.Z)
	m[8] = math.Cos(r.X) * math.Cos(r.Y)
}

// Set matrix rotation to quateronion 'q'
func (m *Mat3) SetQuaternion(q *Qtrnn) {
	xx, xy, xz, xw := q.X*q.X, q.X*q.Y, q.X*q.Z, q.X*q.W
	yy, yz, yw := q.Y*q.Y, q.Y*q.Z, q.Y*q.W
	zz, zw := q.Z*q.Z, q.Z*q.W

	m[0] = 1.0 - 2.0*(yy+zz)
	m[3] = 2.0 * (xy - zw)
	m[6] = 2.0 * (xz + yw)
	m[1] = 2.0 * (xy + zw)
	m[4] = 1.0 - 2.0*(xx+zz)
	m[7] = 2.0 * (yz - xw)
	m[2] = 2.0 * (xz - yw)
	m[5] = 2.0 * (yz + xw)
	m[8] = 1.0 - 2.0*(xx+yy)
}

// Multiply 'm' by 'o' and return result
func (m Mat3) Mul(o Mat3) Mat3 {
	var result Mat3
	for row := 0; row < 3; row++ {
		ca := row
		cb := ca + 3
		cc := ca + 6
		result[ca] =
			m[ca]*o[0] +
				m[cb]*o[1] +
				m[cc]*o[2]
		result[cb] =
			m[ca]*o[3] +
				m[cb]*o[4] +
				m[cc]*o[5]
		result[cc] =
			m[ca]*o[6] +
				m[cb]*o[7] +
				m[cc]*o[8]
	}
	return result
}

// Multiply 'm' by 'v' and return result
func (m Mat3) Mulv(v V3) V3 {
	return V3{
		v.X*m[0] + v.Y*m[3] + v.Z*m[6],
		v.X*m[1] + v.Y*m[4] + v.Z*m[7],
		v.X*m[2] + v.Y*m[5] + v.Z*m[8]}
}

// Unexported rotate used by RotateLocal & RotateGlobal
func (m Mat3) rotate(angle float64, axis V3) Mat3 {
	sinAngle := math.Sin(angle * math.Pi / 180)
	cosAngle := math.Cos(angle * math.Pi / 180)
	oneMinusCos := float64(1) - cosAngle

	axis = axis.Unit()

	xx := axis.X * axis.X
	yy := axis.Y * axis.Y
	zz := axis.Z * axis.Z
	xy := axis.X * axis.Y
	yz := axis.Y * axis.Z
	zx := axis.Z * axis.X
	xs := axis.X * sinAngle
	ys := axis.Y * sinAngle
	zs := axis.Z * sinAngle

	var rotation Mat3

	rotation[0] = (oneMinusCos * xx) + cosAngle
	rotation[3] = (oneMinusCos * xy) - zs
	rotation[6] = (oneMinusCos * zx) + ys

	rotation[1] = (oneMinusCos * xy) + zs
	rotation[4] = (oneMinusCos * yy) + cosAngle
	rotation[7] = (oneMinusCos * yz) - xs

	rotation[2] = (oneMinusCos * zx) - ys
	rotation[5] = (oneMinusCos * yz) + xs
	rotation[8] = (oneMinusCos * zz) + cosAngle

	return rotation.Mul(m)
}

func (m Mat3) String() string {
	s := "["
	for col := 0; col < 3; col++ {
		for row := 0; row < 3; row++ {
			s += strconv.FormatFloat(m[col+row*3], 'g', 2, 64)
			if col != 2 || row != 2 {
				s += ", "
			}
		}
	}
	s += "]"
	return s
}
