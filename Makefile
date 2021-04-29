.PHONY:
.SILENT:

init:
	mkdir .bin

build:
	go build -o ./.bin/yml2openmetrics ./cmd/yml2openmetrics/main.go

run: build
	./.bin/yml2openmetrics

delimage:
	sudo docker rm yml2openmetrics

image: build
	sudo docker build -t yml2openmetrics:v1 .

start:
	sudo docker run --name yml2openmetrics -p 8080:8080 yml2openmetrics:v1
