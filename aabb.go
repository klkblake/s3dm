package s3dm

type AABB struct {
	Min *V3
	Max *V3
	// Avoid allocations in IntersectsPlane
	temp *V3
}

func NewAABB(min *V3, max *V3) *AABB {
	return &AABB{min, max, NewV3(0, 0, 0)}
}

func (aabb *AABB) IntersectsPlane(plane *Plane) int {
	min := aabb.Min
	max := aabb.Max
	temp := aabb.temp
	temp.Set(min)
	res := plane.Side(temp)
	// Ordered using Gray Code.
	temp.X = max.X
	if plane.Side(temp) != res {
		return 0
	}
	temp.Y = max.Y
	if plane.Side(temp) != res {
		return 0
	}
	temp.X = min.X
	if plane.Side(temp) != res {
		return 0
	}
	temp.Z = max.Z
	if plane.Side(temp) != res {
		return 0
	}
	temp.X = max.X
	if plane.Side(temp) != res {
		return 0
	}
	temp.Y = min.Y
	if plane.Side(temp) != res {
		return 0
	}
	temp.X = min.X
	if plane.Side(temp) != res {
		return 0
	}
	return res
}
