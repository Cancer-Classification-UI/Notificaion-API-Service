#!/bin/bash
cd "$(dirname "$0")"
SCRIPT_DIR="$(pwd)"
echo "Starting container..."
echo "Running image..."
docker run -d $(
if [ -f "$SCRIPT_DIR/../.env" ]; then
    grep -oP 'APP_PORT=\K\d+' $SCRIPT_DIR/../.env | awk '{ print "-p "$1":"$1 }'
else
    echo "-p 8087:8087"
fi
) --name notification-api ccu-notification-api
echo "Container started"