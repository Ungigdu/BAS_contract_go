package BAS_Ethereum

import (
	"github.com/Ungigdu/BAS_contract_go/contracts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
)

var (
	RopstenNetworkAccessPoint = "https://ropsten.infura.io/v3/8b8db3cca50a4fcf97173b7619b1c4c3"
	BASTokenAddress           = "0x3058A7Ed6a0E15691F9e309cbe982A820928e055"
	BASManagerSimpleAddress         = "0x8Cd10021DA4fDb92DdC5dd063b89B087125D7e19"
)

var Token *contracts.BASToken
var Manager *contracts.BASManagerSimple

type DomainRecord struct {
	keyHash [32]byte
	IPv4    [4]byte
	IPv6 	[16]byte
	BCLength [1]byte
	BCAddress [32]byte
	EthAddress common.Address
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
	if manager,err:=contracts.NewBASManagerSimple(common.HexToAddress(BASManagerSimpleAddress),conn);err!=nil{
		panic(err)
	}else{
		Manager = manager
	}
}

func openWallet(cipherKey,password string) (*bind.TransactOpts, error){
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
func ApproveToken(cipherKey,password,key string,ethAccount common.Address,value *big.Int) (string,error){
	auth,err:= openWallet(cipherKey,password)
	if err!=nil{
		return "",err
	}
	return buildTxResponse(Token.Approve(auth,ethAccount,value))
}

func Rent(cipherKey,password,key string, year uint8 ,data []byte)(string,error){
	auth,err:= openWallet(cipherKey,password)
	if err!=nil{
		return "",err
	}
	return buildTxResponse(Manager.Rent(auth,key,year,data))
}

func Change(cipherKey,password,key string,data []byte)(string,error){
	auth,err:= openWallet(cipherKey,password)
	if err!=nil{
		return "",err
	}
	return buildTxResponse(Manager.Change(auth,key,data))
}

func buildQueryResult(h [32]byte,f [4]byte, s [16]byte, l [1]byte, b [32]byte,a common.Address,err error) (DomainRecord,error){
	if err!=nil{
		return DomainRecord{},err
	}
	return DomainRecord{
		keyHash:   h,
		IPv4:      f,
		IPv6:      s,
		BCLength:  l,
		BCAddress: b,
		EthAddress:a,
	},nil
}

func QueryByString(key string) (DomainRecord,error){
	return buildQueryResult(Manager.QueryByString(nil,key))
}


func QueryByBCAddress(bcAddress [32]byte)(DomainRecord,error)  {
	return buildQueryResult(Manager.QueryByBCAddress(nil,bcAddress))
}