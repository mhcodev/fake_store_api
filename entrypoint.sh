#!/bin/bash

echo "Starting the application in $ENVIRONMENT mode..."

if [ "$ENVIRONMENT" = "development" ]; then
  echo "Running air..."
  exec air
elif [ "$ENVIRONMENT" = "production" ]; then
  echo "Running go run main.go..."
  exec ./tmp/main
else
  echo "Unknown environment: $ENVIRONMENT. Exiting."
  exit 1
fi
