include $(GOROOT)/src/Make.inc

TARG=s3dm
GOFILES=\
	mat4.go\
	v3.go\
	qtrnn.go\
	ray.go\
	sphere.go\
	plane.go\

include $(GOROOT)/src/Make.pkg

