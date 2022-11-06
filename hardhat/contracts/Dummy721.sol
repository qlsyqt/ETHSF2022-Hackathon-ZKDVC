//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.17;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721Enumerable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/Context.sol";
import {SnarkProof, IssuanceInputs, RevocationInputs} from "./lib/Structs.sol";

//Mint on issuance
//Burn on revocation
//Transfer is not allowed
//本来徽章是用1155做的，但是1155目前的功能不支持遍历
contract Dummy721 is Context, Ownable,ERC721Enumerable{
    
    constructor(uint256 amount) ERC721("Dummy","Dummy"){
        for (uint i=0;i<amount;i++) {
            uint256 token = uint256(keccak256(abi.encodePacked(msg.sender, i)));
            _mint(msg.sender, token);
        }
    }

}   
