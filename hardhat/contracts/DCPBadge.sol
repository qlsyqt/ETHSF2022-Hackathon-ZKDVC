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
contract DCPBadge is Context, Ownable,ERC721Enumerable{
    
    IDcpClaimVerifier private dcpClaimVerifier;
    IDcpRevocationVerifier private dcpRevocationVerifier;
    string private uri;

    mapping(uint256=>string) private categories;

    event DCPBadgeMint(uint256 indexed id, address to, address operator);
    event DCPBadgeBurnt(uint256 indexed id, address to, address operator);

    constructor(string memory uri_, address dcpClaimVerifier_, address dcpRevocationVerifier_) ERC721("DCPBadge","DCPBadge"){
        uri = uri_;
        dcpClaimVerifier = IDcpClaimVerifier(dcpClaimVerifier_);
        dcpRevocationVerifier = IDcpRevocationVerifier(dcpRevocationVerifier_);
    }

    function mint(address to_, SnarkProof memory snarkProof_, IssuanceInputs memory inputs_) external {
        require(to_ != address(0), "DCPBadge: Invalid to address");
        require(dcpClaimVerifier.verifyProof(snarkProof_.a, snarkProof_.b, snarkProof_.c, inputs_.publicInputs), "DCPBadge: Invalid proofs");
        //TODO: also verify connections between args and inputs
    
        uint256 id = uint256(keccak256(abi.encodePacked(inputs_.dataCategory, inputs_.subCategory)));    
        _mint(to_, id);
        emit DCPBadgeMint(id, to_, _msgSender());
    }


    //TODO: mocked .Will be deleted later
    function mintNative(string memory dataCategory, string memory subCategory) external {
        uint256 id = uint256(keccak256(abi.encodePacked(dataCategory, subCategory)));    
        _mint(msg.sender, id);
        categories[id] = dataCategory;
        emit DCPBadgeMint(id, _msgSender(), _msgSender());
    }

    function getTokenCategory(uint256 token) external view returns(string memory){
        require(_exists(token), "Token not exist");
        return categories[token];
    } 

    function burn(address from_, SnarkProof memory snarkProof_, RevocationInputs memory inputs_) external {
        require(from_ != address(0), "DCPBadge: Invalid to address");
        require(dcpRevocationVerifier.verifyProof(snarkProof_.a, snarkProof_.b, snarkProof_.c, inputs_.publicInputs), "DCPBadge: Invalid proofs");
        //TODO: also verify connections between args and inputs
        
        uint256 id = uint256(keccak256(abi.encodePacked(inputs_.dataCategory, inputs_.subCategory)));    
        require(ERC721.ownerOf(id) == from_, "DCPBadge: Not owner");
        _burn(id);
        emit DCPBadgeBurnt(id, from_, _msgSender());
    }

    function _beforeTokenTransfer(
        address from,
        address to,
        uint256 id
    ) internal override ( ERC721Enumerable) {
        if (from != address(0)) {
            revert("Transfer is not allowed.");
        }
        super._beforeTokenTransfer(from, to, id);
    }

    function _baseURI() internal view override returns (string memory) {
        return uri;
    }

    function setBaseURI(string calldata uri_) external onlyOwner {
        uri = uri_;
    }
    function setClaimVerifier(address dcpClaimVerifier_) external onlyOwner{
        dcpClaimVerifier = IDcpClaimVerifier(dcpClaimVerifier_);
    }

    function setRevocationVerifier(address dcpRevocationVerifier_) external onlyOwner{
        dcpRevocationVerifier = IDcpRevocationVerifier(dcpRevocationVerifier_);
    }
    
}   

interface IDcpClaimVerifier {
        function verifyProof(
            uint[2] memory a,
            uint[2][2] memory b,
            uint[2] memory c,
            uint[10] memory input
        ) external view returns (bool r) ;
}

interface IDcpRevocationVerifier {
        function verifyProof(
            uint[2] memory a,
            uint[2][2] memory b,
            uint[2] memory c,
            uint[5] memory input
        ) external view returns (bool r) ;
}