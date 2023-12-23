.PHONY: all
all: vet test build

.PHONY: build
build:
	go build ./cmd/tm

.PHONY: vet
vet:
	go vet -structtag=false ./...

.PHONY: test
test:
	go test -v -count=1 ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: demo
# see https://github.com/charmbracelet/vhs
demo:
	vhs demo.tape
