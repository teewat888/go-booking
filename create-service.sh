#!/bin/bash

if [ $# -ne 1 ]; then 
    echo "Usage: $0 <new-module-name"
    echo "Example: $0 github.com/myorg/user-service"
    exit 1
fi 

NEW_MODULE_NAME=$1
OLD_MODULE_NAME=github.com/teewat888/go-booking/boilerplate

SERVICE_NAME=$(basename $NEW_MODULE_NAME)

if [ -d "$SERVICE_NAME" ]; then 
    echo "Error: Directory '$SERVICE_NAME' already exists"
    exit 1
fi

echo "Creating new service dir: $SERVICE_NAME"
cp -r boilerplate $SERVICE_NAME

cd $SERVICE_NAME

if [[ "$OSTYPE" == "darwin"* ]]; then 
    sed -i '' "s|module $OLD_MODULE_NAME|module $NEW_MODULE_NAME|g" go.mod

    find . -type f -name "*.go" -exec sed -i '' "s|$OLD_MODULE_NAME|$NEW_MODULE_NAME|g" {} \;

    sed -i '' "s|msgo-boilerplate|$SERVICE_NAME|g" makefile
    sed -i '' "s|msgo-Boilerplate|$SERVICE_NAME|g" internal/config/config.go
else
    sed -i "s|module $OLD_MODULE_NAME|module $NEW_MODULE_NAME|g" go.mod
    find . -type f -name "*.go" -exec sed -i "s|$OLD_MODULE_NAME|$NEW_MODULE_NAME|g" {} \;
    sed -i "s|msgo-boilerplate|$SERVICE_NAME|g" makefile
    sed -i "s|msgo-Boilerplate|$SERVICE_NAME|g" internal/config/config.go
fi

rm -f go.sum
go mod tidy

echo "New service created successfully: $SERVICE_NAME"