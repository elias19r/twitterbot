#!/bin/bash

echo "Compile a static binary..."
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags '-w -s' ./cmd/twitterbot

echo "Source .env file..."
source ./.env

echo "Build docker image..."
docker build -q \
  --build-arg CONSUMER_KEY=$CONSUMER_KEY \
  --build-arg CONSUMER_SECRET=$CONSUMER_SECRET \
  --build-arg ACCESS_TOKEN=$ACCESS_TOKEN \
  --build-arg ACCESS_TOKEN_SECRET=$ACCESS_TOKEN_SECRET \
  -t twitterbot .

echo "Remove static binary..."
rm ./twitterbot

echo "Done."
