#!/bin/bash
echo "Killing and removing container..."
docker kill notification-api
docker rm notification-api
echo "Container stopped"
