# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html

.PHONY: generate
generate:
	mkdir -p build
	cd cmd/temporal-aws-sdk-gen/ && go build -o ../../build/temporal-aws-sdk-gen
	build/temporal-aws-sdk-gen

.PHONY: help
.DEFAULT_GOAL := help
help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
