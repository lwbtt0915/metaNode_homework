pragma solidity ^0.8.0;

contract Voting {
     //投票映射
     mapping (address =>uint8) public voteMapping;
      //候选人
     address[] private _candidatesList;

     //投票
     function  vote(address  addr,uint8 _vote)   external {
            voteMapping[addr]+=_vote;
            _candidatesList.push(addr);
     }

 
     //获取得票数
     function getVotes(address addr)  external  returns (uint8 _vote) {
             return voteMapping[addr];
     }


    
     // 重置所有人得票数为0
     function resetVotes() external  {
            for (uint256 i = 0; i < _candidatesList.length; i++) {
            voteMapping[_candidatesList[i]] = 0;
        }
     }

}