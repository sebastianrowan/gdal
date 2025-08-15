// +build windows,amd64

package gdal

/*
#cgo windows LDFLAGS: -Lc:/OSGeo4W/lib -lgdal_i
#cgo windows CFLAGS: -IC:/OSGeo4W/include
*/
import "C"
