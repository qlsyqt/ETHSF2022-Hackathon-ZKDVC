// We require the Hardhat Runtime Environment explicitly here. This is optional
// but useful for running the script in a standalone fashion through `node <script>`.
//
// You can also run a script with `npx hardhat run <script>`. If you do that, Hardhat
// will compile your contracts, add the Hardhat Runtime Environment's members to the
// global scope, and execute the script.
const {ethers} = require("hardhat");

async function main() {
  const [owner] = await ethers.getSigners();
  console.log(`owner address:" ${owner.address}`)
  const ERC721 = await ethers.getContractFactory("Dummy721", owner);
  const instance = await ERC721.deploy(5);
  await instance.deployed();
  console.log(
    'deploy complete:',instance.address
  );
}

// We recommend this pattern to be able to use async/await everywhere
// and properly handle errors.
main().catch((error) => {
  console.error(error);
  process.exitCode = 1;
});
