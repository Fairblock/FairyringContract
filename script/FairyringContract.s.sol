// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.25;

import {Script, console} from "forge-std/Script.sol";
import {FairyringContract} from "../src/FairyringContract.sol";


contract FairyringContractScript is Script {
    function setUp() public {}

    function run() public {
        vm.startBroadcast();

        FairyringContract fc = new FairyringContract();
        console.log("FairyringContract Deployed at:");
        console.logAddress(address(fc));

        vm.stopBroadcast();
    }
}
