GOLIST?=$$(go list -buildvcs=false ./... | grep -vE "(vendor)")

syntax:
	@echo "==> Checking for linting errors..."
	@golangci-lint run
	@echo "OK"

test:
	@echo "==> Running unit test ..."
	@go test -buildvcs=false -timeout 2s $(GOLIST) -cover
	@echo "OK"

