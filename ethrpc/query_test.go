package ethrpc

import (
	"fmt"
	"testing"
)

var ()

//
//func TestEthRpc_BlockByNumber(t *testing.T) {
//	d,e := q.BlockByNumber(big.NewInt(13000000))
//	if e != nil {
//		t.Fatal("block by number err:", e)
//	} else {
//		t.Log("get block d:", d)
//	}
//}
//
//func TestEthRpc_BlockNumber(t *testing.T) {
//	d,e := q.BlockNumber()
//	if e != nil {
//		t.Fatal("block number err:", e)
//	} else {
//		t.Log("get blocknumber ", d)
//	}
//}
//
//func TestEthRpc_BlockTransactionCountByNumber(t *testing.T) {
//	d,e := q.BlockTransactionCountByNumber(big.NewInt(13884904))
//	if e != nil {
//		t.Fatal("txcount by block number err:", e)
//	} else {
//		t.Log("get txcount by block number d:", d)
//	}
//}

//func TestEthRpc_TransactionByHash(t *testing.T) {
//	d,e := q.TransactionByHash(common.HexToHash("0x5d8703a18fc2028ab9a0cb8176bbb6c7113acd497ebe6c0d3dc7a8436415fc7d"))
//	if e != nil {
//		t.Fatal("transaction by hash failed, e:", e)
//	} else {
//		t.Log("get transaction by hash got ", d)
//	}
//}

//func TestEthRpc_TransactionByBlockNumberAndIndex(t *testing.T) {
//	d,e := q.TransactionByBlockNumberAndIndex(big.NewInt(13884904), 25)
//	if e != nil {
//		t.Fatal("TransactionByBlockNumberAndIndex by hash failed, e:", e)
//	} else {
//		t.Log("get TransactionByBlockNumberAndIndex by hash got ", d)
//	}
//}
//
//func TestEthRpc_TransactionCount(t *testing.T) {
//	addr := common.HexToAddress("0x0D29168B39f9bE8AaAb6ac3Cae1C11548e6ff9C6")
//	d,e := q.TransactionCount(addr, nil)
//	if e != nil {
//		t.Fatal("TransactionCount by hash failed, e:", e)
//	} else {
//		t.Log("get TransactionCount by hash got ", d)
//	}
//}
//
//func TestEthRpc_SendRawTransaction(t *testing.T) {
//	// todo: add raw transaction.
//	raw := common.Hex2Bytes("")
//	d,e := q.SendRawTransaction(raw)
//	if e != nil {
//		t.Fatal("SendRawTransaction by hash failed, e:", e)
//	} else {
//		t.Log("SendRawTransaction got ", d)
//	}
//}

//func TestEthRpc_TransactionReceipt(t *testing.T) {
//	hash := common.HexToHash("0x5d8703a18fc2028ab9a0cb8176bbb6c7113acd497ebe6c0d3dc7a8436415fc7d")
//	d,e := q.TransactionReceipt(hash)
//	if e != nil {
//		t.Fatal("TransactionReceipt failed, e:", e)
//	} else {
//		t.Log("get TransactionReceipt got ", d)
//	}
//}
//
//func TestEthRpc_CallContract(t *testing.T) {
//	to := common.HexToAddress("0xbe05ac1fb417c9ea435b37a9cecd39bc70359d31")
//	data := common.FromHex("0x18160ddd") // total supply
//	msg := ethereum.CallMsg{}
//	msg.To = &to
//	msg.Data = data
//	msg.Gas = 10000000
//	d,e := q.CallContract(msg,nil)
//	if e != nil {
//		t.Fatal("TransactionReceipt failed, e:", e)
//	} else {
//		t.Log("get TransactionReceipt got ", d)
//	}
//}
//
//
//func TestEthRpc_GetCode(t *testing.T) {
//	addr := common.HexToAddress("0xbe05ac1fb417c9ea435b37a9cecd39bc70359d31")
//	d,e := q.GetCode(addr,nil)
//	if e != nil {
//		t.Fatal("GetCode failed, e:", e)
//	} else {
//		t.Log("get GetCode got ", d)
//	}
//}
//
//func TestEthRpc_StorageAt(t *testing.T) {
//	addr := common.HexToAddress("0xbe05ac1fb417c9ea435b37a9cecd39bc70359d31")
//	pos := common.HexToHash("1")
//	fmt.Println("pos is ", pos.String())
//	d,e := q.StorageAt(addr, pos,nil)
//	if e != nil {
//		t.Fatal("StorageAt failed, e:", e)
//	} else {
//		fmt.Println("get StorageAt got ", d)
//	}
//}

func TestEthRpc_GasPrice(t *testing.T) {
	d, e := GasPrice()
	if e != nil {
		t.Fatal("GasPrice failed, e:", e)
	} else {
		fmt.Println("get GasPrice got ", d)
	}
}
