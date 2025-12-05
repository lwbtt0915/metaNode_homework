const {deployments, upgrades, getUnnamedAccounts} = require("hardhat");
//const { deployments, upgrades, ethers } = require("hardhat");

const fs = require("fs");
const path =require("path");


module.exports = async({getNamedAccounts, deployments}) => {
    const {save} = deployments;
    const {deployer} = await  getUnnamedAccounts();

     console.log("部署用户地址", deployer);
     const NftAuction = await ethers.getContractFactory("NftAuction");


     // 通过代理部署合约
     const NftAuctionProxy = await upgrades.deployProxy(NftAuction, [], {
          initializer:"initialize",
     })


     await NftAuctionProxy.waitForDeployment();

     const proxyAddress = await NftAuctionProxy.getAddress();
     console.log("代理合约地址", proxyAddress);

     const implAddress = await upgrades.erc1967.getImplementationAddress(proxyAddress)
     console.log("实现合约地址", implAddress);


     const storePath = path.resolve(__dirname, "../cache/proxyNftAuction.json");
     console.log("storePath", JSON.stringify({
          proxyAddress,
          implAddress: await  upgrades.erc1967.getImplementationAddress(proxyAddress),
          abi: NftAuction.interface.format("json"),
     }));

     fs.writeFileSync(
          storePath,
          JSON.stringify({
               proxyAddress,
               implAddress: await  upgrades.erc1967.getImplementationAddress(proxyAddress),
               abi: NftAuction.interface.format("json"),
          })
     );


     await save("NftAuctionProxy", {
          abi: NftAuction.interface.format("json"),
          address: proxyAddress,
          args:[],
          log: true,
     })
};


module.exports.tags = ["deployNftAuction"];