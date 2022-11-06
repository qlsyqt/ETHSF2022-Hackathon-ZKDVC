const {ethers} = require("hardhat");

async function main() {
    const [owner] = await ethers.getSigners();
  
    const State = await ethers.getContractFactory("State", owner);
    const state = await State.deploy();
    await state.deployed();

    const Verifier = await ethers.getContractFactory("StateTransitionVerifier");
    const verifier = await Verifier.deploy();
    await verifier.deployed();

    await (await state.connect(owner).initialize(verifier.address)).wait();


    const DCPClaimVerifier = await ethers.getContractFactory("DCPClaimVerifier", owner);
    const dcpClaimVerifier = await DCPClaimVerifier.deploy();
    await dcpClaimVerifier.deployed();


    const DCPRevocationVerifier = await ethers.getContractFactory("DCPRevocationVerifier", owner);
    const dcpRevocationVerifier = await DCPRevocationVerifier.deploy();
    await dcpRevocationVerifier.deployed();


    const DCPBadge = await ethers.getContractFactory("DCPBadge", owner);
    const dcpBadge = await DCPBadge.deploy("", dcpClaimVerifier.address, dcpRevocationVerifier.address);
    await dcpBadge.deployed();

    console.log("deploy success, by ",owner.address);
    console.log("---------------------------------------------");
    console.log("verifier :", verifier.address);
    console.log("state :", state.address);

    console.log("---------------------------------------------");
    console.log("dedcpClaimVerifierployer :", dcpClaimVerifier.address);
    console.log("dcpRevocationVerifier :", dcpRevocationVerifier.address);
    console.log("dcpBadge :", dcpBadge.address);

}

try {
    main();
} catch (error) {
    console.log(error);
}