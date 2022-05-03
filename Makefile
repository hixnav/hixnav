name = hixnav

build-web:
	cd web & npm run build

web: build-web
	go-bindata -o cmd/bindata.go -pkg cmd ./web/dist/...

serve:
	cd web && npm run serve

build-linux:
	set CGO_ENABLED=0
	set GOOS=linux
	set GOARCH=amd64
	go build -o bin/$(name)-linux-amd64 main.go
	
build-mac:
	set CGO_ENABLED=0
	set GOOS=darwin
	set GOARCH=amd64
	go build -o bin/$(name)-mac-amd64 main.go

build-win:
	set CGO_ENABLED=0
	set GOOS=windows
	set GOARCH=amd64
	go build -o bin/$(name)-winx64.exe main.go

build: build-linux build-mac build-win

run:
	go run main.go