pragma solidity >=0.4.24;

import "./owned.sol";
import "./safemath.sol";
import "./ERC20.sol";

contract BAS_Manager_Simple is owned{
    using SafeMath for uint256;
    ERC20 public token;

    mapping (bytes32=>Record) public DataRecords;
    mapping (bytes32=>bytes32) private BCAddressMapping;
    uint public price = 100 * 10**18;

    struct Record{
        bytes4 IPv4;
        bytes16 IPv6;
        bytes1 BCLength;
        bytes32 BCAddress;
        address EthAddress;
    }
    constructor (address tokenAddr) public{
        token = ERC20(tokenAddr);
    }
    function changePrice(uint16 newPrice) public onlyOwner  returns(uint){
        price = newPrice;
        return price;
    }
    function checkRent(uint8 y) public view returns(uint256) {
        return y*price;
    }
    function checkAllowance() public view returns(uint256){
        return token.allowance(msg.sender,address(this));
    }
    function executeAllowanceBySender(uint256 value) internal{
        token.transferFrom(msg.sender,address(this),value);
    }
    function rent(string memory key,uint8 y,bytes memory data) public{
        uint256 r = checkRent(y);
        require(checkAllowance()>=r,"not enough token");
        executeAllowanceBySender(r);
        insert(key,data);
    }
    function split(bytes memory data) internal pure returns (bytes4 f, bytes16 s, bytes1 l, bytes32 b){
        assembly{
            f := mload(add(data,32))
            s := mload(add(data,36))
            l := mload(add(data,52))
            b := mload(add(data,53))
        }
    }
    function hash (string memory key) public pure returns (bytes32){
        return keccak256(abi.encodePacked(key));
    }
    function checkAvailablity(bytes32 hashKey,bytes4 f, bytes16 s, bytes32 b) internal view returns (bool){
        if(extractOwner(hashKey)==address(0) && BCAddressMapping[b]==0){
            return true;
        }else{
            return false;
        }
    }
    function extractOwner(bytes32 hashKey) internal view returns (address){
        Record memory r = DataRecords[hashKey];
        return r.EthAddress;
    }
    function forget(bytes4 f, bytes16 s,bytes32 b) internal{
        BCAddressMapping[b] = 0;
    }
    function remember(bytes32 hashKey,bytes4 f, bytes16 s, bytes1 l,bytes32 b) internal{
        Record memory r = Record({IPv4:f,IPv6:s,BCLength:l,BCAddress:b,EthAddress:msg.sender});
        DataRecords[hashKey] = r;
        if(b!=0){
            BCAddressMapping[b] = hashKey;
        }
    }
    function insert (string memory key, bytes memory data) internal{
        require(data.length==53,"bad format");
        bytes32 hashKey = hash(key);
        (bytes4 f, bytes16 s,bytes1 l, bytes32 b) = split(data);
        require(checkAvailablity(hashKey, f, s, b),"not available");
        remember(hashKey,f,s,l,b);
    }
    function change (string memory key, bytes memory data) public{
        require(data.length==53,"bad format");
        bytes32 hashKey = hash(key);
        Record memory old = DataRecords[hashKey];
        forget(old.IPv4,old.IPv6,old.BCAddress);
        (bytes4 f, bytes16 s, bytes1 l,bytes32 b) = split(data);
        require(extractOwner(hashKey)==msg.sender,"not owned");
        remember(hashKey,f,s,l,b);
    }
    function queryByHash(bytes32 hashKey) public view returns (bytes32,bytes4,bytes16,bytes1,bytes32,address){
        Record memory r = DataRecords[hashKey];
        return (hashKey,r.IPv4,r.IPv6,r.BCLength,r.BCAddress,r.EthAddress);
    }
    function queryByString(string memory key)public view returns (bytes32,bytes4,bytes16,bytes1,bytes32,address){
        bytes32 hashKey = hash(key);
        return queryByHash(hashKey);
    }
    function queryByBCAddress(bytes32 bc) public view returns (bytes32,bytes4,bytes16,bytes1,bytes32,address) {
        bytes32 hashKey = BCAddressMapping[bc];
        return queryByHash(hashKey);
    }
}