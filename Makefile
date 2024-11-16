APP = app

BRANCH = $(shell git branch --show-current)
HASH = $(shell git show --pretty=format:%h --no-patch)
TAG = $(shell git for-each-ref refs/tags --sort=-taggerdate --format='%(refname:short)' --count=1)
GIT_INFO = $(BRANCH)-$(HASH)
COMMIT = $(HASH)

.PHONY: info
info:
	@echo '$(GIT_INFO)'
	@echo '$(BRANCH)'
	@echo '$(COMMIT)'

.PHONY: build
build: info
	@go build -v -a -ldflags="-X 'main.GitInfo=$(GIT_INFO)' -X 'main.Tag=$(TAG)' -X 'main.Commit=$(COMMIT)' -X 'main.Branch=$(BRANCH)'" -o bin/acloset cmd/acloset/main.go

.PHONY: swag
swag:
	@go install github.com/swaggo/swag/cmd/swag@v1.16.4
	@swag init -g cmd/app/main.go

.PHONY: package
package:
	go mod tidy

.PHONY: test
test:
	@go test -coverpkg=./... -coverprofile=coverage.out ./...
	@cat coverage.out | grep -v 'github.com/omnia-core/go-echo-template/pkg\|cmd\|log' > exclude.out
	@go tool cover -html=exclude.out

.PHONY: run
run:
	@go run ./cmd/$(APP)/main.go
