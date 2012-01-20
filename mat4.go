package s3dm

import (
	"math"
	"strconv"
)

// 4x4 matrix in column-major order.
type Mat4 [4 * 4]float64

func NewMat4() *Mat4 {
	matrix := new(Mat4)
	matrix.SetIdentity()
	return matrix
}

func NewPerspectiveMat4(fovy, aspect, near, far float64) *Mat4 {
	top := near * math.Tan(fovy*0.5)
	right := aspect * top
	return &Mat4{
		near / right, 0, 0, 0,
		0, near / top, 0, 0,
		0, 0, -(far + near) / (far - near), -1,
		0, 0, -2 * far * near / (far - near), 0}
}

func NewOrthographicMat4(width, height, near, far float64) *Mat4 {
	return &Mat4{
		2/width, 0, 0, 0,
		0, 2/height, 0, 0,
		0, 0, 2/(near-far), 0,
		-1, -1, (near+far)/(near-far), 1}
}

func (m *Mat4) SetIdentity() {
	*m = [4 * 4]float64{
		1, 0, 0, 0,
		0, 1, 0, 0,
		0, 0, 1, 0,
		0, 0, 0, 1}
}

func (m *Mat4) GetMatrix() [4 * 4]float64 {
	return [4 * 4]float64(*m)
}

func (m *Mat4) GetFloat32Matrix() [4 * 4]float32 {
	matrix := new([4 * 4]float32)
	for i, v := range m {
		matrix[i] = float32(v)
	}
	return *matrix
}

func (m *Mat4) Position() *V3 {
	return NewV3(m[12], m[13], m[14])
}

func (m *Mat4) Mul(o *Mat4) *Mat4 {
	res := new(Mat4)
	for row := 0; row < 4; row++ {
		ca := row
		cb := ca + 4
		cc := ca + 8
		cd := ca + 12
		res[ca] =
			m[ca]*o[0] +
			m[cb]*o[1] +
			m[cc]*o[2] +
			m[cd]*o[3]
		res[cb] =
			m[ca]*o[4] +
			m[cb]*o[5] +
			m[cc]*o[6] +
			m[cd]*o[7]
		res[cc] =
			m[ca]*o[8] +
			m[cb]*o[9] +
			m[cc]*o[10] +
			m[cd]*o[11]
		res[cd] =
			m[ca]*o[12] +
			m[cb]*o[13] +
			m[cc]*o[14] +
			m[cd]*o[15]
	}
	return res
}

func (m *Mat4) Transpose() *Mat4 {
	return &Mat4{
		m[0], m[4], m[8], m[12],
		m[1], m[5], m[9], m[13],
		m[2], m[6], m[10], m[14],
		m[3], m[7], m[11], m[15]}
}

func (m *Mat4) String() string {
	s := "["
	for col := 0; col < 4; col++ {
		for row := 0; row < 4; row++ {
			s += strconv.FormatFloat(m[col+row*4], 'g', 2, 64)
			if col != 3 || row != 3 {
				s += ", "
			}
		}
	}
	s += "]"
	return s
}
