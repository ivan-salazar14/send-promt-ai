#!/bin/bash

# Configuration
URL="http://localhost:8080/process"
TOKEN="my-secure-token"
PROMPT="Explain the importance of concurrency in Go in one sentence."

echo "Starting concurrent test with 5 simultaneous requests..."

# We use an ampersand (&) to run curl commands in the background
for i in {1..5}
do
   curl -X POST $URL \
        -H "Content-Type: application/json" \
        -H "Authorization: Bearer $TOKEN" \
        -d "{\"prompt\": \"$PROMPT (Request #$i)\"}" &
done

# Wait for all background processes to finish
wait
echo -e "\nAll requests completed."