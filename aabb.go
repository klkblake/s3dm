package s3dm

type AABB struct {
	Min Position
	Max Position
}

func (aabb AABB) MoveGlobal(v V3) AABB {
	aabb.Min = aabb.Min.Add(v)
	aabb.Max = aabb.Max.Add(v)
	return aabb
}

func (aabb AABB) Intersects(other AABB) bool {
	min1 := aabb.Min
	max1 := aabb.Max
	min2 := other.Min
	max2 := other.Max
	return min1.X.Lt(max2.X) && max1.X.Gt(min2.X) &&
		min1.Z.Lt(max2.Z) && max1.Z.Gt(min2.Z) &&
		min1.Y.Lt(max2.Y) && max1.Y.Gt(min2.Y)
}

func (aabb AABB) IntersectsPlane(plane Plane) float64 {
	box := [2]Position{aabb.Min, aabb.Max}
	var px, py, pz int
	if plane.Normal.X > 0 {
		px = 1
	}
	if plane.Normal.Y > 0 {
		py = 1
	}
	if plane.Normal.Z > 0 {
		pz = 1
	}
	d := plane.Side(Position{box[px].X, box[py].Y, box[pz].Z})
	if d < 0 {
		return d
	}
	d = plane.Side(Position{box[1-px].X, box[1-py].Y, box[1-pz].Z})
	if d > 0 {
		return d
	}
	return 0
}

func (aabb AABB) IntersectsFrustum(frustum *Frustum) float64 {
	// TODO: exploit temporal coherence.
	res := float64(1)
	for _, plane := range frustum.Planes {
		intersects := aabb.IntersectsPlane(plane)
		if intersects < 0 {
			return intersects
		}
		if intersects == 0 {
			res = 0
		}
	}
	return res
}
