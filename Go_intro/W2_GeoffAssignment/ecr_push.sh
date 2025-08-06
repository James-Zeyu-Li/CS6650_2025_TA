#!/bin/bash

# 设置变量（你可以根据自己的实际值修改）
AWS_ACCOUNT_ID=589535382240
REGION=us-west-2
REPO_NAME=neu/week2
IMAGE_NAME=go-echo
TAG=Geoff

# Build image
echo "Building Docker image..."
docker build -t $IMAGE_NAME .

# Tag docker image
echo "Tagging image for ECR..."
docker tag $IMAGE_NAME:latest $AWS_ACCOUNT_ID.dkr.ecr.$REGION.amazonaws.com/$REPO_NAME:$TAG

# 登录 ECR
echo "Logging into ECR"
aws ecr get-login-password --region $REGION | \
docker login --username AWS --password-stdin $AWS_ACCOUNT_ID.dkr.ecr.$REGION.amazonaws.com

echo "Pushing image to ECR"
docker push $AWS_ACCOUNT_ID.dkr.ecr.$REGION.amazonaws.com/$REPO_NAME:$TAG

echo "Done! Your image is available at:"
echo "AWS_ACCOUNT_ID.dkr.ecr.$REGION.amazonaws.com/$REPO_NAME:$TAG"
