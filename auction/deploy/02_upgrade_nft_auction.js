const {ethers, upgrades} =require("hardhat")
const fs = require("fs")
const path = require("path")


module.exports = async function({getNamedAccounts, deployments}) {
    const {save} = deployments
    const {deployer} = await getNamedAccounts()
    console.log("部署用户地址", deployer);


    const storePath = path.resolve(__dirname, "../cache/proxyNftAuction.json");
    console.log("storePath:::", storePath);

    const storeData = fs.readFileSync(storePath, "utf-8");

    console.log("storeData:::", storeData);
    const {proxyAddress, implAddress, abi } = JSON.parse(storeData);


    //升级合约
    const nftAuctionV2 = await ethers.getContractFactory("NftAuctionV2");
    console.log("NftAuctionV2:::", nftAuctionV2);

    //升级代理合约
    const nftAuctionProxyV2 = await upgrades.upgradeProxy(proxyAddress, nftAuctionV2);

    console.log("NftAuctionProxyV2:::", nftAuctionProxyV2);

    await nftAuctionProxyV2.waitForDeployment()
    const proxyAddressV2 = await nftAuctionProxyV2.getAddress();


    console.log("proxyAddressV2:::", proxyAddressV2);


    //保存代理合约地址
//     fs.writeFileSync(
//         storePath,
//         JSON.stringify({
//              proxyAddress: proxyAddressV2,
//              implAddress: await  upgrades.erc1967.getImplementationAddress(proxyAddressV2),
//              abi: NftAuction.interface.format("json"),
//         })
//    );


   await save("NftAuctionProxyV2", {
         abi,
        address: proxyAddressV2,
   })


}



module.exports.tags = ["upgradeNftAuction"];