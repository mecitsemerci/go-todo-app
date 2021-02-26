test:
	@echo "Test started"
	@docker-compose up -d
	@sleep 1 && \
		MONGO_URL="mongodb://`docker-compose port mongo 27017`/" \
		MONGO_TODO_DB=TodoDb \
		JAEGER_DISABLED=true \
		go test -cover ./internal/... -covermode=atomic
	@docker-compose down

lint:
	golangci-lint run --fix -v


.PHONY:test lint