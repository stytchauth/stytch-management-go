.PHONY: help
help:
	@echo "Available commands"
	@grep -E '^[a-zA-Z_-]+:.*?# .*$$' $(MAKEFILE_LIST) | sort 

# Include .env file if it exists
-include .env

.PHONY: lint
lint:
	golangci-lint run --fix

TEST_CMD := gotestsum ./pkg/... -v
ifeq (, $(shell command -v gotestsum))
	TEST_CMD := go test ./pkg/... -v
endif

.PHONY: test
test:
	@if [ ! -f .env ]; then \
		echo "Error: .env file not found. Please create a .env file with STYTCH_WORKSPACE_KEY_ID, STYTCH_WORKSPACE_KEY_SECRET, and optionally STYTCH_WORKSPACE_BASE_URI"; \
		exit 1; \
	fi
	STYTCH_WORKSPACE_KEY_ID="$(STYTCH_WORKSPACE_KEY_ID)" \
	STYTCH_WORKSPACE_KEY_SECRET="$(STYTCH_WORKSPACE_KEY_SECRET)" \
	STYTCH_WORKSPACE_BASE_URI="$(STYTCH_WORKSPACE_BASE_URI)" \
	$(TEST_CMD)

.PHONY: tests
tests: test # A useful alias
