.PHONY: help
help: ## ヘルプ
	@grep -E '^[%/a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-22s\033[0m %s\n", $$1, $$2}'
	@echo ''

.PHONY: gqlgen
gqlgen: ## GraphQLのコード生成
	docker-compose exec api go run github.com/99designs/gqlgen generate
