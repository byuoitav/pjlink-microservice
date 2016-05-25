#!/usr/bin/env bash
# Pasted into Jenkins to build (will eventually be fleshed out to work with a Docker Hub and Amazon AWS)

echo "Stopping running application"
docker stop pjlink-service
docker rm pjlink-service

echo "Building container"
docker build -t byuoitav/pjlink-service .

echo "Starting the new version"
docker run -d --restart=always --name pjlink-service -p 8005:8005 byuoitav/pjlink-service:latest
