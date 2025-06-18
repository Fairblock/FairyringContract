// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

interface IGateway {
    function latestEncryptionKey() external view returns (bytes memory);
    function getRandomnessByHeight(uint256) external view returns (bytes32);
    function latestRandomnessHashOnly() external view returns (bytes32);
    function latestRandomness() external view returns (bytes32, uint256);
    function encryptionKeyExists(bytes memory) external view returns (bool);
    function submitFID(address _requester, string memory _fid, uint256 _id) external;
    function requestGeneralID() external;

}