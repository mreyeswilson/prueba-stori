.PHONY: build

AWS_REGION ?= us-east-1
S3_BUCKET ?= prueba-stori-$(AWS_REGION)
TEST_TO_FILE ?= go test -coverprofile=coverage.out
MOCK_TO_FILE = mockery --dir
COVERAGE_THRESHOLD = 100.0

build:
	sam build

start: build
	sam local start-api -p 5000

coverage:
	$(TEST_TO_FILE) ./internal/application/... ./internal/infraestructure/repositories/...
	go tool cover -html=coverage.out

check-coverage:
	$(TEST_TO_FILE) ./internal/application/... ./internal/infraestructure/repositories/...
	@COVERAGE=$(shell go tool cover -func=coverage.out | grep total: | awk '{print $$3}' | sed 's/%//'); \
	if [ "$$COVERAGE" != "$(COVERAGE_THRESHOLD)" ]; then \
		echo "Coverage is $$COVERAGE%, which is less than the required $(COVERAGE_THRESHOLD)%"; \
		exit 1; \
	else \
		echo "Coverage is $$COVERAGE%, meeting the requirement"; \
	fi

gen-mocks:
	$(MOCK_TO_FILE) ./internal/domain/interfaces --output=mocks --all

gofmt:
	find . -name \*.go -exec gofmt -s -w {} \;

deploy: gofmt build
	sam deploy --stack-name prueba-stori \
		--region $(AWS_REGION) \
		--s3-bucket $(S3_BUCKET) \
		--capabilities CAPABILITY_IAM \
		--no-fail-on-empty-changeset \
		--no-confirm-changeset