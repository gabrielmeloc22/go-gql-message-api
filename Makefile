.PHONY: dev

dev:
			@find . -name "*.go" | entr -r go run ./cmd/api/main.go