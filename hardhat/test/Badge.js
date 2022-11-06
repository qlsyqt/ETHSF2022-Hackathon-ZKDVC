const {time, loadFixture} = require("@nomicfoundation/hardhat-network-helpers");
const {anyValue} = require("@nomicfoundation/hardhat-chai-matchers/withArgs")
const {expect} = require('chai');
const { isCallTrace } = require("hardhat/internal/hardhat-network/stack-traces/message-trace");


describe("Badge Deployment", function(){

    async function deployFixture() {
        // Contracts are deployed using the first signer/account by default
        const [owner, dummy] = await ethers.getSigners();
    
        const Badge = await ethers.getContractFactory("DCPBadge");
        const instance = await Badge.deploy("aaa", dummy.address, dummy.address);

        await instance.deployed();

        console.log(instance.address);
        return { instance, owner };
      }
    
    it("Basic", async function(){
        const {instance, owner} = await loadFixture(deployFixture);
        
        const dataCategory = "dc";
        const subCategory = "sc";
        const receipt = await (await instance.connect(owner).mintNative(dataCategory, subCategory)).wait();


        const token = await instance.tokenOfOwnerByIndex(owner.address, 0);

        // instance.getCategory
        const category = await instance.getTokenCategory(token);
        expect(category).to.equal(dataCategory);
    })

})