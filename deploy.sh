#!/usr/bin/env bash
# Pasted into Jenkins to build (will eventually be fleshed out to work with a Docker Hub and Amazon AWS)

echo "Stopping running application"
docker stop pjlink-microservice
docker rm pjlink-microservice

echo "Building container"
docker build -t byuoitav/pjlink-microservice .

echo "Starting the new version"
docker run -d --restart=always --name pjlink-microservice -p 8005:8005 byuoitav/pjlink-microservice:latest
