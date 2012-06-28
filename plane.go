package s3dm

type Plane struct {
	Origin V3
	Normal V3
}

func (p Plane) Intersect(r *Ray) (V3, V3) {
	po, pn := p.Origin, p.Normal
	ro, rd := r.Origin, r.Dir

	denom := pn.Dot(rd)
	if denom == 0 {
		return V3{}, V3{}
	}
	c := pn.Dot(po)
	t := (c - pn.Dot(ro)) / denom
	if t <= 0 {
		return V3{}, V3{}
	}
	// If hitting underside; flip plane normal
	if denom > 0 {
		return ro.Add(rd.Muls(t)), pn.Muls(-1)
	}
	return ro.Add(rd.Muls(t)), pn
}

func (p Plane) Side(point V3) float64 {
	return p.Normal.Dot(point.Sub(p.Origin))
}
