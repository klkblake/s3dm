package s3dm


import "math"

var QtrnnIdentity = Qtrnn{1, 0, 0, 0}

type Qtrnn struct {
	X, Y, Z, W float64
}

func AxisAngle(axis V3, angle float64) Qtrnn {
	a := axis.Muls(math.Sin(angle/2))
	return Qtrnn{a.X, a.Y, a.Z, math.Cos(angle*0.5)}
}

func (q Qtrnn) Mul(p Qtrnn) Qtrnn {
	x1, y1, z1, w1 := q.X, q.Y, q.Z, q.W
	x2, y2, z2, w2 := p.X, p.Y, p.Z, p.W
	return Qtrnn{
		X: x1*w2 + y1*z2 - z1*y2 + w1*x2,
		Y: -x1*z2 + y1*w2 + z1*x2 + w1*y2,
		Z: x1*y2 - y1*x2 + z1*w2 + w1*z2,
		W: -x1*x2 - y1*y2 - z1*z2 + w1*w2,
	}
}

func (q Qtrnn) Matrix() Mat3 {
	x, y, z, w := q.X, q.Y, q.Z, q.W
	xx, xy, xz, xw := x*x, x*y, x*z, x*w
	yy, yz, yw := y*y, y*z, y*w
	zz, zw := z*z, z*w
	ww := w * w
	return Mat3{
		xx - yy - zz + ww, 2 * (xy + zw), 2 * (xz - yw),
		2 * (xy - zw), -xx + yy - zz + ww, 2 * (yz + xw),
		2 * (xz + yw), 2 * (yz - xw), -xx - yy + zz + ww,
	}
}
