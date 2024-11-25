#!/bin/bash

cat ./out/FairyringContract.sol/FairyringContract.json | jq '.abi' > fairyringContract.json
abigen --abi fairyringContract.json --pkg contract --out fairyring_contract.go --type FairyringContract

rm fairyringContract.json