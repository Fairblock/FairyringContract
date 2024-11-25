// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {Test, console} from "forge-std/Test.sol";
import {FairyringContract} from "../src/FairyringContract.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract FairyringContractTest is Test {
    FairyringContract public fairyringContract;
    bytes public encryptionKey = hex"aa0e6f1b8b1c0d4e8d32f8332a525638db413ea59f2eef45479e9c7e6af7690f6dc296181c012c912be23767b2892d26";
    bytes public decryptionKey = hex"92c53b134ee580542aa0ec2ae8a8233764c1fa3fe024c577ae245df08a8ddede1dee54667ecaad61327cb9a1cbf1d15d08e8e00b64d0931744120fc0e4ce961cd4d3e04beb69bc6b9f075472262d311be62b4d70ab8e69120923ce82af53c7a1";
    uint256 public decryptionKeyHeight = 1;

    function setUp() public {
        fairyringContract = new FairyringContract();
    }

    function test_submitEncryptionKey() public {
        fairyringContract.submitEncryptionKey(encryptionKey);
        assertTrue(keccak256(fairyringContract.latestEncryptionKey()) == keccak256(encryptionKey));
    }

    function test_submitDecryptionKey() public {
        fairyringContract.submitEncryptionKey(encryptionKey);
        fairyringContract.submitDecryptionKey(encryptionKey, decryptionKey, decryptionKeyHeight);
        (bytes32 randoness, uint256 height) = fairyringContract.latestRandomness();
        assertEq(height, decryptionKeyHeight);
        assertTrue(randoness == keccak256(decryptionKey));
    }

    function test_Fail_NotOwnerSubmitEncryptionKey() public {
        vm.startPrank(address(1));
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, address(1)));
        fairyringContract.submitEncryptionKey(encryptionKey);
        vm.stopPrank();
    }

    function test_Fail_NotOwnerSubmitDecryptionKey() public {
        fairyringContract.submitEncryptionKey(encryptionKey);
        vm.startPrank(address(1));
        vm.expectRevert(abi.encodeWithSelector(Ownable.OwnableUnauthorizedAccount.selector, address(1)));
        fairyringContract.submitDecryptionKey(encryptionKey, decryptionKey, decryptionKeyHeight);
        vm.stopPrank();
    }

    function test_Fail_SubmitExistsEncryptionKey() public {
        fairyringContract.submitEncryptionKey(encryptionKey);
         vm.expectRevert(bytes("encryption key already exists"));
        fairyringContract.submitEncryptionKey(encryptionKey);
    }

    function test_Fail_SubmitKeyForNotExistsEncryptionKey() public {
        vm.expectRevert(bytes("encryption key does not exists"));
        fairyringContract.submitDecryptionKey(encryptionKey, decryptionKey, decryptionKeyHeight);
    }

    function test_Fail_SubmitExistsDecryptionKey() public {
        fairyringContract.submitEncryptionKey(encryptionKey);
        fairyringContract.submitDecryptionKey(encryptionKey, decryptionKey, decryptionKeyHeight);
        vm.expectRevert(bytes("decryption key for given height already exists"));
        fairyringContract.submitDecryptionKey(encryptionKey, decryptionKey, decryptionKeyHeight);
    }
}
