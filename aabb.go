package s3dm

type AABB struct {
	Min V3
	Max V3
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

func (aabb AABB) IntersectsPlane(plane Plane) int {
	box := [2]V3{aabb.Min, aabb.Max}
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
	if plane.Side(V3{box[px].X, box[py].Y, box[pz].Z}) < 0 {
		return -1
	}
	if plane.Side(V3{box[1-px].X, box[1-py].Y, box[1-pz].Z}) > 0 {
		return 1
	}
	return 0
}
