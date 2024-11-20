require("@nomicfoundation/hardhat-toolbox");
require("dotenv").config();
/** @type import('hardhat/config').HardhatUserConfig */
module.exports = {
  solidity: "0.8.27",
  networks: {
    hardhat: {},
    sepolia: {
      // url:"https://eth-sepolia.g.alchemy.com/v2/" + process.env.ALCHEMY_API_KEY, 
      url: "https://sepolia.infura.io/v3/" + process.env.INFURA_ID,
      accounts: [`0x${process.env.PRIVATE_KEY}`],
    },
  },
};
