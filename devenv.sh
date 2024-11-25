#!/bin/bash

echo "Killing previous anvil devnet..."
pkill anvil

# Start devnet
anvil --accounts 3 --balance 100000 --config-out anvil_config.json > devenv.log 2>&1 &
PID=$!
echo "Starting local node in the background, pid: $PID"
sleep 1
pkey=$(cat anvil_config.json | jq -r '.private_keys[0]')
addr=$(cast wallet address --private-key $pkey)

deploy_out=$(forge create src/FairyringContract.sol:FairyringContract --private-key $pkey --json | jq)
# Deploy contract
frc=$(echo $deploy_out | jq -r '.deployedTo')

echo "Contract deployed at: $frc"
echo "Setting up the contract..."

shares=$(fairyringd share-generation generate 1 1)
pubkey=$(echo $shares | jq -r '.MasterPublicKey')
master_share=$(echo $shares | jq -r '.Shares[0].Value')

out=$(cast send $frc "submitEncryptionKey(bytes)" $pubkey --private-key $pkey --json)
echo "Submitted Encryption key, status: $(echo $out | jq -r '.status'), TXID: $(echo $out | jq -r '.transactionHash')"
height=1

echo ""
echo "Press Enter to exit"
echo ""

while true;
do
    derived=$(fairyringd share-generation derive $master_share 1 $height | jq -r '.Keyshare')
    out=$(cast send $frc "submitDecryptionKey(bytes,bytes,uint256)" $pubkey $derived $height --private-key $pkey --json)
    echo "Contract Address: $frc | Submitted Randomness for Height: $height"
    height=$((height+1))

    if read -t 1 -n 1 key; then
        break
    fi
done

kill $PID
rm devenv.log
rm anvil_config.json
echo "Cleaned dev env data"