// SPDX-License-Identifier: MIT
pragma solidity ^0.8.25;

import {IGateway} from "../src/IGateway.sol";
import {Context} from "@openzeppelin/contracts/utils/Context.sol";
import {Ownable} from "@openzeppelin/contracts/access/Ownable.sol";

contract NumberGuessing is Context, Ownable {
    uint256 public roundDuration;

    struct Round {
        uint256 startBlock;
        uint256 endAtBlock;
        bool ended;
        address[] winners;
        uint256 winningNumber;
    }

    mapping(uint256 => Round) public rounds;
    mapping(uint256 => mapping(uint256 => address[])) public roundGussess;

    uint256 public currentRoundID = 1;
    IGateway public gateway;

    constructor(
        uint256 _roundDuration,
        IGateway _gateway
    ) Ownable(_msgSender()) {
        roundDuration = _roundDuration;
        gateway = _gateway;
    }

    function startNewRound() external onlyOwner {
        require(!rounds[currentRoundID].ended, "Current round has not ended");
        address[] memory emptyArray;
        rounds[currentRoundID] = Round({
            startBlock: block.number,
            endAtBlock: block.number + roundDuration,
            ended: false,
            winners: emptyArray,
            winningNumber: 0
        });
    }

    function endCurrentRound() external onlyOwner {
        Round storage currentRound = rounds[currentRoundID];
        
        require(!currentRound.ended, "The round has already ended");
        require(currentRound.endAtBlock <= block.number, "The round is still ongoing");

        currentRound.winningNumber = uint256(gateway.latestRandomnessHashOnly()) % 100;

        currentRound.ended = true;
        currentRound.winners = roundGussess[currentRoundID][currentRound.winningNumber];

        currentRoundID++;
    }

    function makeAGuess(uint256 guess) external {
        require(tx.origin == msg.sender, "Not EOA");
        require(!rounds[currentRoundID].ended, "The round has already ended");
        require(guess >= 1 && guess <= 100, "Guess must be between 1 and 100");

        roundGussess[currentRoundID][guess].push(msg.sender);
    }

    function getRoundWinners(uint256 roundID) external view returns (address[] memory) {
        require(roundID > 0 && roundID < currentRoundID, "Invalid Round ID");
        require(rounds[roundID].ended, "Target round has not ended");
        return rounds[roundID].winners;
    }

    function getLastRoundInfo() external view returns (Round memory) {
        require(currentRoundID > 1, "No Last Round");
        return rounds[currentRoundID - 1];
    }

    function getGuessedAddrs(uint256 roundID, uint256 gussedNumber) external view returns (address[] memory) {
        require(roundID > 0 && roundID <= currentRoundID, "Invalid Round ID");
        return roundGussess[roundID][gussedNumber];
    }
}