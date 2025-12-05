const {ethers, deployments,upgrades} = require("hardhat")
const { expect } = require("chai")

// describe("Starting", async function() {

//     it("Should be able to deploy", async function() {
//         const Contract = await ethers.getContractFactory("NftAuction")
//         const contract = await Contract.deploy()

//         await contract.waitforDeployment()

//        await contract.createAuction(
//            100 * 1000,
//            ethers.parseEther("0.000000001"),
//            ethers.ZeroAddress,
//            1
//        )


//        const auction = await contract.auctions(0);
//        console.log(auction);
//     })
// })





describe("Test upgrade", async function() {
    
    it("Should be able to deploy", async function () {

        //1.部署业务合约
        await deployments.fixture("deployNftAuction")

        const nftAuctionProxy = await deployments.get("NftAuctionProxy")

        //2.调用createAuction 方法创建拍卖
        const nftAuction = await ethers.getContractAt("NftAuction", nftAuctionProxy.address)

        await nftAuction.createAuction(
            100 * 1000,
            ethers.parseEther("0.01"),
            ethers.ZeroAddress,
            1
        )
    
        const auction = await nftAuction.auctions(0);
        console.log("创建拍卖成功：：", auction);


        const implAddress2 = await upgrades.erc1967.getImplementationAddress(nftAuctionProxy.address);
        console.log("implAddress2:::", implAddress2);
        //3.升级合约
        await deployments.fixture(["upgradeNftAuction"]);

        //4. 读取合约的 auction[0]
        const auction2 = await nftAuction.auctions(0);
        console.log("升级后读取拍卖成功：：", auction2);



         const nftAuctionV2 = await ethers.getContractAt("NftAuctionV2", nftAuctionProxy.address);
         const hel = await nftAuctionV2.testHello();
         console.log("hello::", hel);

        expect(auction2.startTime).to.equal(auction.startTime);
        //不相等后面再说
       // expect(implAddress1).to.not(implAddress2);

    }) 
   
})