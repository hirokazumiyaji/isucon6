all: isuda isutar

deps:
	glide install

isuda: deps
	GOOS=linux GOARCH=amd64 go build -o isuda isuda.go type.go util.go

isutar: deps
	GOOS=linux GOARCH=amd64 go build -o isutar isutar.go type.go util.go

.PHONY: all deps
