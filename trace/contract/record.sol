// SPDX-License-Identifier: MIT
pragma solidity 0.8.19;

contract Lock {
    mapping(address => uint256) public Locked;
    mapping(Op=>mapping(uint64 => bool)) public used;

    enum Op {
        deposit,
        withdrawl
    }
    event Assert(
        address from,
        uint64 nonce,
        uint256 amount,
        uint256 blocknum,
        uint256 timestamp,
        Op op
    );

    constructor() {}

    function deposit(uint64 nonce) external payable {
        uint256 amount = msg.value;
        address sender = msg.sender;
        Locked[sender] += amount;
        require(!used[Op.deposit][nonce]);
        used[Op.deposit][nonce] = true;
        emit Assert(
            sender,
            nonce,
            amount,
            block.number,
            block.timestamp,
            Op.deposit
        );
    }

    function withdrawl(uint64 amount, uint64 nonce) external {
        address sender = msg.sender;
        Locked[sender] -= amount;
        payable(sender).transfer(amount);
        require(!used[Op.withdrawl][nonce]);
        used[Op.withdrawl][nonce] = true;
        emit Assert(
            sender,
            nonce,
            amount,
            block.number,
            block.timestamp,
            Op.withdrawl
        );
    }
}
