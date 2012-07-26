package s3dm

type Plane struct {
	Origin Position
	Normal V3
}

func (p Plane) Intersect(r *Ray) (Position, V3) {
	po, pn := p.Origin, p.Normal
	ro, rd := r.Origin, r.Dir

	denom := pn.Dot(rd)
	if denom == 0 {
		return Position{}, V3{}
	}
	t := (pn.Dot(po.Sub(ro).V3())) / denom
	if t <= 0 {
		return Position{}, V3{}
	}
	// If hitting underside; flip plane normal
	if denom > 0 {
		return ro.Addf(rd.Muls(t)), pn.Muls(-1)
	}
	return ro.Addf(rd.Muls(t)), pn
}

func (p Plane) Side(point Position) float64 {
	return p.Normal.Dot(point.Sub(p.Origin).V3())
}
