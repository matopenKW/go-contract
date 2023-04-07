// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.17;

contract HistoryContract {
    struct History {
        uint256 timestamp;
        string data;
    }

    mapping(address => History[]) private histories;

    event DataStored(address indexed user, uint256 indexed recordIndex, string data, uint256 timestamp);

    function storeHistory(string memory _data) public {
        History memory newHistory = History(block.timestamp, _data);
        histories[msg.sender].push(newHistory);

        uint256 recordIndex = histories[msg.sender].length - 1;
        emit DataStored(msg.sender, recordIndex, _data, block.timestamp);
    }

    function getHistory(uint256 _index) public view returns (History memory) {
        require(_index < histories[msg.sender].length, "Invalid index");
        return histories[msg.sender][_index];
    }

    function getHistoryFromTimestamp(uint256 timestamp) public view returns (History memory) {
        uint len = histories[msg.sender].length;
        for (uint i = 0; i< len; i++){
            if (histories[msg.sender][i].timestamp == timestamp) {
                return histories[msg.sender][i];
            }
        }
        return History(0, "");
    }
}
