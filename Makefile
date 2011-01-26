include $(GOROOT)/src/Make.inc

TARG=s3dm
GOFILES=\
	mat4.go\
	v3.go\
	ray.go\
	sphere.go\

include $(GOROOT)/src/Make.pkg

