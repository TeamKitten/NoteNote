all: appbuild bindata build

build:
	@go build -o NoteNote

appbuild:
	@npm run build

bindata:
	@go-bindata-assetfs ./dist/...

## DOES NOT WORK ON APPLE SILICON MAC!!
#dev:
#	@realize start
