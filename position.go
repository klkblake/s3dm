package s3dm

import "github.com/klkblake/fixed"

type Position struct {
	X, Y, Z fixed.Fixed
}

func (p Position) Add(o Position) Position {
	return Position{p.X.Add(o.X), p.Y.Add(o.Y), p.Z.Add(o.Z)}
}

func (p Position) Addf(v V3) Position {
	return Position{p.X.Add(fixed.New(v.X)), p.Y.Add(fixed.New(v.Y)), p.Z.Add(fixed.New(v.Z))}
}

func (p Position) Sub(o Position) Position {
	return Position{p.X.Sub(o.X), p.Y.Sub(o.Y), p.Z.Sub(o.Z)}
}

func (p Position) V3() V3 {
	return V3{p.X.Float64(), p.Y.Float64(), p.Z.Float64()}
}
