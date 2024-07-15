.PHONY: test

help: # コマンド確認
	@echo "\033[32mAvailable targets:\033[0m"
	@grep "^[a-zA-Z\-]*:" Makefile | grep -v "grep" | sed -e 's/^/make /' | sed -e 's/://'

run:
	docker compose exec app sh -c "go run ./cmd/main.go"


## テスト処理の共通化
define tests
	$(if $(TEST_PATH),\
		$(if $(TEST_TAGS),\
			go test -timeout 10m -tags=$(TEST_TAGS) $(TEST_PATH) $(TEST_OPTIONS),\
			go test -timeout 10m $(TEST_PATH) $(TEST_OPTIONS)\
		),\
		$(if $(TEST_TAGS),\
			go test -timeout 10m -tags=$(TEST_TAGS) ./... $(TEST_OPTIONS),\
			go test -timeout 10m ./... $(TEST_OPTIONS)\
		)\
	)
endef

test: lint
	cd app && go test ./...

# コマンドの例：
# $ make test-domain
# $ make test-domain path=./internal/domain/cart opts="-run TestCart_AddProduct"
test-domain:
	$(eval TEST_PATH=$(or ${path},./internal/domain...))
	$(eval TEST_OPTIONS=${opts})
	$(call tests)

# コマンドの例：
# $ make test-infrastructure
# $ make test-infrastructure path=./internal/infrastructure/mysql/repository... opts="-run TestOrderRepository_Save"
test-infrastructure:
	$(eval TEST_PATH=$(or ${path},./internal/infrastructure...))
	$(eval TEST_OPTIONS=${opts})
	$(call tests)

test-integration: lint test-integration-read test-integration-write

# コマンドの例：
# $ make test-integration-read
# $ make test-integration-read path=./internal/server/api_test/api_read_test opts="-run TestOrder_GetCart"
test-integration-read:
	$(eval TEST_PATH=$(or ${path},./internal/server/api_test/api_read_test...))
	$(eval TEST_TAGS=integration_read)
	$(eval TEST_OPTIONS=${opts})
	$(call tests)

# コマンドの例：
# $ make test-integration-write
# $ make test-integration-write path=./internal/server/api_test/api_write_test opts="-run TestOrder_PostCart"
test-integration-write:
	$(eval TEST_PATH=$(or ${path},./internal/server/api_test/api_write_test...))
	$(eval TEST_TAGS=integration_write)
	$(eval TEST_OPTIONS=${opts})
	$(call tests)
