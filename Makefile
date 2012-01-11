include $(GOROOT)/src/Make.inc

TARG=s3dm
GOFILES=\
	mat3.go\
	v3.go\
	qtrnn.go\
	xform.go\
	ray.go\
	sphere.go\
	plane.go\
	tri.go\
	trimesh.go\
	frustum.go\
	aabb.go\

include $(GOROOT)/src/Make.pkg

