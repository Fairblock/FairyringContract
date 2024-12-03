## Fairyring Contract

Fairyring Contract is an EVM smart contract written in Solidity that allows `fairyport` to relay decryption key from `Fairyring` to target EVM chain.

Fairyrng Contract stores encryption keys & decryption keys relayed by `fairyport` from `Fairyring`, developer can query the latest decryption key and encryption key from this smart contract. The contract also included a function for querying a verifiable randomness that's calculated from the decryption key.

If the target EVM chain contains the `decrypt` pre-compiles from Fairblock, this contract can also verify the randomness on-chain. (Verifying randomness is NOT yet implemented)

We also created an example contract that utilize the Fairyring Contract and a script that spin up a local development environment using `anvil` from `foundry`, The script spin up a local EVM chain, deploy the Fairyring Contract and the example number guessing contract, and submit decryption key to the contract every block.

## Usage

### Build

```shell
$ forge build
```

### Test

```shell
$ forge test
```

### Deploy

```shell
$ forge script script/FairyringContract.s.sol:FairyringContractScript --rpc-url <your_rpc_url> --private-key <your_private_key>
```

### Deploy Number Guessing example contract

```shell
$ forge create example/NumberGuessing.sol:NumberGuessing --constructor-args <RoundDuration> <FairyringContractAddress> --rpc-url <your_rpc_url> --private-key <your_private_key>
```

### Development Environment

```shell
$ ./devenv.sh
```

