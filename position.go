package s3dm

type Position struct {
	X, Y, Z Fixed64
}

func (p Position) Add(v V3) Position {
	return Position{p.X + NewFixed64(v.X), p.Y + NewFixed64(v.Y), p.Z + NewFixed64(v.Z)}
}

func (p Position) Sub(o Position) V3 {
	return V3{(p.X - o.X).Float64(), (p.Y - o.Y).Float64(), (p.Z - o.Z).Float64()}
}
