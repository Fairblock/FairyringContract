#!/bin/bash

cat ./out/Gateway.sol/Gateway.json | jq '.abi' > gateway.json
abigen --abi gateway.json --pkg contract --out gateway.go --type Gateway

rm gateway.json