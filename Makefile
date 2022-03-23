build:
	GOARCH=amd64 GOOS=linux go build -o hixnav main.go
	
build-win:
	go build -o hixnav.exe main.go

run:
	go run main.go