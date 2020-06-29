.PHONY: all
build-lambda:
	docker container run --rm -v $(PWD)/terraform/src/modules/lambda:/lambda -it -w /lambda python:3.7-buster /bin/bash build_lambda.sh
