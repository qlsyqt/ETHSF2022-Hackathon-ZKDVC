//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.17;

struct SnarkProof {
    uint256[2] a;
    uint256[2][2] b;
    uint256[2] c;
}

struct IssuanceInputs {
    string dataCategory;
    string subCategory;
    uint256[10] publicInputs;
}

struct RevocationInputs {
    string dataCategory;
    string subCategory;
    uint256[5] publicInputs;
}
