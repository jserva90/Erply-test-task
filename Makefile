run:
	go run ./cmd

lint:
	golangci-lint run --max-same-issues=0 --max-issues-per-linter=0

swag:
	@command -v swag >/dev/null 2>&1 || { \
		echo "swag not found, installing..."; \
		go get -u github.com/swaggo/swag/cmd/swag; \
	}
	swag init -g ./cmd/main.go

