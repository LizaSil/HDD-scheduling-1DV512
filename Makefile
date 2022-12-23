build:
	GO111MODULE="on" && go build -o

run:
	GO111MODULE="on" && go run *.go

compile:
	GOARCH=amd64 GOOS=darwin go build -o darwin64
	GOARCH=arm64 GOOS=darwin go build -o darwinARM64
	GOARCH=amd64 GOOS=linux go build -o linux64
	GOARCH=amd64 GOOS=windows go build -o windows64