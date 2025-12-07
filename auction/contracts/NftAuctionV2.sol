// SPDX-License-Identifier: GPL-3.0

pragma solidity ^0.8;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";


contract NftAuctionV2 is Initializable {



    struct Auction{
       address seller;
       uint256 duration;
       uint256 startPrice;
       bool ended;
       address highestBidder;
       uint256 startTime;
       uint256 highestBid;
       //nft 合约地址
       address nftContract;
       // NFT id 
       uint256 tokenId;
    }

    mapping(uint256 => Auction) public auctions;

    uint256 public nextAuctionId;

    address public admin;

   function  initialize() initializer public {
        admin= msg.sender;
    }



    //创建拍卖
    function createAuction(uint256 _duration, uint256 _startPrice ,
    address _nftAddress ,uint256 _tokenId)public {
        require(msg.sender == admin ,"Only admin create auctions");
        require(_duration >0 ,"Duration must be greater than 0");
        require(_startPrice > 0, "Start price must be greater than 0");


        auctions[nextAuctionId] =Auction ({
           seller: msg.sender,
            duration:_duration,
           startPrice: _startPrice,
           ended: false,
           highestBidder: address(0),
           startTime: block.timestamp,
           highestBid: 0,
           nftContract: _nftAddress,
           tokenId: _tokenId
        });

        nextAuctionId++;


    }


    //参与拍卖买单
    function placeBid(uint256 _auctionId) external  payable {
        Auction storage auction = auctions[_auctionId];
       require(!auction.ended && (auction.startTime+auction.duration <= block.timestamp), "Auction has ended.");
       // 判断出价是否大于当前最高出价
       require(msg.value > auction.highestBid && msg.value >= auction.startPrice, "Bid must ve higher than the current highest bid.");

       //返回之前的最高出价者
       if(auction.highestBidder != address(0)) {
            payable(auction.highestBidder).transfer(auction.highestBid);
       }

       auction.highestBidder = msg.sender;
       auction.highestBid = msg.value;
    }




    function testHello() public pure returns(string memory) {
        return "Hello, World!";
    }
}
