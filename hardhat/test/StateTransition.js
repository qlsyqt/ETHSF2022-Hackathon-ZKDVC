const {
    time,
    loadFixture,
  } = require("@nomicfoundation/hardhat-network-helpers");
  const { anyValue } = require("@nomicfoundation/hardhat-chai-matchers/withArgs");
  const { expect } = require("chai");

  const { waffle } = require("hardhat");
  require("@nomiclabs/hardhat-waffle");

  
  describe("state", function () {
    // We define a fixture to reuse the same setup in every test.
    // We use loadFixture to run this setup once, snapshot that state,
    // and reset Hardhat Network to that snapshot in every test.
    // async function setStateAndVerifier() {
  
    //   // Contracts are deployed using the first signer/account by default
    //   const [owner] = await ethers.getSigners();
  
    //   const State = await ethers.getContractFactory("State", owner);
    //   const state = await State.deploy();
    //   await state.deployed();

    //   const Verifier = await ethers.getContractFactory("StateTransitionVerifier");
    //   const verifier = await Verifier.deploy();
    //   await verifier.deployed();

    //   await (await state.connect(owner).initialize(verifier.address)).wait();
  
    //   return { state};
    // }
  
    // it(" test update genesis state to a new state", async function() {
    //     const {state} = await loadFixture(setStateAndVerifier);
    //     console.log(state.address);
    //     const [owner, second] = await ethers.getSigners();

    //     //genesis state to a new state.
    //     //did generated from bjj public key c0fd4324c51cdc9ee10f901fb575a4d11b95fd4f2cf6eaf7052e2a19bbf9e299 and revocation nonce 1.
    //     const id = BigInt("26592832776784078789941335656836033229340999336920335567255332674191097856");
    //     const oldState = BigInt("6779716606974252845410404765497113182232614656151360295172244272544411487259");
    //     const newState = BigInt("4261880322420347246246787914971343760440730417906531557674114299270977011541");
    //     var isOldStateGenesis = true;
    //     const a = [BigInt("19906498506743515988231000715034702069804532539319427094231258289242945036433"),
    //     BigInt("16453554573493818404981023035227585844717081809803769303069333304508973623807")]
    //     const b = [[ BigInt("10959501866750999589686894037049971570933928933336256642988639488364741414089"),    
    //         BigInt("5617966927892693058820744748843761755536598213483999945665668773831731711461"), 
    //                    ],
    //                 [    BigInt("4741235647821872646751096733509153996636577406682280623147782070687106971212"),
    //                     BigInt("15308259226983523787237346826171248952127720081608066981273175425338262231053"), 
    //                    ]
    //                 ]
    //     const c = [BigInt("19069303097284893759313273250342034797071976886801554427714044381134282618830") ,
    //         BigInt("16273866848661828260616242221159357034444830503370082159601148718616714905552")]
                    
    //     const tx = await state.connect(second).transitState(id, oldState, newState, isOldStateGenesis, a, b, c)
    //     console.log(tx.nonce);
    //     const receipt = await tx.wait();
    //     console.log(receipt);

    // })
    it(" get tx count", async function() {
        console.log(waffle.provider);

    })

  });
  
  