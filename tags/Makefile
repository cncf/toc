IMAGE_NAME=golang:1.18
export GOPROXY?=https://proxy.golang.org

default: \
	generate \

generate:
	go run ./generator/app.go

generate-dockerized:
	docker run --rm -e WHAT -e GOPROXY -v $(shell pwd):/go/src/app:Z $(IMAGE_NAME) make -C /go/src/app generate

test:
	go test -v ./generator/...

test-dockerized:
	docker run --rm -v $(shell pwd):/go/src/app:Z $(IMAGE_NAME) make -C /go/src/app test

.PHONY: default generate generate-dockerized test test-dockerized
