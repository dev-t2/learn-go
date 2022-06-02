// SPDX-License-Identifier: MIT
pragma solidity >=0.4.22 <0.9.0;

contract Lottery {
  uint256 constant internal BET_BLOCK_LIMIT = 256;
  uint256 constant internal BET_BLOCK_INTERVAL = 3;
  uint256 constant internal BET_AMOUNT = 5 * 10 ** 15;

  uint256 private totalPot;

  struct BetInfo {
    uint256 answerBlockNumber;
    address payable bettor;
    bytes1 challenges;
  }

  uint256 private first;
  uint256 private last;

  mapping (uint256 => BetInfo) private bets;

  function getTotalPot() public view returns (uint256 pot) {
    return totalPot;
  }

  function getBetInfo(uint256 index) public view returns (uint256 answerBlockNumber, address bettor, bytes1 challenges) {
    BetInfo memory betInfo = bets[index];

    answerBlockNumber = betInfo.answerBlockNumber;
    bettor = betInfo.bettor;
    challenges = betInfo.challenges;
  }

  function pushBet(bytes1 challenges) public returns (bool) {
    BetInfo memory betInfo;

    betInfo.answerBlockNumber = block.number + BET_BLOCK_INTERVAL;
    betInfo.bettor = payable(msg.sender);
    betInfo.challenges = challenges;

    bets[last] = betInfo;

    last = last + 1;

    return true;
  }

  function popBet(uint256 index) public returns (bool) {
    delete bets[index];

    return true;
  }
}