package test

import (
	"fmt"
	"github.com/AsimovNetwork/asimov/blockchain"
	"github.com/AsimovNetwork/asimov/chaincfg"
	"github.com/AsimovNetwork/asimov/common"
	"github.com/AsimovNetwork/asimov/database/dbimpl/ethdb"
	"github.com/AsimovNetwork/asimov/protos"
	"github.com/AsimovNetwork/asimov/vm/fvm"
	"github.com/AsimovNetwork/asimov/vm/fvm/abi"
	"github.com/AsimovNetwork/asimov/vm/fvm/core/state"
	"github.com/AsimovNetwork/asimov/vm/fvm/core/vm"
	"math/big"
	"strings"
	"testing"
	"time"
)

func TestUnPackResult(t *testing.T) {
	header := protos.BlockHeader{
		Height:    1,
		Timestamp: time.Unix(time.Now().Unix(), 0),
	}

	// Account address
	address := common.HexToAddress("0x948ab52cc7b5107efd4b03a51f0d1688b4a49a54")

	// FVM Context
	// bc := newBlockChain()
	bc := &blockchain.BlockChain{}
	context := fvm.NewFVMContext(address, new(big.Int).SetInt64(1), &header, bc, nil)

	db := ethdb.NewMemDatabase()
	stateDB, _ := state.New(common.Hash{}, state.NewDatabase(db))

	// FVM
	vmenv := vm.NewFVM(context, stateDB, chaincfg.ActiveNetParams.FvmParam, *bc.GetVmConfig())

	// contract creator
	sender := vm.AccountRef(address)
	// code := common.Hex2Bytes("608060405234801561001057600080fd5b506040516102a83803806102a8833981018060405281019080805182019291905050508060009080519060200190610049929190610050565b50506100f5565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061009157805160ff19168380011785556100bf565b828001600101855582156100bf579182015b828111156100be5782518255916020019190600101906100a3565b5b5090506100cc91906100d0565b5090565b6100f291905b808211156100ee5760008160009055506001016100d6565b5090565b90565b6101a4806101046000396000f300608060405260043610610041576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063954ab4b214610046575b600080fd5b34801561005257600080fd5b5061005b6100d6565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561009b578082015181840152602081019050610080565b50505050905090810190601f1680156100c85780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b606060008054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561016e5780601f106101435761010080835404028352916020019161016e565b820191906000526020600020905b81548152906001019060200180831161015157829003601f168201915b50505050509050905600a165627a7a723058203301ef54e16b133c1d5ccfec6c93620b1d958eed02ca86302c21b972c62205700029")
	// ode += construct
	/*
		how to encode
		1. install rust：curl https://sh.rustup.rs -sSf | sh
		2. install ethabi：cargo install ethabi-cli
		3. ethabi encode params -v string "Hello World"
	*/
	code := common.Hex2Bytes("60806040523480156200001157600080fd5b506040516200100b3803806200100b8339810180604052810190808051820192919060200180518201929190505050816000908051906020019062000058929190620001de565b50806001908051906020019062000071929190620001de565b506000604051908082528060200260200182016040528015620000a957816020015b6060815260200190600190039081620000935790505b5060029080519060200190620000c192919062000265565b506040805190810160405280600981526020017f7a68616e676a696e670000000000000000000000000000000000000000000000815250600360405180807f6e616d65000000000000000000000000000000000000000000000000000000008152506004019050908152602001604051809103902090805190602001906200014b929190620001de565b506040805190810160405280600281526020017f3332000000000000000000000000000000000000000000000000000000000000815250600360405180807f6167650000000000000000000000000000000000000000000000000000000000815250600301905090815260200160405180910390209080519060200190620001d5929190620001de565b505050620003f8565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200022157805160ff191683800117855562000252565b8280016001018555821562000252579182015b828111156200025157825182559160200191906001019062000234565b5b509050620002619190620002cc565b5090565b828054828255906000526020600020908101928215620002b9579160200282015b82811115620002b8578251829080519060200190620002a7929190620002f4565b509160200191906001019062000286565b5b509050620002c891906200037b565b5090565b620002f191905b80821115620002ed576000816000905550600101620002d3565b5090565b90565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200033757805160ff191683800117855562000368565b8280016001018555821562000368579182015b82811115620003675782518255916020019190600101906200034a565b5b509050620003779190620002cc565b5090565b620003a991905b80821115620003a557600081816200039b9190620003ac565b5060010162000382565b5090565b90565b50805460018160011615610100020316600290046000825580601f10620003d45750620003f5565b601f016020900490600052602060002090810190620003f49190620002cc565b5b50565b610c0380620004086000396000f3006080604052600436106100a4576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806303942978146100a957806331543cf4146101125780633350fe42146101c35780633cf0a6e71461022c5780634622ab03146102dc5780634c21eb071461038257806360f96a8f146103eb5780639508484b1461047b578063954ab4b2146104bd578063ef690cc01461054d575b600080fd5b3480156100b557600080fd5b50610110600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506105dd565b005b34801561011e57600080fd5b5061013d60048036038101908080359060200190929190505050610719565b604051808060200183151515158152602001828103825284818151815260200191508051906020019080838360005b8381101561018757808201518184015260208101905061016c565b50505050905090810190601f1680156101b45780820380516001836020036101000a031916815260200191505b50935050505060405180910390f35b3480156101cf57600080fd5b5061022a600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610739565b005b34801561023857600080fd5b506102416107b3565b604051808060200184600381111561025557fe5b60ff16815260200183151515158152602001828103825285818151815260200191508051906020019080838360005b8381101561029f578082015181840152602081019050610284565b50505050905090810190601f1680156102cc5780820380516001836020036101000a031916815260200191505b5094505050505060405180910390f35b3480156102e857600080fd5b50610307600480360381019080803590602001909291905050506107fe565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561034757808201518184015260208101905061032c565b50505050905090810190601f1680156103745780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561038e57600080fd5b506103e9600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506108b9565b005b3480156103f757600080fd5b506104006108fb565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610440578082015181840152602081019050610425565b50505050905090810190601f16801561046d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561048757600080fd5b50610490610999565b604051808363ffffffff1663ffffffff168152602001821515151581526020019250505060405180910390f35b3480156104c957600080fd5b506104d26109aa565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156105125780820151818401526020810190506104f7565b50505050905090810190601f16801561053f5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561055957600080fd5b50610562610a4c565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156105a2578082015181840152602081019050610587565b50505050905090810190601f1680156105cf5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b60008090505b60028054905081101561071557816040518082805190602001908083835b6020831015156106265780518252602082019150602081019050602083039250610601565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390206000191660028281548110151561066657fe5b9060005260206000200160405180828054600181600116156101000203166002900480156106cb5780601f106106a95761010080835404028352918201916106cb565b820191906000526020600020905b8154815290600101906020018083116106b7575b50509150506040518091039020600019161415610708576002818154811015156106f157fe5b9060005260206000200160006107079190610aea565b5b80806001019150506105e3565b5050565b606060008060206040519081016040528060008152509091509150915091565b6003816040518082805190602001908083835b602083101515610771578051825260208201915060208101905060208303925061074c565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060006107b09190610aea565b50565b60606000806001806040805190810160405280600981526020017f7a68616e676a696e6700000000000000000000000000000000000000000000008152509190925092509250909192565b60028181548110151561080d57fe5b906000526020600020016000915090508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108b15780601f10610886576101008083540402835291602001916108b1565b820191906000526020600020905b81548152906001019060200180831161089457829003601f168201915b505050505081565b60028190806001815401808255809150509060018203906000526020600020016000909192909190915090805190602001906108f6929190610b32565b505050565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109915780601f1061096657610100808354040283529160200191610991565b820191906000526020600020905b81548152906001019060200180831161097457829003601f168201915b505050505081565b600080600080819150915091509091565b606060008054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a425780601f10610a1757610100808354040283529160200191610a42565b820191906000526020600020905b815481529060010190602001808311610a2557829003601f168201915b5050505050905090565b60008054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ae25780601f10610ab757610100808354040283529160200191610ae2565b820191906000526020600020905b815481529060010190602001808311610ac557829003601f168201915b505050505081565b50805460018160011615610100020316600290046000825580601f10610b105750610b2f565b601f016020900490600052602060002090810190610b2e9190610bb2565b5b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610b7357805160ff1916838001178555610ba1565b82800160010185558215610ba1579182015b82811115610ba0578251825591602001919060010190610b85565b5b509050610bae9190610bb2565b5090565b610bd491905b80821115610bd0576000816000905550600101610bb8565b5090565b905600a165627a7a723058204408f295a5e0bd74ceda73bb1122a71918abbd0d6cfdbdfd17e9fcecebf46ecd0029000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000800000000000000000000000000000000000000000000000000000000000000006616466616466000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000066164666164660000000000000000000000000000000000000000000000000000")
	res, addr, returnGas, _, err := vmenv.Create(
		sender,
		code,
		uint64(4604216),
		common.Big0,
		nil,
		nil,
	)

	if err != nil {
		panic(err)
	}

	fmt.Println(res)
	fmt.Println(addr)
	fmt.Println(returnGas)

	callCode := common.Hex2Bytes("3cf0a6e7")
	ret, leftGas, err := vmenv.Call(sender, addr, callCode, uint64(999999999), common.Big0, nil)
	fmt.Println(leftGas)
	fmt.Println("Contract's result is : ")

	abiStr := `[
	{
		"constant": false,
		"inputs": [
			{
				"name": "_name",
				"type": "string"
			}
		],
		"name": "deleteArray",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "_key",
				"type": "string"
			}
		],
		"name": "deleteKey",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"constant": false,
		"inputs": [
			{
				"name": "_name",
				"type": "string"
			}
		],
		"name": "put",
		"outputs": [],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "function"
	},
	{
		"inputs": [
			{
				"name": "_greeting",
				"type": "string"
			},
			{
				"name": "_parent",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "nonpayable",
		"type": "constructor"
	},
	{
		"constant": true,
		"inputs": [
			{
				"name": "index",
				"type": "uint256"
			}
		],
		"name": "getTemplate",
		"outputs": [
			{
				"name": "",
				"type": "string"
			},
			{
				"name": "",
				"type": "bool"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "greeting",
		"outputs": [
			{
				"name": "",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [
			{
				"name": "",
				"type": "uint256"
			}
		],
		"name": "names",
		"outputs": [
			{
				"name": "",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "parent",
		"outputs": [
			{
				"name": "",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "say",
		"outputs": [
			{
				"name": "",
				"type": "string"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "testReturnTemplate",
		"outputs": [
			{
				"name": "",
				"type": "string"
			},
			{
				"name": "",
				"type": "uint8"
			},
			{
				"name": "",
				"type": "bool"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	},
	{
		"constant": true,
		"inputs": [],
		"name": "testReturnValue",
		"outputs": [
			{
				"name": "",
				"type": "uint32"
			},
			{
				"name": "",
				"type": "bool"
			}
		],
		"payable": false,
		"stateMutability": "view",
		"type": "function"
	}
]`
	definition, err := abi.JSON(strings.NewReader(abiStr))
	outType := &[]interface{}{new(string), new(uint8), new(bool)}
	err = definition.Unpack(outType, "testReturnTemplate", ret)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(*((*outType)[0]).(*string))
	fmt.Println(*((*outType)[1]).(*uint8))
	fmt.Println(*((*outType)[2]).(*bool))
}
