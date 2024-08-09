#!/bin/bash

# Set image and container name
IMAGE_NAME="ascii-art-web"
CONTAINER_NAME="ascii-art-web-container"

# Build image
docker build -t $IMAGE_NAME

# Run container
docker run -p 9000:9000 --name $CONTAINER_NAME $IMAGE_NAME