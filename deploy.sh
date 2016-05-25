#!/usr/bin/env bash

echo "Stopping running application"
ssh $DEPLOY_USERNAME@$DEPLOY_HOST 'docker stop pjlink-microservice'
ssh $DEPLOY_USERNAME@$DEPLOY_HOST 'docker rm pjlink-microservice'

echo "Pulling latest version"
ssh $DEPLOY_USERNAME@$DEPLOY_HOST 'docker pull byuoitav/pjlink-microservice:latest'

echo "Starting the new version"
ssh $DEPLOY_USERNAME@$DEPLOY_HOST 'docker run -d --restart=always --name pjlink-microservice -p 8005:8005 byuoitav/pjlink-microservice:latest'

echo "Success!"

exit 0
