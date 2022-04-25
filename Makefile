name = hixnav

build-front:
	cd web & npm run build

front: build-front
	go-bindata -o cmd/bindata.go -pkg cmd ./web/dist/...


build-linux:
	set CGO_ENABLED=0
	set GOOS=linux
	set GOARCH=amd64
	go build -o release/$(name)-linux-amd64 main.go
	
build-mac:
	set CGO_ENABLED=0
	set GOOS=darwin
	set GOARCH=amd64
	go build -o release/$(name)-mac-amd64 main.go

build-win:
	set CGO_ENABLED=0
	set GOOS=windows
	set GOARCH=amd64
	go build -o release/$(name)-winx64.exe main.go

build: build-linux build-mac build-win

build-web:
	cd web && npm run build

serve:
	cd web && npm run serve

run:
	go run main.go

push:
	git push origin master
	git push coding master