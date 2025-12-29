const { expect } = require("chai");
const { ethers } = require("hardhat");

describe("Context Comparison", function () {
    let contextTest, caller, owner, other, contextAddr;

    beforeEach(async () => {
        [owner, other] = await ethers.getSigners();

        const ContextTest = await ethers.getContractFactory("ContextTest");
        contextTest = await ContextTest.deploy();
        await contextTest.waitForDeployment();

        const Caller = await ethers.getContractFactory("Caller");
        caller = await Caller.deploy();
        await caller.waitForDeployment();
    });

    it("直接调用 ContextTest", async () => {
        const tx = await contextTest.connect(other).whoAmI();
        const receipt = await tx.wait();

        const log = receipt.logs.find((l) => l.fragment.name === "Log");
        expect(log.args.msgSender).to.equal(other.address);
        expect(log.args._msgSender).to.equal(other.address);
        expect(log.args.thisAddress).to.equal(await contextTest.getAddress());
    });

    it("通过 Caller 合约调用 ContextTest", async () => {
        const tx = await caller.connect(other).callWhoAmI(await contextTest.getAddress());
        const receipt = await tx.wait();

        // 逐条解析 log，直到找到 Log 事件
        let parsedEvent;

        for (const log of receipt.logs) {
            try {
                const parsed = contextTest.interface.parseLog(log);
                if (parsed?.name === "Log") {
                    parsedEvent = parsed;
                    break;
                }
            } catch (e) {
                // 忽略非本合约事件
            }
        }

        expect(parsedEvent, "Log event not found").to.not.be.undefined;

        const [msgSender, _msgSender, thisAddress] = parsedEvent.args;

        expect(msgSender).to.equal(await caller.getAddress());
        expect(_msgSender).to.equal(await caller.getAddress());
        expect(thisAddress).to.equal(await contextTest.getAddress());
    });
});
