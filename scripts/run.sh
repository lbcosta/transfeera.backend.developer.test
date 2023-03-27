#!/bin/bash

RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

cleanup() {
    echo -e "${GREEN}Stopping containers...${NC}"
    docker-compose down
    exit 0
}

trap cleanup SIGINT

cd "$(dirname "$0")" || exit

go test ../src/api/v1/handlers
if [ $? -ne 0 ]; then
    echo -e "${RED}Error: Tests failed in 'handlers'. Check the test output above for details.${NC}"
    exit 1
fi

go test ../src/api/v1/services
if [ $? -ne 0 ]; then
    echo -e "${RED}Error: Tests failed in 'services'. Check the test output above for details.${NC}"
    exit 1
fi

docker-compose up --build -d

echo -e "${GREEN}Containers are running. Press CTRL+C to stop them.${NC}"

while true; do
    sleep 1
done