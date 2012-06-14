package s3dm

// The triangle interface is used by both Tri and TriMesh
type triangle interface {
	Points() (V3, V3, V3)
	Normal() V3
	Center() V3
}

// Checks if p1 and p2 are on same side of line from a to b
func sameSide(p1, p2, a, b V3) bool {
	bsuba := b.Sub(a)
	cp1 := bsuba.Cross(p1.Sub(a))
	cp2 := bsuba.Cross(p2.Sub(a))
	return cp1.Dot(cp2) >= 0
}

func pointInside(t triangle, p V3) bool {
	p1, p2, p3 := t.Points()
	return sameSide(p, p1, p2, p3) &&
		sameSide(p, p2, p1, p3) &&
		sameSide(p, p3, p1, p2)
}

func intersectTriangle(t triangle, r *Ray) (V3, V3) {
	p := NewPlane(t.Center(), t.Normal())
	i, n := p.Intersect(r)
	if (n.X != 0 || n.Y != 0 || n.Z != 0) && pointInside(t, i) {
		return i, t.Normal()
	}
	return V3{}, V3{}
}

// StaticTri doesn't have a transformation so it can be used with TriMesh
type staticTri struct {
	p1, p2, p3 V3
}

func (t *staticTri) Points() (V3, V3, V3) {
	return t.p1, t.p2, t.p3
}

func (t *staticTri) Normal() V3 {
	// Cross( p2 - p1, p3 - p1 )
	return t.p2.Sub(t.p1).Cross(t.p3.Sub(t.p1)).Unit()
}

func (t *staticTri) Center() V3 {
	// (p1 + p2 + p3) / 3
	return t.p1.Add(t.p2.Add(t.p3)).Divs(3)
}

// Tri is basicly a staticTri with a transform
type Tri struct {
	Xform
	st staticTri
}

func NewTri(p1, p2, p3 V3) *Tri {
	t := new(Tri)
	t.ResetXform()
	t.st = staticTri{p1, p2, p3}
	return t
}

func (t *Tri) Copy() *Tri {
	return NewTri(t.Points())
}

func (t *Tri) Points() (V3, V3, V3) {
	return t.Mulv(t.st.p1), t.Mulv(t.st.p2), t.Mulv(t.st.p3)
}

func (t *Tri) SetPoints(p1, p2, p3 V3) {
	t.st.p1 = p1
	t.st.p2 = p2
	t.st.p3 = p3
}

func (t *Tri) Normal() V3 {
	return t.Mulv(t.st.Normal())
}

func (t *Tri) Center() V3 {
	return t.Mulv(t.st.Center())
}

func (t *Tri) Intersect(r *Ray) (V3, V3) {
	return intersectTriangle(t, r)
}
