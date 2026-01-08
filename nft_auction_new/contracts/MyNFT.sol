
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.28;

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";



contract MyNFT is ERC721,ERC721URIStorage,Ownable {

    //    // 下一个可用的tokenId
    uint256 private _nextTokenId;


   //构造函数，初始化NFT合约
    constructor() ERC721("MyNFT", "MNFT") Ownable(msg.sender) {}



   /**
     * @dev 安全铸造NFT函数，只有合约拥有者可以调用
     * @param to NFT接收者地址
     * @param uri NFT的元数据URI
     * @return 返回铸造的tokenId
     */
    function safeMint(address to, string memory uri) external onlyOwner returns(uint256) {
        uint256 tokenId = _nextTokenId++;
        // 安全铸造NFT（会检查接收者是否能处理ERC721令牌）
        _safeMint(to, tokenId);
         // 设置token的URI（元数据）
        _setTokenURI(tokenId, uri);
        return  tokenId;
    }


    /**
     * @dev 重写tokenURI函数，返回token的URI
     * @param tokenId 要查询的token ID
     * @return token的URI字符串
     */
     function tokenURI(uint256 tokenId) public view override(ERC721, ERC721URIStorage) returns(string memory){
         return super.tokenURI(tokenId);
     }


     /**
     * @dev 重写supportsInterface函数，支持接口查询
     * @param interfaceId 接口ID
     * @return 是否支持该接口
     */
    function supportsInterface(bytes4 interfaceId)
        public
        view
        override(ERC721, ERC721URIStorage)
        returns (bool)
    {
        return super.supportsInterface(interfaceId);
    }
}