// 导入hardhat运行时环境
const hre = require("hardhat");

/**
 * @dev 主部署函数
 */
async function main() {
  // 获取部署者账户（第一个账户）
  const [deployer] = await hre.ethers.getSigners();
  
  console.log("Deploying contracts with account:", deployer.address);

  // 1. 部署 NFT 合约
  console.log("\n1. Deploying AuctionNFT contract...");
  const AuctionNFT = await hre.ethers.getContractFactory("MyNFT");
  const nft = await AuctionNFT.deploy();
  await nft.waitForDeployment(); // 等待部署完成
  const nftAddress = await nft.getAddress();
  console.log("AuctionNFT deployed to:", nftAddress);

  // 2. 部署英式拍卖合约
  console.log("\n2. Deploying EnglishAuction contract...");
  const EnglishAuction = await hre.ethers.getContractFactory("NFTAuctionByEnglishAuction");
  const englishAuction = await EnglishAuction.deploy();
  await englishAuction.waitForDeployment();
  const englishAuctionAddress = await englishAuction.getAddress();
  console.log("EnglishAuction deployed to:", englishAuctionAddress);

  // 4. 保存合约地址到配置文件
  console.log("\n4. Saving contract addresses to config file...");
  const config = {
    nft: nftAddress,
    englishAuction: englishAuctionAddress,
    network: hre.network.name,
    timestamp: new Date().toISOString()
  };

  // 使用文件系统模块
  const fs = require("fs");
  const configFileName = `config-${hre.network.name}.json`;
  
  // 写入配置文件
  fs.writeFileSync(
    configFileName,
    JSON.stringify(config, null, 2) 
  );
  
  console.log(`Config saved to ${configFileName}!`);
  
  // 5. 输出部署摘要
  console.log("\n=== Deployment Summary ===");
  console.log("Network:", hre.network.name);
  console.log("Deployer:", deployer.address);
  console.log("NFT Contract:", nftAddress);
  console.log("English Auction:", englishAuctionAddress);
  console.log("===========================\n");
}

// 执行主函数并处理错误
main()
  .then(() => process.exit(0))
  .catch((error) => {
    console.error("Deployment failed:", error);
    process.exit(1);
  });