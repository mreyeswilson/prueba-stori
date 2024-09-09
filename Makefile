.PHONY: build

AWS_REGION ?= us-east-1
S3_BUCKET ?= prueba-stori-$(AWS_REGION)
TEST_TO_FILE ?= go test -coverprofile=coverage.out
MOCK_TO_FILE = mockery --dir

build:
	sam build

start: build
	sam local start-api -p 5000

coverage:
	$(TEST_TO_FILE) ./internal/application/... ./internal/infraestructure/repositories/...
	go tool cover -html=coverage.out

gen-mocks:
	$(MOCK_TO_FILE) ./internal/infraestructure/aws --all --output=mocks/aws
	$(MOCK_TO_FILE) ./internal/infraestructure/aws --all --output=mocks/aws

gofmt:
	find . -name \*.go -exec gofmt -s -w {} \;

deploy: gofmt build
	sam deploy --stack-name prueba-stori \
		--region $(AWS_REGION) \
		--s3-bucket $(S3_BUCKET) \
		--capabilities CAPABILITY_IAM \
		--no-fail-on-empty-changeset \
		--no-confirm-changeset