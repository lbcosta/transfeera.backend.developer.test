#!/bin/bash

RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
NC='\033[0m' # No Color

cleanup() {
    echo -e "${GREEN}Stopping containers...${NC}"
    docker-compose down
    exit 0
}

trap cleanup SIGINT

cd "$(dirname "$0")" || exit

output=$(go test -v ../src/api/v1/handlers 2>&1)
if [ $? -ne 0 ]; then
    echo -e "${output}" | awk '/FAIL/ {print "'${RED}'" $0 "'${NC}'"}; /PASS/ {print "'${GREEN}'" $0 "'${NC}'"}; /RUN / {print "'${YELLOW}'" $0 "'${NC}'"}; !/FAIL|PASS|RUN / {print}'
    echo -e "${RED}Error: Tests failed in 'handlers'. Check the test output above for details.${NC}"
    exit 1
else
        echo -e "${output}" | awk '/FAIL/ {print "'${RED}'" $0 "'${NC}'"}; /PASS/ {print "'${GREEN}'" $0 "'${NC}'"}; /RUN / {print "'${YELLOW}'" $0 "'${NC}'"}; !/FAIL|PASS|RUN / {print}'

fi

output=$(go test -v ../src/api/v1/services 2>&1)
if [ $? -ne 0 ]; then
    echo -e "${output}" | awk '/FAIL/ {print "'${RED}'" $0 "'${NC}'"}; /PASS/ {print "'${GREEN}'" $0 "'${NC}'"}; /RUN / {print "'${YELLOW}'" $0 "'${NC}'"}; !/FAIL|PASS|RUN / {print}'
    echo -e "${RED}Error: Tests failed in 'services'. Check the test output above for details.${NC}"
    exit 1
else
        echo -e "${output}" | awk '/FAIL/ {print "'${RED}'" $0 "'${NC}'"}; /PASS/ {print "'${GREEN}'" $0 "'${NC}'"}; /RUN / {print "'${YELLOW}'" $0 "'${NC}'"}; !/FAIL|PASS|RUN / {print}'

fi

docker-compose up --build -d

echo -e "${GREEN}Containers are running. Press CTRL+C to stop them.${NC}"

while true; do
    sleep 1
done