package main

import (
	"fmt"
	"github.com/Ungigdu/BAS_contract_go/BAS_Ethereum"
)

func main(){
	BAS_Ethereum.RecoverContract()
	fmt.Println(BAS_Ethereum.QueryByString("nbs"))
	fmt.Println(BAS_Ethereum.QueryByIPv4([4]byte{45,77,222,59}))
	fmt.Println(BAS_Ethereum.QueryByBCAddress([32]byte{93,52,39,43,194,110,93,110,71,95,191,37,246,56,50,176,47,37,251,42,0,0,0,0,0,0,0,0,0,0,0,0}))
}
