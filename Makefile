BINARY_NAME=ytt

build:
	go build -o bin/${BINARY_NAME} cmd/app/main.go

compile:
	echo "Compiling for ubuntu and windows"
	GOOS=linux GOARCH=amd64 go build -o bin/${BINARY_NAME} cmd/app/main.go
	GOOS=windows GOARCH=386 go build -o bin/${BINARY_NAME}.exe cmd/app/main.go