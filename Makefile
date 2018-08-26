default: test

test:
	@go test -cover ./...

.PHONY: test
