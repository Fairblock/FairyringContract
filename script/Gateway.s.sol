// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.25;

import {Script, console} from "forge-std/Script.sol";
import {Gateway} from "../src/Gateway.sol";


contract GatewayScript is Script {
    function setUp() public {}

    function run() public {
        vm.startBroadcast();

        Gateway fc = new Gateway();
        console.log("Gateway Deployed at:");
        console.logAddress(address(fc));

        vm.stopBroadcast();
    }
}
