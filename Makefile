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
	@if [ -n $(sudo docker ps -aq --filter "name=yml2openmetrics") ]; then\
		sudo docker rm -f yml2openmetrics;\
	fi
	sudo docker run --name yml2openmetrics -p 8080:8080 yml2openmetrics:v1
