package s3dm

type AABB struct {
	Min V3
	Max V3
}

func (aabb AABB) MoveGlobal(v V3) AABB {
	aabb.Min.AddLocal(v)
	aabb.Max.AddLocal(v)
	return aabb
}

func (aabb AABB) IntersectsPlane(plane *Plane) int {
	min := aabb.Min
	max := aabb.Max
	var temp V3
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
