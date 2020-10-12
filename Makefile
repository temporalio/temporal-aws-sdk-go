# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help bins
.DEFAULT_GOAL := help

# general build-product folder, cleaned as part of `make clean`
BUILD := .build
BIN := bin

# Gather all templates
ALL_TEMPLATES :=  $(shell find templates -name "*.tmpl")

$(BUILD):
	@mkdir -p $(BUILD)

$(BUILD)/generate: $(BUILD) $(ALL_TEMPLATES) ## generate code based on templates in templates directory
	go get go.temporal.io/aws-sdk-generator
	aws-sdk-generator --template-dir templates --output-dir .
	touch $(BUILD)/generate

$(BUILD)/clients: $(BUILD)/generate
	go build ./clients/...
	touch $(BUILD)/clients

$(BIN)/awsactivities: $(BUILD)/generate
	go build -o $@ awsactivities.go

$(BIN)/samples/ec2demo-starter: $(BUILD)/generate
	go build -o $@ samples/ec2demo/starter/main.go

$(BIN)/samples/ec2demo-worker: $(BUILD)/generate
	go build -o $@ samples/ec2demo/worker/main.go

$(BIN)/samples/s3list: $(BUILD)/generate
	go build -o $@ samples/s3list/worker/main.go

generate: $(BUILD)/generate ## Regenerate code if templates changed

bins: $(BUILD)/clients $(BIN)/awsactivities $(BIN)/samples/ec2demo-worker $(BIN)/samples/ec2demo-starter $(BIN)/samples/s3list ## Build binaries

clean: ## Remove .build directory. Doesn't revert generated code changes.
	rm -rf $(BUILD)
	rm -rf $(BIN)

help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
