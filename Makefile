M = $(shell printf "\033[34;1m▶\033[0m")

install-hooks: hooks-check-deps
	@husky init
	@husky add pre-push ""
	yes | cp -r githooks/. .husky/hooks

hooks-check-deps:
	@if [ -z `which husky` ]; then \
		echo "installing husky";\
		go install github.com/automation-co/husky@latest;\
	fi

lint: lint-check-deps
	$(info $(M) running linter...)
	@golangci-lint run ./...

lint-check-deps:
	@if [ -z `which golangci-lint` ]; then \
		echo "installing golangci-lint";\
		brew install golangci-lint;\
	fi

format:
	$(info $(M) running gofmt...)
	@go fmt ./...

test:
	@go test -v -race -coverpkg=./... -coverprofile=cov.out -outputdir=./ ./...

coverage:test
	@go tool cover -html=./cov.out -o ./coverage.html
