.PHONY: all
build-lambda:
	docker container run --rm -v $(PWD)/terraform/src/modules/lambda:/lambda -it -w /lambda python:3.7-buster /bin/bash build_lambda.sh
deploy-bff:
	docker-compose up -d deploy
	docker-compose exec deploy sh ./deploy.sh
	docker-compose stop deploy
deploy-front:
	cd app; docker-compose up -d front
	cd app; docker-compose exec front yarn build
	cd app; docker-compose stop front
	docker-compose up -d deploy
	docker-compose exec deploy sh ./deploy-front.sh
	docker-compose stop deploy
