const { ethers } = require("hardhat");

async function main() {
  console.log("开始部署 ERC20 代币合约...");

  // 获取部署者账户
  const [deployer] = await ethers.getSigners();
  console.log("部署者地址:", deployer.address);
  console.log("部署者余额:", ethers.formatEther(await ethers.provider.getBalance(deployer.address)), "ETH");

  // 合约参数
  const tokenName = "My Test Token";
  const tokenSymbol = "MTT";
  const decimals = 18;
  const initialSupply = 1000000; // 100万
  const maxSupply = 10000000; // 1000万

  console.log("\n合约参数:");
  console.log("代币名称:", tokenName);
  console.log("代币符号:", tokenSymbol);
  console.log("小数位数:", decimals);
  console.log("初始供应量:", initialSupply);
  console.log("最大供应量:", maxSupply);

  // 部署合约
  const MyToken = await ethers.getContractFactory("MyToken");
  const myToken = await MyToken.deploy(
    tokenName,
    tokenSymbol,
    decimals,
    initialSupply,
    maxSupply
  );

  await myToken.waitForDeployment();
  const tokenAddress = await myToken.getAddress();

  console.log("\n✅ 代币合约部署成功!");
  console.log("合约地址:", tokenAddress);
  console.log("交易哈希:", myToken.deploymentTransaction().hash);

  // 验证合约信息
  console.log("\n📊 合约信息验证:");
  console.log("总供应量:", ethers.formatUnits(await myToken.totalSupply(), decimals));
  console.log("部署者余额:", ethers.formatUnits(await myToken.balanceOf(deployer.address), decimals));
  console.log("最大供应量:", ethers.formatUnits(await myToken.maxSupply(), decimals));
  console.log("所有者地址:", await myToken.owner());

  return {
    tokenAddress,
    myToken
  };
}

// 部署并验证
async function deployAndVerify() {
  try {
    const { tokenAddress, myToken } = await main();
    
    // 这里可以添加额外的测试或操作
    console.log("\n🎉 部署流程完成!");
    
    return {
      tokenAddress,
      myToken
    };
  } catch (error) {
    console.error("❌ 部署失败:", error);
    process.exit(1);
  }
}

// 如果是直接运行此脚本
if (require.main === module) {
  deployAndVerify();
}

module.exports = deployAndVerify;