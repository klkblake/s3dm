include $(GOROOT)/src/Make.inc

TARG=s3dm
GOFILES=\
	mat3.go\
	v3.go\
	qtrnn.go\
	transform.go\
	ray.go\
	sphere.go\
	plane.go\

include $(GOROOT)/src/Make.pkg

