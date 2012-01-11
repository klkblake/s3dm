package s3dm

type AABB struct {
	Min *V3
	Max *V3
}

func NewAABB(min *V3, max *V3) *AABB {
	return &AABB{min, max}
}

func (aabb *AABB) Verticies() []*V3 {
	verticies := make([]*V3, 8)
	verticies[0] = aabb.Min
	verticies[1] = aabb.Max
	verticies[2] = NewV3(aabb.Min.X, aabb.Min.Y, aabb.Max.Z)
	verticies[3] = NewV3(aabb.Max.X, aabb.Max.Y, aabb.Min.Z)
	verticies[4] = NewV3(aabb.Min.X, aabb.Max.Y, aabb.Min.Z)
	verticies[5] = NewV3(aabb.Max.X, aabb.Min.Y, aabb.Max.Z)
	verticies[6] = NewV3(aabb.Min.X, aabb.Max.Y, aabb.Max.Z)
	verticies[7] = NewV3(aabb.Max.X, aabb.Min.Y, aabb.Min.Z)
	return verticies
}

func (aabb *AABB) IntersectsPlane(plane *Plane) int {
	verticies := aabb.Verticies()
	res := plane.Side(verticies[0])
	for i := 1; i < len(verticies); i++ {
		if plane.Side(verticies[i]) != res {
			return 0
		}
	}
	return res
}
