# A Self-Documenting Makefile: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
.PHONY: help bins
.DEFAULT_GOAL := help

OS = $(shell uname | tr A-Z a-z)

# general build-product folder, cleaned as part of `make clean`
BUILD := .build
BIN := bin

AWS_SDK_GENERATOR_VERSION = 0.0.3

# Gather all templates
ALL_TEMPLATES :=  $(shell find templates -name "*.tmpl")

$(BUILD):
	@mkdir -p $(BUILD)

$(BIN):
	@mkdir -p $(BIN)

$(BIN)/temporal-aws-sdk-generator: $(BIN)/temporal-aws-sdk-generator-${AWS_SDK_GENERATOR_VERSION}
	@ln -sf temporal-aws-sdk-generator-${AWS_SDK_GENERATOR_VERSION} $(BIN)/temporal-aws-sdk-generator
$(BIN)/temporal-aws-sdk-generator-${AWS_SDK_GENERATOR_VERSION}: $(BIN)
	curl -L https://github.com/temporalio/temporal-aws-sdk-generator/releases/download/v${AWS_SDK_GENERATOR_VERSION}/temporal-aws-sdk-generator_${OS}_amd64.tar.gz | tar -zOxf - temporal-aws-sdk-generator > $(BIN)/temporal-aws-sdk-generator-${AWS_SDK_GENERATOR_VERSION} && chmod +x ./bin/temporal-aws-sdk-generator-${AWS_SDK_GENERATOR_VERSION}

$(BUILD)/generate: $(BUILD) $(ALL_TEMPLATES) $(BIN)/temporal-aws-sdk-generator ## generate code based on templates in templates directory
	$(BIN)/temporal-aws-sdk-generator --template-dir templates --output-dir .
	touch $(BUILD)/generate

$(BUILD)/clients: $(BUILD)/generate
	go build ./clients/...
	touch $(BUILD)/clients

$(BIN)/aws-sdk-worker: $(BUILD)/generate
	go build -o $@ aws-sdk-worker.go

$(BIN)/samples/ec2demo-starter: $(BUILD)/generate
	go build -o $@ samples/ec2demo/starter/main.go

$(BIN)/samples/ec2demo-worker: $(BUILD)/generate
	go build -o $@ samples/ec2demo/worker/main.go

$(BIN)/samples/s3list: $(BUILD)/generate
	go build -o $@ samples/s3list/worker/main.go

generate: $(BUILD)/generate ## Regenerate code if templates changed

bins: $(BUILD)/clients $(BIN)/aws-sdk-worker $(BIN)/samples/ec2demo-worker $(BIN)/samples/ec2demo-starter $(BIN)/samples/s3list ## Build binaries

clean: ## Remove .build directory. Doesn't revert generated code changes.
	rm -rf $(BUILD)
	rm -rf $(BIN)

help:
	@grep -h -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
