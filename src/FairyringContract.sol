// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {IFairyringContract} from "./IFairyringContract.sol";
import {Context} from "@openzeppelin/contracts/utils/Context.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract FairyringContract is IFairyringContract, Context, Ownable {

    event RequestGeneralID(address indexed requester, uint256 id);
    event RequestGeneralKey(address indexed requester, uint256 id);

    uint256 generalRequestFee = 0;
    
    bytes public latestEncryptionKey;
    mapping(bytes => bool) public encryptionKeyExists;

    mapping(uint256 => bytes) decryptionKeys;
    uint256 latestDecryptionKeyHeight = 0;

    mapping(address =>  uint256) public addressGeneralID;
    mapping(address => mapping(uint256 => bool)) public generalIDRequested;
    mapping(address => mapping(uint256 => bool)) public generalKeyRequested;
    mapping(address => mapping(uint256 => bytes)) public generalDecryptionKeys;

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

    function submitGeneralDecryptionKeys(
        address requester,
        uint256 id,
        bytes memory decryptionKey
    ) external onlyOwner() {
        require(generalIDRequested[requester][id], "The given requester & ID have not requested the general identity");
        require(generalKeyRequested[requester][id], "The given requester & ID have not requested the general decryption key");
        require(generalDecryptionKeys[requester][id].length == 0, "Decryption key for the given requester & ID already exists");
        
        generalDecryptionKeys[requester][id] = decryptionKey;

        // TODO: Do a contract callback with the decryption key here ?
    }

    // TODO: Update the allow the requester to set a target contract address 
    // and function to do callback when the decryption key received
    function requestGeneralDecryptionKey(uint256 id) external {
        require(generalIDRequested[_msgSender()][id], "Given ID is not requested");
        require(!generalKeyRequested[_msgSender()][id], "Already request decryption key for the given ID");
        generalKeyRequested[_msgSender()][id] = true;
        emit RequestGeneralKey(_msgSender(), id);
    }

    function requestGeneralID() external {
        uint256 id = addressGeneralID[_msgSender()];
        generalIDRequested[_msgSender()][id] = true;
        addressGeneralID[_msgSender()] = addressGeneralID[_msgSender()] + 1;
        emit RequestGeneralID(_msgSender(), id);
    }

    function setRequestGeneralFee(uint256 newFee) external onlyOwner() {
        generalRequestFee = newFee;
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
