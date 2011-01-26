package s3dm

type Ray struct {
	O *V3
	D *V3
}

func NewRay(o, d *V3) *Ray {
	r := new(Ray)
	r.O = o
	r.D = d
	return r
}
