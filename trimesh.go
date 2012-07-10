package s3dm

type TriMesh struct {
	Xform
	tris []staticTri
}

func NewTriMesh(tris []*Tri) *TriMesh {
	tm := new(TriMesh)
	tm.Xform = XformIdentity

	for _, t := range tris {
		p1, p2, p3 := t.Points()
		tm.tris = append(tm.tris, staticTri{p1, p2, p3})
	}
	return tm
}

// TODO: Prettify...
func (tm *TriMesh) Intersect(r *Ray) (Position, V3) {
	first := float64(-1)
	var fi Position
	var fn V3

	rot := tm.Rotation.Matrix()
	for _, t := range tm.tris {
		p1 := tm.Position.Add(rot.Mulv(t.p1.Sub(Position{})))
		p2 := tm.Position.Add(rot.Mulv(t.p2.Sub(Position{})))
		p3 := tm.Position.Add(rot.Mulv(t.p3.Sub(Position{})))
		tt := staticTri{p1, p2, p3}
		i, n := intersectTriangle(&tt, r)
		if (n.X != 0 || n.Y != 0 || n.Z != 0) && (first == -1 || i.Sub(r.Origin).Length() < first) {
			first = i.Sub(r.Origin).Length()
			fi = i
			fn = n
		}
	}
	return fi, fn
}
