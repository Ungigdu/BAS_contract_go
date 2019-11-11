package BAS_Ethereum

import (
	"BAS_contract_interface/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
)

const (
	RopstenNetworkAccessPoint = "https://ropsten.infura.io/v3/8b8db3cca50a4fcf97173b7619b1c4c3"
	BASTokenAddress           = "0x3058A7Ed6a0E15691F9e309cbe982A820928e055"
	BASManagerAddress         = "0xf3f43AB2A248846D5C320eFE8dD580141F8B1792"
)

var Token *contracts.BASToken
var Manager *contracts.BASManager

type DomainRecord struct {
	keyHash [32]byte
	IPv4    [4]byte
	IPv6 	[16]byte
	BCAddress [32]byte
}

func RecoverContract(){
	conn, err := ethclient.Dial(RopstenNetworkAccessPoint)
	if err != nil {
		panic(err)
	}
	if token,err:=contracts.NewBASToken(common.HexToAddress(BASTokenAddress),conn); err!=nil{
		panic(err)
	}else{
		Token = token
	}
	if manager,err:=contracts.NewBASManager(common.HexToAddress(BASManagerAddress),conn);err!=nil{
		panic(err)
	}else{
		Manager = manager
	}
}

func OpenWallet(cipherKey,password string) (*bind.TransactOpts, error){
	auth, err := bind.NewTransactor(strings.NewReader(cipherKey), password)
	if err != nil {
		return nil, err
	}
	return  auth, nil
}

func buildTxResponse(tx *types.Transaction,err error)(string,error){
	if err!=nil{
		return "",err
	}else{
		return tx.Hash().String(),nil
	}
}
func ApproveToken(ethAccount common.Address,value *big.Int) (string,error){
	return buildTxResponse(Token.Approve(nil,ethAccount,value))
}

func Rent(cipherKey,password,key string, year uint8 ,data []byte)(string,error){
	auth,err:=OpenWallet(cipherKey,password)
	if err!=nil{
		return "",err
	}
	return buildTxResponse(Manager.Rent(auth,key,year,data))
}

func Change(cipherKey,password,key string,data []byte)(string,error){
	auth,err:=OpenWallet(cipherKey,password)
	if err!=nil{
		return "",err
	}
	return buildTxResponse(Manager.Change(auth,key,data))
}

func buildQueryResult(h [32]byte,f [4]byte, s [16]byte, b [32]byte,err error) (DomainRecord,error){
	if err!=nil{
		return DomainRecord{},err
	}
	return DomainRecord{
		keyHash:   h,
		IPv4:      f,
		IPv6:      s,
		BCAddress: b,
	},nil
}

func QueryByString(key string) (DomainRecord,error){
	return buildQueryResult(Manager.QueryByString(nil,key))
}

func QueryByIPv4(IPv4 [4]byte) (DomainRecord,error){
	return buildQueryResult(Manager.QueryByIPv4(nil,IPv4))
}

func QueryByIPv6(IPv6 [16]byte) (DomainRecord,error){
	return buildQueryResult(Manager.QueryByIPv6(nil,IPv6))
}

func QueryByBCAddress(bcAddress [32]byte)(DomainRecord,error)  {
	return buildQueryResult(Manager.QueryByBCAddress(nil,bcAddress))
}