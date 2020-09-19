#!/bin/bash
aws ecr get-login-password | docker login --username AWS --password-stdin https://$AWS_ACCOUNT.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com

docker image build -t bff ./bff/
docker tag bff $AWS_ACCOUNT.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/repo_bff:latest
docker push $AWS_ACCOUNT.dkr.ecr.$AWS_DEFAULT_REGION.amazonaws.com/repo_bff:latest
