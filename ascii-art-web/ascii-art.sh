#!/bin/bash

RED="\033[31m"
GREEN="\033[32m"
NC="\033[0m"

IMAGE_NAME="my_ascii_image"
CONTAINER_NAME="my_ascii_container"

buildImage() {
    echo "Building Docker image..."
    docker build -t $IMAGE_NAME .
    if [ $? -ne 0 ]; then
        echo -e "${RED}Error: Failed to build the Docker image.${NC}"
        exit 1
    fi
    echo -e "${GREEN}Docker image built successfully.${NC}"
}

runContainer() {
    echo "Running Docker container..."
    docker run -d -p 8080:8081 --name $CONTAINER_NAME $IMAGE_NAME
    if [ $? -ne 0 ]; then
        echo -e "${RED}Error: Failed to run the Docker container.${NC}"
        exit 1
    fi
    echo -e "${GREEN}Docker container is running.${NC}"
}

if [ ! -f Dockerfile ]; then
    echo -e "${RED}Error: Dockerfile not found.${NC}"
    exit 1
fi

buildImage

runContainer

echo -e "${GREEN}Build and run process completed successfully.${NC}"
