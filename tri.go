package s3dm

// The triangle interface is used by both Tri and TriMesh
type triangle interface {
	Points() (Position, Position, Position)
	Normal() V3
	Center() Position
}

// Checks if p1 and p2 are on same side of line from a to b
func sameSide(p1, p2, a, b Position) bool {
	bsuba := b.Sub(a)
	cp1 := bsuba.Cross(p1.Sub(a))
	cp2 := bsuba.Cross(p2.Sub(a))
	return cp1.Dot(cp2) >= 0
}

func pointInside(t triangle, p Position) bool {
	p1, p2, p3 := t.Points()
	return sameSide(p, p1, p2, p3) &&
		sameSide(p, p2, p1, p3) &&
		sameSide(p, p3, p1, p2)
}

func intersectTriangle(t triangle, r *Ray) (Position, V3) {
	p := Plane{t.Center(), t.Normal()}
	i, n := p.Intersect(r)
	if (n.X != 0 || n.Y != 0 || n.Z != 0) && pointInside(t, i) {
		return i, t.Normal()
	}
	return Position{}, V3{}
}

// StaticTri doesn't have a transformation so it can be used with TriMesh
type staticTri struct {
	p1, p2, p3 Position
}

func (t *staticTri) Points() (Position, Position, Position) {
	return t.p1, t.p2, t.p3
}

func (t *staticTri) Normal() V3 {
	// Cross( p2 - p1, p3 - p1 )
	return t.p2.Sub(t.p1).Cross(t.p3.Sub(t.p1)).Unit()
}

func (t *staticTri) Center() Position {
	// (p1 + p2 + p3) / 3
	a := t.p2.Sub(t.p1)
	b := t.p3.Sub(t.p1)
	return t.p1.Add(a.Add(b).Muls(1. / 3))
}

// Tri is basicly a staticTri with a transform
type Tri struct {
	Xform
	p1, p2, p3 V3
}

func NewTri(p1, p2, p3 V3) *Tri {
	t := new(Tri)
	t.Xform = XformIdentity
	t.p1, t.p2, t.p3 = p1, p2, p3
	return t
}

func (t *Tri) Copy() *Tri {
	return &Tri{t.Xform, t.p1, t.p2, t.p3}
}

func (t *Tri) Points() (Position, Position, Position) {
	rot := t.Rotation.Matrix()
	return t.Position.Add(rot.Mulv(t.p1)), t.Position.Add(rot.Mulv(t.p2)), t.Position.Add(rot.Mulv(t.p3))
}

func (t *Tri) SetPoints(p1, p2, p3 V3) {
	t.p1 = p1
	t.p2 = p2
	t.p3 = p3
}

func (t *Tri) Normal() V3 {
	p1 := t.Position.Add(t.p1.Rotate(t.Rotation))
	p2 := t.Position.Add(t.p2.Rotate(t.Rotation))
	p3 := t.Position.Add(t.p3.Rotate(t.Rotation))
	st := staticTri{p1, p2, p3}
	return st.Normal()
}

func (t *Tri) Center() Position {
	p1 := t.Position.Add(t.p1.Rotate(t.Rotation))
	p2 := t.Position.Add(t.p2.Rotate(t.Rotation))
	p3 := t.Position.Add(t.p3.Rotate(t.Rotation))
	st := staticTri{p1, p2, p3}
	return st.Center()
}

func (t *Tri) Intersect(r *Ray) (Position, V3) {
	return intersectTriangle(t, r)
}
