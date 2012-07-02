package s3dm

import "strconv"

type Mat3 [3 * 3]float64

var Mat3Identity = Mat3{
	1, 0, 0,
	0, 1, 0,
	0, 0, 1,
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

func (m Mat3) RawMatrix() [3 * 3]float64 {
	return [3 * 3]float64(m)
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
