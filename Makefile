name = hixnav
version = v1.0.0


build-linux:
	SET CGO_ENABLED=0
	SET GOOS=linux
	SET GOARCH=amd64
	go build -o release/$(name)_$(version)-linux main.go
	
build-mac:
	SET CGO_ENABLED=0
	SET GOOS=darwin
	SET GOARCH=amd64
	go build -o release/$(name)_$(version)-mac main.go

build-win:
	SET CGO_ENABLED=0
	SET GOOS=windows
	SET GOARCH=amd64
	go build -o release/$(name)_$(version)-winx64.exe main.go

build: build-linux build-mac build-win

run:
	go run main.go