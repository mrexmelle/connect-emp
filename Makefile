
PROJECT_NAME=connect-emp
VERSION=0.1.0
IMAGE_NAME=ghcr.io/mrexmelle/$(PROJECT_NAME)
GO_SOURCES=$(shell find . -name '*.go' -not -path "./vendor/*")

$(PROJECT_NAME): $(GO_SOURCES)
	go build -o $@ ./cmd/main.go

clean:
	rm -rf $(PROJECT_NAME)

distclean:
	rm -rf $(PROJECT_NAME) docs

docker-build:
	docker build -t $(IMAGE_NAME):$(VERSION) .

docker-push:
	docker push $(IMAGE_NAME):$(VERSION)

docs:
	swag init --parseDependency -g cmd/main.go

install-deps:
	GOPRIVATE=github.com/mrexmelle/connect-org go mod tidy

test:
	go test ./internal/...

all: $(PROJECT_NAME)