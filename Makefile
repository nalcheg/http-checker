run: ## Run application attached to docker-compose.
	docker-compose up --build
run-detached: ## Run application detached from docker-compose.
	docker-compose up --build -d
test: ## Run application tests.
	go test ./application/...

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help
