# This is how we want to name the binary output
OUTPUT_NAME=codis_status_exporter
DIST_DIR="dist"

VERSION="v2.0.0"
# These are the values we want to pass for Version and BuildTime
GITTAG=2021.05.26.release
#GITTAG=`git describe --abbrev=0`
BUILD_TIME=`date +%FT%T%z`

linux:
	echo  "Compiling......."
	mkdir -p ${DIST_DIR}
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build  -o ${DIST_DIR}/${OUTPUT_NAME}.linux.amd64 main.go

all:
	go build ${LDFLAGS} -o ${OUTPUT_NAME} main.go
