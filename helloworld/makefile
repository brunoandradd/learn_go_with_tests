.DEFAULT_GOAL := build

.PHONY:fmt vet build test

fmt:
	 go fmt ./...

vet:
	 go vet ./...

build: fmt vet
	 go build ./...

clean:
	rm cap1

w:
	watchexec -c clear -o do-nothing -d 100ms 'go run hello.go'

test:
	watchexec -c clear -o do-nothing -d 100ms 'go test'
