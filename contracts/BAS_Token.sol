pragma solidity >=0.4.24;

import "./ERC20.sol";

contract BAS_Token is ERC20{

    string  public constant  name = "blockchain address system";
    string  public constant  symbol = "BAS";
    uint8   public constant  decimals = 18;
    uint256 public constant INITIAL_SUPPLY = 4.2e8 * (10 ** uint256(decimals));

    constructor() public{
        _mint(msg.sender, INITIAL_SUPPLY);
    }
}