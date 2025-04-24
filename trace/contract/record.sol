// SPDX-License-Identifier: MIT
pragma solidity 0.8.20;

import "@openzeppelin/contracts/access/Ownable.sol";

contract WorkRecord is Ownable {
    struct WorkDetails {
        string workLocation;
        string workContent;
        string person;
        string advice;
        string workTime;
        string remarks;
        string imagesUrl;
    }

    constructor() Ownable(msg.sender) {}

    // Mapping from number to WorkDetails
    mapping(string => WorkDetails) private records;

    // Function to write records, only callable by the owner
    function addRecord(
        string memory number,
        string memory workLocation,
        string memory workContent,
        string memory person,
        string memory advice,
        string memory workTime,
        string memory remarks,
        string memory imagesUrl
    ) public onlyOwner {
        records[number] = WorkDetails({
            workLocation: workLocation,
            workContent: workContent,
            person: person,
            advice: advice,
            workTime: workTime,
            remarks: remarks,
            imagesUrl: imagesUrl
        });
    }

    // Function to read records by number
    function getRecord(string memory number)
    public
    view
    returns (WorkDetails memory)
    {
        return records[number];
    }
}
