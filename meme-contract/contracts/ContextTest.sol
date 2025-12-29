// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/utils/Context.sol";

contract ContextTest is Context {
    event Log(address msgSender, address _msgSender, address thisAddress);

    function whoAmI() external {
        emit Log(msg.sender, _msgSender(), address(this));
    }
}

contract Caller {
    function callWhoAmI(address contextTest) external {
        ContextTest(contextTest).whoAmI();
    }
}
