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
	return aabb.Min.X <= other.Max.X && aabb.Max.X >= other.Min.X &&
		aabb.Min.Z <= other.Max.Z && aabb.Max.Z >= other.Min.Z &&
		aabb.Min.Y <= other.Max.Y && aabb.Max.Y >= other.Min.Y
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
