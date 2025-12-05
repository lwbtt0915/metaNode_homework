// SPDX-License-Identifier: MIT
pragma solidity ^0.8.19;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Burnable.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC20/extensions/ERC20Permit.sol";

contract MyToken is ERC20, ERC20Burnable, Ownable, ERC20Permit {
    uint8 private _decimals;
    uint256 public maxSupply;

    event TokensMinted(address indexed to, uint256 amount);
    event TokensBurned(address indexed from, uint256 amount);

    constructor(
        string memory name,
        string memory symbol,
        uint8 decimalUnits,
        uint256 initialSupply,
        uint256 _maxSupply
    ) 
        ERC20(name, symbol) 
        Ownable(msg.sender)
        ERC20Permit(name)
    {
        _decimals = decimalUnits;
        maxSupply = _maxSupply * 10 ** decimalUnits;
        
        uint256 initialAmount = initialSupply * 10 ** decimalUnits;
        require(initialAmount <= maxSupply, "Initial supply exceeds max supply");
        
        _mint(msg.sender, initialAmount);
    }

    function decimals() public view virtual override returns (uint8) {
        return _decimals;
    }

    /**
     * @dev 铸造新代币，仅所有者可调用
     */
    function mint(address to, uint256 amount) public onlyOwner {
        require(totalSupply() + amount <= maxSupply, "Exceeds max supply");
        _mint(to, amount);
        emit TokensMinted(to, amount);
    }

    /**
     * @dev 批量转账
     */
    function batchTransfer(
        address[] memory recipients, 
        uint256[] memory amounts
    ) public returns (bool) {
        require(recipients.length == amounts.length, "Arrays length mismatch");
        
        for (uint256 i = 0; i < recipients.length; i++) {
            require(transfer(recipients[i], amounts[i]), "Transfer failed");
        }
        return true;
    }

    /**
     * @dev 获取剩余可铸造数量
     */
    function getRemainingMintable() public view returns (uint256) {
        return maxSupply - totalSupply();
    }

    /**
     * @dev 重写 burn 函数，添加事件
     */
    function burn(uint256 amount) public override {
        super.burn(amount);
        emit TokensBurned(msg.sender, amount);
    }

    /**
     * @dev 重写 burnFrom 函数，添加事件
     */
    function burnFrom(address account, uint256 amount) public override {
        super.burnFrom(account, amount);
        emit TokensBurned(account, amount);
    }
}