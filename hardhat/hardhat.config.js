require("@nomicfoundation/hardhat-toolbox");
require('hardhat-abi-exporter');

/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  networks: {
      hardhat: {

      },
      development: {
        url: "http://127.0.0.1:8545"
      }
  },
  solidity: {
    compilers: [
      {
        version: "0.8.17",
      },
      {
        version: "0.8.15",
      },
      {
        version: "0.6.11",
      },
    ],
  },
  abiExporter: {
    path: './data/abi',
    clear: true,
    flat: true,
    spacing: 2
  }
};
