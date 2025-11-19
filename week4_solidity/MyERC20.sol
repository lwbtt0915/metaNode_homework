// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;


contract MyERC20 {
    //代币名称
    string public name;
    //代币符号
    string public symbol;
    //小数位
    uint8 public constant decimals = 18;
    //总供应量
    uint256 public totalSupply;
    //余额映射 Address ——》 余额
    mapping (address => uint256) public balanceOf;
    //授权映射 owner → spender → 剩余授权额度
    mapping (address => mapping(address => uint256)) public allowance;
    //合约所有者
    address private immutable _owner;

    //转账事件：from 向 to 转移 value 代币时触发
    event Transfer(address indexed from, address indexed to, uint256 value);

    //授权事件：owner 授权 spender 转移 value 代币时触发
    event Approval(address indexed owner, address indexed  spender, uint256 value);


    modifier onlyOwner() {
        require(msg.sender == _owner, "ERC20: caller is not the owner");
        _;
    }


    constructor
    (
        string memory _name,
        string memory _symbol,
        uint256 _totalSupply
    )
    {
        name = _name;
        symbol = _symbol;
        totalSupply = _totalSupply;
        balanceOf[msg.sender] = _totalSupply;
        _owner = msg.sender;
    }

    //从调用者账户向to 转移 value 最小代币单位
    function transfer(address to, uint256 value) public returns (bool){
        address owner =msg.sender;
        _transfer(owner, to, value);

        return true;
    }


    //内部转账逻辑
    function _transfer(
        address from,
        address to,
        uint256 value
    ) internal {
        require(from != address(0), "ERC20: transfer from the zero address"); // 禁止从0地址转账
        require(to != address(0), "ERC20: transfer to the zero address"); // 禁止转账到0地址
        require(balanceOf[from] >= value, "ERC20: insufficient balance"); // 余额不足校验

        // 扣减发送者余额，增加接收者余额
        balanceOf[from] -= value;
        balanceOf[to] += value;

        // 触发转账事件（严格遵循 ERC20 标准）
        emit Transfer(from, to, value);
    }

    //授权：调用者（owner）授权 spender 转移 value 最小代币单位
    function approval(address spender,uint256 value) public returns(bool) {
        address owner = msg.sender;
        _approve(owner, spender, value);
        return true;
    }


    //内部授权逻辑
    function _approve(address owner, address spender, uint256 value) internal {
        require(owner != address(0), "ERC20: transfer from the zero address"); // 禁止从0地址转账
        require(spender != address(0), "ERC20: transfer to the zero address"); // 禁止转账到0地址
        require(balanceOf[owner] >= value, "ERC20: insufficient balance"); // 余额不足校验

        // 扣减发送者余额，增加接收者余额
        balanceOf[owner] -= value;
        balanceOf[spender] += value;

        // 触发转账事件（严格遵循 ERC20 标准）
        emit Transfer(owner, spender, value);
    }



    //授权转账：spender 从 from 账户向 to 转移 value
    function transferFrom(address from, address to, uint256 value) public returns (bool) {
        address spender = msg.sender;
        //扣减授权额度
        _updateAllowance(from, spender, value);

        //执行转账
        _transfer( from, to, value);

        return true;
    }


    //内部授权额度扣减逻辑
    function _updateAllowance(address owner,address spender,uint256 value) internal {
        uint256 currentAllowance = allowance[owner][spender];
        if (currentAllowance != type(uint256).max) { // 无限授权（max）不扣减
            require(currentAllowance >= value, "ERC20: insufficient allowance"); // 授权额度不足校验
            allowance[owner][spender] = currentAllowance - value;
        }
    }

    // ========== 增发功能（仅所有者） ==========
    /**
     * @dev 增发代币：向 to 地址增发 value 最小代币单位
     * @param to 接收增发代币的地址
     * @param value 增发数量（最小代币单位）
     * @return 增发成功返回 true
     */
    function mint(address to, uint256 value) public onlyOwner returns (bool) {
        _mint(to, value);
        return true;
    }

    /**
     * @dev 内部增发逻辑
     */
    function _mint(address to, uint256 value) internal {
        require(to != address(0), "ERC20: mint to the zero address"); // 禁止增发到0地址

        // 增加总供应量和接收者余额
        totalSupply += value;
        balanceOf[to] += value;

        // 触发转账事件（from=0地址表示增发，符合 ERC20 惯例）
        emit Transfer(address(0), to, value);
    }
}