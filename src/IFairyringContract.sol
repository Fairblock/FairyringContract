// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

interface IFairyringContract {
    function latestEncryptionKey() external view returns (bytes memory);
    function getRandomnessByHeight(uint256) external view returns (bytes32);
    function latestRandomnessHashOnly() external view returns (bytes32);
    function latestRandomness() external view returns (bytes32, uint256);
    function encryptionKeyExists(bytes memory) external view returns (bool);
}