GOPATH := $(HOME)/go
PATH := $(PATH):$(GOPATH)/bin

build:
	GCO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o twampc cmd/twampc/main.go
	GCO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o twampd cmd/twampd/main.go
