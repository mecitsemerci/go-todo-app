unit_test:
	go test -cover ./internal/... -covermode=atomic

.PHONY: unit_test

integration_test:
	@docker-compose up -d
	@sleep 1 && \
		MONGO_URL="mongodb://127.0.0.1:27017/" \
		MONGO_TODO_DB=TodoDbTest \
        MONGO_CONNECTION_TIMEOUT=20 \
        MONGO_MAX_POOL_SIZE=10 \
		JAEGER_DISABLED=true \
		go test -coverpkg ./... ./test/...
	@docker-compose down

.PHONY: integration_test

lint:
	golangci-lint run --fix -v


.PHONY: lint