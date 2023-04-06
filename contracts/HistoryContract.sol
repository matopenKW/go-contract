// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.17;

contract HistoryContract {
    struct History {
        uint256 timestamp;
        bytes32 data;
    }

    mapping(address => History[]) private histories;

    event DataStored(address indexed user, uint256 indexed recordIndex, bytes32 data, uint256 timestamp);

    function storeData(bytes32 _data) public {
        History memory newHistory = History(block.timestamp, _data);
        histories[msg.sender].push(newHistory);

        uint256 recordIndex = histories[msg.sender].length - 1;
        emit DataStored(msg.sender, recordIndex, _data, block.timestamp);
    }

    function getData(uint256 _index) public view returns (History memory) {
        require(_index < histories[msg.sender].length, "Invalid index");
        return histories[msg.sender][_index];
    }
}
