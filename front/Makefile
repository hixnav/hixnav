name = hixnav

serve:
	npm run serve

build:
	npm run build

gobuild: build
	go-bindata -o pkg/bindata.go -pkg pkg ./dist/...
