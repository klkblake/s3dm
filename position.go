package s3dm

import "github.com/klkblake/fixed"

type Position struct {
	X, Y, Z fixed.Fixed
}

func (p Position) Add(v V3) Position {
	return Position{p.X.Add(fixed.New(v.X)), p.Y.Add(fixed.New(v.Y)), p.Z.Add(fixed.New(v.Z))}
}

func (p Position) Sub(o Position) V3 {
	return V3{(p.X.Sub(o.X)).Float64(), (p.Y.Sub(o.Y)).Float64(), (p.Z.Sub(o.Z)).Float64()}
}
