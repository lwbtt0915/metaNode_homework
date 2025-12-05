/** @type import('hardhat/config').HardhatUserConfig */

require("@nomicfoundation/hardhat-toolbox");
require("hardhat-deploy");
require("@openzeppelin/hardhat-upgrades");

module.exports = {
  networks: {
    localhost: {
      url: "http://127.0.0.1:8545",  // 默认端口
      chainId: 31337,  // Hardhat网络默认chainId
    }
  },
  solidity: "0.8.28",
};
