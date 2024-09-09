.PHONY: build

AWS_REGION ?= us-east-1
S3_BUCKET ?= prueba-stori-$(AWS_REGION)

build:
	sam build

start: build
	sam local start-api -p 5000

deploy: build
	sam deploy --stack-name prueba-stori \
		--region $(AWS_REGION) \
		--s3-bucket $(S3_BUCKET) \
		--capabilities CAPABILITY_IAM \
		--no-fail-on-empty-changeset \
		--no-confirm-changeset