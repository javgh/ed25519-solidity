pragma solidity ^0.5.3;

contract Outpost {
    uint public pings = 0;
    uint public paidPings = 0;

    function ping() external {
        pings += 1;
    }

    function paidPing() external payable {
        paidPings += 1;
    }
}
