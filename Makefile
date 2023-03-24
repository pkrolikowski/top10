path = bin
name = top10

build:
	go build -o $(path)/$(name) main.go

compile:
	GOOS=darwin GOARCH=amd64 go build -o  $(path)/$(name)-darwin-amd64 main.go
	GOOS=darwin GOARCH=arm64 go build -o  $(path)/$(name)-darwin-arm64 main.go
	GOOS=freebsd GOARCH=386 go build -o $(path)/$(name)-freebsd-386 main.go
	GOOS=freebsd GOARCH=amd64 go build -o $(path)/$(name)-freebsd-amd64 main.go
	GOOS=linux GOARCH=amd64 go build -o $(path)/$(name)-linux-amd64 main.go
	GOOS=linux GOARCH=arm go build -o $(path)/$(name)-linux-arm main.go
	GOOS=windows GOARCH=amd64 go build -o $(path)/$(name)-windows-amd64 main.go