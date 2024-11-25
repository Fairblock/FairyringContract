// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {IFairyringContract} from "./IFairyringContract.sol";
import {Context} from "@openzeppelin/contracts/utils/Context.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract FairyringContract is IFairyringContract, Context, Ownable {
    
    bytes public latestEncryptionKey;
    mapping(bytes => bool) public encryptionKeyExists;

    mapping(uint256 => bytes) decryptionKeys;
    uint256 latestDecryptionKeyHeight = 0;

    constructor() Ownable(_msgSender()) {}

    function submitEncryptionKey(
        bytes memory encryptionKey
    ) external onlyOwner() {
        require(!encryptionKeyExists[encryptionKey], "encryption key already exists");
        latestEncryptionKey = encryptionKey;
        encryptionKeyExists[encryptionKey] = true;
    }

    function submitDecryptionKey(
        bytes memory encryptionKey,
        bytes memory decryptionKey,
        uint256 height
    ) external onlyOwner() {
        require(decryptionKeys[height].length == 0, "decryption key for given height already exists");
        require(encryptionKeyExists[encryptionKey], "encryption key does not exists");

        decryptionKeys[height] = decryptionKey;
        latestDecryptionKeyHeight = height;
    }

    function getRandomnessByHeight(uint256 height) external view returns (bytes32) {
        return _getDecryptionKey(height);
    }

    function latestRandomnessHashOnly() external view returns (bytes32) {
        return _getDecryptionKey(latestDecryptionKeyHeight);
    }

    function latestRandomness() external view returns (bytes32, uint256) {
        return (_getDecryptionKey(latestDecryptionKeyHeight), latestDecryptionKeyHeight);
    }

    function _getDecryptionKey(uint256 height) internal view returns (bytes32) {
         if (decryptionKeys[height].length == 0) {
            return bytes32(0);
        } else {
            return keccak256(decryptionKeys[height]);
        }
    }
}
