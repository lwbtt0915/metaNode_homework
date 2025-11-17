pragma solidity ^0.8.0;
contract StringConvert {
    //字符串反正函数
    function convert(string  memory str) public pure returns (string memory) {
        bytes memory bs =  bytes(str);

        bytes memory arr = new bytes(bs.length);

        for (uint i =0; i < bs.length;i++) {
            arr[i]= bs[bs.length-1-i];
        }


        return string(arr);
    }
}
