// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package GenericHandler

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// GenericHandlerDepositRecord is an auto generated low-level Go binding around an user-defined struct.
type GenericHandlerDepositRecord struct {
	DestinationChainID uint8
	Depositer          common.Address
	ResourceID         [32]byte
	MetaData           []byte
}

// GenericHandlerABI is the input ABI used to generate the binding from.
const GenericHandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"initialResourceIDs\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"initialContractAddresses\",\"type\":\"address[]\"},{\"internalType\":\"bytes4[]\",\"name\":\"initialDepositFunctionSignatures\",\"type\":\"bytes4[]\"},{\"internalType\":\"bytes4[]\",\"name\":\"initialExecuteFunctionSignatures\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToDepositFunctionSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToExecuteFunctionSignature\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address[36]\",\"name\":\"signers\",\"type\":\"address[36]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"destId\",\"type\":\"uint8\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_metaData\",\"type\":\"bytes\"}],\"internalType\":\"structGenericHandler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"depositFunctionSig\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"executeFunctionSig\",\"type\":\"bytes4\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// GenericHandlerBin is the compiled bytecode used for deploying new contracts.
var GenericHandlerBin = "0x60806040523480156200001157600080fd5b50604051620022563803806200225683398181016040528101906200003791906200056d565b82518451146200007e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000075906200078b565b60405180910390fd5b8151835114620000c5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620000bc90620007cf565b60405180910390fd5b80518251146200010c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200010390620007ad565b60405180910390fd5b846000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060005b8451811015620001cb57620001bd8582815181106200016a57fe5b60200260200101518583815181106200017f57fe5b60200260200101518584815181106200019457fe5b6020026020010151858581518110620001a957fe5b6020026020010151620001d760201b60201c565b80806001019150506200014f565b50505050505062000963565b826002600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908360e01c021790555080600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908360e01c02179055506001600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050505050565b600081519050620003928162000915565b92915050565b600082601f830112620003aa57600080fd5b8151620003c1620003bb826200081f565b620007f1565b91508181835260208401935060208101905083856020840282011115620003e757600080fd5b60005b838110156200041b578162000400888262000381565b845260208401935060208301925050600181019050620003ea565b5050505092915050565b600082601f8301126200043757600080fd5b81516200044e620004488262000848565b620007f1565b915081818352602084019350602081019050838560208402820111156200047457600080fd5b60005b83811015620004a857816200048d88826200053f565b84526020840193506020830192505060018101905062000477565b5050505092915050565b600082601f830112620004c457600080fd5b8151620004db620004d58262000871565b620007f1565b915081818352602084019350602081019050838560208402820111156200050157600080fd5b60005b838110156200053557816200051a888262000556565b84526020840193506020830192505060018101905062000504565b5050505092915050565b60008151905062000550816200092f565b92915050565b600081519050620005678162000949565b92915050565b600080600080600060a086880312156200058657600080fd5b6000620005968882890162000381565b955050602086015167ffffffffffffffff811115620005b457600080fd5b620005c28882890162000425565b945050604086015167ffffffffffffffff811115620005e057600080fd5b620005ee8882890162000398565b935050606086015167ffffffffffffffff8111156200060c57600080fd5b6200061a88828901620004b2565b925050608086015167ffffffffffffffff8111156200063857600080fd5b6200064688828901620004b2565b9150509295509295909350565b600062000662603c836200089a565b91507f696e697469616c5265736f7572636549447320616e6420696e697469616c436f60008301527f6e7472616374416464726573736573206c656e206d69736d61746368000000006020830152604082019050919050565b6000620006ca603d836200089a565b91507f70726f7669646564206465706f73697420616e6420657865637574652066756e60008301527f6374696f6e207369676e617475726573206c656e206d69736d617463680000006020830152604082019050919050565b6000620007326040836200089a565b91507f70726f766964656420636f6e74726163742061646472657373657320616e642060008301527f66756e6374696f6e207369676e617475726573206c656e206d69736d617463686020830152604082019050919050565b60006020820190508181036000830152620007a68162000653565b9050919050565b60006020820190508181036000830152620007c881620006bb565b9050919050565b60006020820190508181036000830152620007ea8162000723565b9050919050565b6000604051905081810181811067ffffffffffffffff821117156200081557600080fd5b8060405250919050565b600067ffffffffffffffff8211156200083757600080fd5b602082029050602081019050919050565b600067ffffffffffffffff8211156200086057600080fd5b602082029050602081019050919050565b600067ffffffffffffffff8211156200088957600080fd5b602082029050602081019050919050565b600082825260208201905092915050565b6000620008b882620008f5565b9050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6200092081620008ab565b81146200092c57600080fd5b50565b6200093a81620008bf565b81146200094657600080fd5b50565b6200095481620008c9565b81146200096057600080fd5b50565b6118e380620009736000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063a5c3a98511610071578063a5c3a9851461017e578063ba484c09146101ae578063bba8185a146101de578063c54c2a11146101fa578063cb6244631461022a578063ec97d3b41461025a576100a9565b8063138f0b8b146100ae578063318c136e146100e157806338995da9146100ff5780634402027f1461011b5780637f79bea81461014e575b600080fd5b6100c860048036038101906100c391906110cf565b61028a565b6040516100d8949392919061156c565b60405180910390f35b6100e9610504565b6040516100f69190611536565b60405180910390f35b6101196004803603810190610114919061113d565b610528565b005b6101356004803603810190610130919061120b565b610899565b6040516101459493929190611669565b60405180910390f35b6101686004803603810190610163919061101a565b61099b565b6040516101759190611551565b60405180910390f35b6101986004803603810190610193919061101a565b6109bb565b6040516101a591906115cc565b60405180910390f35b6101c860048036038101906101c391906111cf565b6109db565b6040516101d59190611647565b60405180910390f35b6101f860048036038101906101f3919061106c565b610b53565b005b610214600480360381019061020f9190611043565b610b6d565b6040516102219190611536565b60405180910390f35b610244600480360381019061023f919061101a565b610ba0565b60405161025191906115cc565b60405180910390f35b610274600480360381019061026f919061101a565b610bc0565b60405161028191906115b1565b60405180910390f35b600080600080610298610bd8565b606060405190506064358082016060016040526064360360648337506000600260008b815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600660008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610378576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161036f90611627565b60405180910390fd5b6000600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460e01b9050600060e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916146104e857606081846040516020016104279291906114f7565b604051602081830303815290604052905060008373ffffffffffffffffffffffffffffffffffffffff168260405161045f919061151f565b6000604051808303816000865af19150503d806000811461049c576040519150601f19603f3d011682016040523d82523d6000602084013e6104a1565b606091505b50509050806104e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104dc90611607565b60405180910390fd5b50505b6000808360009650965096509650505050945094509450949050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b610530610bd8565b6000606060c4359150604051905081810160200160405260c4360360c482376000600260008a815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600660008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610613576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060a90611627565b60405180910390fd5b6000600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460e01b9050600060e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161461078357606081846040516020016106c29291906114f7565b604051602081830303815290604052905060008373ffffffffffffffffffffffffffffffffffffffff16826040516106fa919061151f565b6000604051808303816000865af19150503d8060008114610737576040519150601f19603f3d011682016040523d82523d6000602084013e61073c565b606091505b5050905080610780576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161077790611607565b60405180910390fd5b50505b60405180608001604052808a60ff1681526020018873ffffffffffffffffffffffffffffffffffffffff1681526020018b815260200184815250600160008b60ff1660ff16815260200190815260200160002060008a67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548160ff021916908360ff16021790555060208201518160000160016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604082015181600101556060820151816002019080519060200190610889929190610e12565b5090505050505050505050505050565b6001602052816000526040600020602052806000526040600020600091509150508060000160009054906101000a900460ff16908060000160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690806001015490806002018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109915780601f1061096657610100808354040283529160200191610991565b820191906000526020600020905b81548152906001019060200180831161097457829003601f168201915b5050505050905084565b60066020528060005260406000206000915054906101000a900460ff1681565b60056020528060005260406000206000915054906101000a900460e01b81565b6109e3610e92565b600160008360ff1660ff16815260200190815260200160002060008467ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206040518060800160405290816000820160009054906101000a900460ff1660ff1660ff1681526020016000820160019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200160018201548152602001600282018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b425780601f10610b1757610100808354040283529160200191610b42565b820191906000526020600020905b815481529060010190602001808311610b2557829003601f168201915b505050505081525050905092915050565b610b5b610bd8565b610b6784848484610c68565b50505050565b60026020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60046020528060005260406000206000915054906101000a900460e01b81565b60036020528060005260406000206000915090505481565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610c66576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c5d906115e7565b60405180910390fd5b565b826002600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600360008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208190555081600460008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908360e01c021790555080600560008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548163ffffffff021916908360e01c02179055506001600660008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610e5357805160ff1916838001178555610e81565b82800160010185558215610e81579182015b82811115610e80578251825591602001919060010190610e65565b5b509050610e8e9190610ed6565b5090565b6040518060800160405280600060ff168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160008019168152602001606081525090565b5b80821115610eef576000816000905550600101610ed7565b5090565b600081359050610f028161183a565b92915050565b600082601f830112610f1957600080fd5b6024610f2c610f27826116e2565b6116b5565b91508183856020840282011115610f4257600080fd5b60005b83811015610f725781610f588882610ef3565b845260208401935060208301925050600181019050610f45565b5050505092915050565b600081359050610f8b81611851565b92915050565b600081359050610fa081611868565b92915050565b60008083601f840112610fb857600080fd5b8235905067ffffffffffffffff811115610fd157600080fd5b602083019150836001820283011115610fe957600080fd5b9250929050565b600081359050610fff8161187f565b92915050565b60008135905061101481611896565b92915050565b60006020828403121561102c57600080fd5b600061103a84828501610ef3565b91505092915050565b60006020828403121561105557600080fd5b600061106384828501610f7c565b91505092915050565b6000806000806080858703121561108257600080fd5b600061109087828801610f7c565b94505060206110a187828801610ef3565b93505060406110b287828801610f91565b92505060606110c387828801610f91565b91505092959194509250565b6000806000806104c085870312156110e657600080fd5b60006110f487828801610f7c565b945050602061110587828801610f08565b9350506104a085013567ffffffffffffffff81111561112357600080fd5b61112f87828801610fa6565b925092505092959194509250565b60008060008060008060a0878903121561115657600080fd5b600061116489828a01610f7c565b965050602061117589828a01611005565b955050604061118689828a01610ff0565b945050606061119789828a01610ef3565b935050608087013567ffffffffffffffff8111156111b457600080fd5b6111c089828a01610fa6565b92509250509295509295509295565b600080604083850312156111e257600080fd5b60006111f085828601610ff0565b925050602061120185828601611005565b9150509250929050565b6000806040838503121561121e57600080fd5b600061122c85828601611005565b925050602061123d85828601610ff0565b9150509250929050565b6112508161174d565b82525050565b61125f8161174d565b82525050565b61126e8161175f565b82525050565b61127d8161176b565b82525050565b61128c8161176b565b82525050565b61129b81611775565b82525050565b6112b26112ad82611775565b61181f565b82525050565b60006112c382611704565b6112cd818561170f565b93506112dd8185602086016117ec565b6112e681611829565b840191505092915050565b60006112fc82611704565b6113068185611720565b93506113168185602086016117ec565b61131f81611829565b840191505092915050565b600061133582611704565b61133f8185611731565b935061134f8185602086016117ec565b80840191505092915050565b6000611368601e8361173c565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b60006113a860268361173c565b91507f64656c656761746563616c6c20746f20636f6e7472616374416464726573732060008301527f6661696c656400000000000000000000000000000000000000000000000000006020830152604082019050919050565b600061140e602b8361173c565b91507f70726f766964656420636f6e747261637441646472657373206973206e6f742060008301527f77686974656c69737465640000000000000000000000000000000000000000006020830152604082019050919050565b600060808301600083015161147f60008601826114d9565b5060208301516114926020860182611247565b5060408301516114a56040860182611274565b50606083015184820360608601526114bd82826112b8565b9150508091505092915050565b6114d3816117c1565b82525050565b6114e2816117df565b82525050565b6114f1816117df565b82525050565b600061150382856112a1565b600482019150611513828461132a565b91508190509392505050565b600061152b828461132a565b915081905092915050565b600060208201905061154b6000830184611256565b92915050565b60006020820190506115666000830184611265565b92915050565b60006080820190506115816000830187611265565b61158e6020830186611265565b61159b6040830185611256565b6115a860608301846114ca565b95945050505050565b60006020820190506115c66000830184611283565b92915050565b60006020820190506115e16000830184611292565b92915050565b600060208201905081810360008301526116008161135b565b9050919050565b600060208201905081810360008301526116208161139b565b9050919050565b6000602082019050818103600083015261164081611401565b9050919050565b600060208201905081810360008301526116618184611467565b905092915050565b600060808201905061167e60008301876114e8565b61168b6020830186611256565b6116986040830185611283565b81810360608301526116aa81846112f1565b905095945050505050565b6000604051905081810181811067ffffffffffffffff821117156116d857600080fd5b8060405250919050565b600067ffffffffffffffff8211156116f957600080fd5b602082029050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b6000611758826117a1565b9050919050565b60008115159050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b600060ff82169050919050565b60005b8381101561180a5780820151818401526020810190506117ef565b83811115611819576000848401525b50505050565b6000819050919050565b6000601f19601f8301169050919050565b6118438161174d565b811461184e57600080fd5b50565b61185a8161176b565b811461186557600080fd5b50565b61187181611775565b811461187c57600080fd5b50565b611888816117cb565b811461189357600080fd5b50565b61189f816117df565b81146118aa57600080fd5b5056fea2646970667358221220aff76f55c204ba515a88176b3665bff6657e4f53488fcd2dfb84eea4005fa6ee64736f6c634300060c0033"

// DeployGenericHandler deploys a new Ethereum contract, binding an instance of GenericHandler to it.
func DeployGenericHandler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address, initialResourceIDs [][32]byte, initialContractAddresses []common.Address, initialDepositFunctionSignatures [][4]byte, initialExecuteFunctionSignatures [][4]byte) (common.Address, *types.Transaction, *GenericHandler, error) {
	parsed, err := abi.JSON(strings.NewReader(GenericHandlerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(GenericHandlerBin), backend, bridgeAddress, initialResourceIDs, initialContractAddresses, initialDepositFunctionSignatures, initialExecuteFunctionSignatures)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GenericHandler{GenericHandlerCaller: GenericHandlerCaller{contract: contract}, GenericHandlerTransactor: GenericHandlerTransactor{contract: contract}, GenericHandlerFilterer: GenericHandlerFilterer{contract: contract}}, nil
}

// GenericHandler is an auto generated Go binding around an Ethereum contract.
type GenericHandler struct {
	GenericHandlerCaller     // Read-only binding to the contract
	GenericHandlerTransactor // Write-only binding to the contract
	GenericHandlerFilterer   // Log filterer for contract events
}

// GenericHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type GenericHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GenericHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GenericHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GenericHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GenericHandlerSession struct {
	Contract     *GenericHandler   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GenericHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GenericHandlerCallerSession struct {
	Contract *GenericHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// GenericHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GenericHandlerTransactorSession struct {
	Contract     *GenericHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// GenericHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type GenericHandlerRaw struct {
	Contract *GenericHandler // Generic contract binding to access the raw methods on
}

// GenericHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GenericHandlerCallerRaw struct {
	Contract *GenericHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// GenericHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GenericHandlerTransactorRaw struct {
	Contract *GenericHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGenericHandler creates a new instance of GenericHandler, bound to a specific deployed contract.
func NewGenericHandler(address common.Address, backend bind.ContractBackend) (*GenericHandler, error) {
	contract, err := bindGenericHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GenericHandler{GenericHandlerCaller: GenericHandlerCaller{contract: contract}, GenericHandlerTransactor: GenericHandlerTransactor{contract: contract}, GenericHandlerFilterer: GenericHandlerFilterer{contract: contract}}, nil
}

// NewGenericHandlerCaller creates a new read-only instance of GenericHandler, bound to a specific deployed contract.
func NewGenericHandlerCaller(address common.Address, caller bind.ContractCaller) (*GenericHandlerCaller, error) {
	contract, err := bindGenericHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GenericHandlerCaller{contract: contract}, nil
}

// NewGenericHandlerTransactor creates a new write-only instance of GenericHandler, bound to a specific deployed contract.
func NewGenericHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*GenericHandlerTransactor, error) {
	contract, err := bindGenericHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GenericHandlerTransactor{contract: contract}, nil
}

// NewGenericHandlerFilterer creates a new log filterer instance of GenericHandler, bound to a specific deployed contract.
func NewGenericHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*GenericHandlerFilterer, error) {
	contract, err := bindGenericHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GenericHandlerFilterer{contract: contract}, nil
}

// bindGenericHandler binds a generic wrapper to an already deployed contract.
func bindGenericHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GenericHandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenericHandler *GenericHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GenericHandler.Contract.GenericHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenericHandler *GenericHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericHandler.Contract.GenericHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenericHandler *GenericHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenericHandler.Contract.GenericHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GenericHandler *GenericHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GenericHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GenericHandler *GenericHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GenericHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GenericHandler *GenericHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GenericHandler.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_GenericHandler *GenericHandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_GenericHandler *GenericHandlerSession) BridgeAddress() (common.Address, error) {
	return _GenericHandler.Contract.BridgeAddress(&_GenericHandler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_GenericHandler *GenericHandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _GenericHandler.Contract.BridgeAddress(&_GenericHandler.CallOpts)
}

// ContractAddressToDepositFunctionSignature is a free data retrieval call binding the contract method 0xcb624463.
//
// Solidity: function _contractAddressToDepositFunctionSignature(address ) view returns(bytes4)
func (_GenericHandler *GenericHandlerCaller) ContractAddressToDepositFunctionSignature(opts *bind.CallOpts, arg0 common.Address) ([4]byte, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_contractAddressToDepositFunctionSignature", arg0)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// ContractAddressToDepositFunctionSignature is a free data retrieval call binding the contract method 0xcb624463.
//
// Solidity: function _contractAddressToDepositFunctionSignature(address ) view returns(bytes4)
func (_GenericHandler *GenericHandlerSession) ContractAddressToDepositFunctionSignature(arg0 common.Address) ([4]byte, error) {
	return _GenericHandler.Contract.ContractAddressToDepositFunctionSignature(&_GenericHandler.CallOpts, arg0)
}

// ContractAddressToDepositFunctionSignature is a free data retrieval call binding the contract method 0xcb624463.
//
// Solidity: function _contractAddressToDepositFunctionSignature(address ) view returns(bytes4)
func (_GenericHandler *GenericHandlerCallerSession) ContractAddressToDepositFunctionSignature(arg0 common.Address) ([4]byte, error) {
	return _GenericHandler.Contract.ContractAddressToDepositFunctionSignature(&_GenericHandler.CallOpts, arg0)
}

// ContractAddressToExecuteFunctionSignature is a free data retrieval call binding the contract method 0xa5c3a985.
//
// Solidity: function _contractAddressToExecuteFunctionSignature(address ) view returns(bytes4)
func (_GenericHandler *GenericHandlerCaller) ContractAddressToExecuteFunctionSignature(opts *bind.CallOpts, arg0 common.Address) ([4]byte, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_contractAddressToExecuteFunctionSignature", arg0)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// ContractAddressToExecuteFunctionSignature is a free data retrieval call binding the contract method 0xa5c3a985.
//
// Solidity: function _contractAddressToExecuteFunctionSignature(address ) view returns(bytes4)
func (_GenericHandler *GenericHandlerSession) ContractAddressToExecuteFunctionSignature(arg0 common.Address) ([4]byte, error) {
	return _GenericHandler.Contract.ContractAddressToExecuteFunctionSignature(&_GenericHandler.CallOpts, arg0)
}

// ContractAddressToExecuteFunctionSignature is a free data retrieval call binding the contract method 0xa5c3a985.
//
// Solidity: function _contractAddressToExecuteFunctionSignature(address ) view returns(bytes4)
func (_GenericHandler *GenericHandlerCallerSession) ContractAddressToExecuteFunctionSignature(arg0 common.Address) ([4]byte, error) {
	return _GenericHandler.Contract.ContractAddressToExecuteFunctionSignature(&_GenericHandler.CallOpts, arg0)
}

// ContractAddressToResourceID is a free data retrieval call binding the contract method 0xec97d3b4.
//
// Solidity: function _contractAddressToResourceID(address ) view returns(bytes32)
func (_GenericHandler *GenericHandlerCaller) ContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_contractAddressToResourceID", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ContractAddressToResourceID is a free data retrieval call binding the contract method 0xec97d3b4.
//
// Solidity: function _contractAddressToResourceID(address ) view returns(bytes32)
func (_GenericHandler *GenericHandlerSession) ContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _GenericHandler.Contract.ContractAddressToResourceID(&_GenericHandler.CallOpts, arg0)
}

// ContractAddressToResourceID is a free data retrieval call binding the contract method 0xec97d3b4.
//
// Solidity: function _contractAddressToResourceID(address ) view returns(bytes32)
func (_GenericHandler *GenericHandlerCallerSession) ContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _GenericHandler.Contract.ContractAddressToResourceID(&_GenericHandler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_GenericHandler *GenericHandlerCaller) ContractWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_contractWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_GenericHandler *GenericHandlerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _GenericHandler.Contract.ContractWhitelist(&_GenericHandler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_GenericHandler *GenericHandlerCallerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _GenericHandler.Contract.ContractWhitelist(&_GenericHandler.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0x4402027f.
//
// Solidity: function _depositRecords(uint8 , uint64 ) view returns(uint8 _destinationChainID, address _depositer, bytes32 _resourceID, bytes _metaData)
func (_GenericHandler *GenericHandlerCaller) DepositRecords(opts *bind.CallOpts, arg0 uint8, arg1 uint64) (struct {
	DestinationChainID uint8
	Depositer          common.Address
	ResourceID         [32]byte
	MetaData           []byte
}, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_depositRecords", arg0, arg1)

	outstruct := new(struct {
		DestinationChainID uint8
		Depositer          common.Address
		ResourceID         [32]byte
		MetaData           []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DestinationChainID = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Depositer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.ResourceID = *abi.ConvertType(out[2], new([32]byte)).(*[32]byte)
	outstruct.MetaData = *abi.ConvertType(out[3], new([]byte)).(*[]byte)

	return *outstruct, err

}

// DepositRecords is a free data retrieval call binding the contract method 0x4402027f.
//
// Solidity: function _depositRecords(uint8 , uint64 ) view returns(uint8 _destinationChainID, address _depositer, bytes32 _resourceID, bytes _metaData)
func (_GenericHandler *GenericHandlerSession) DepositRecords(arg0 uint8, arg1 uint64) (struct {
	DestinationChainID uint8
	Depositer          common.Address
	ResourceID         [32]byte
	MetaData           []byte
}, error) {
	return _GenericHandler.Contract.DepositRecords(&_GenericHandler.CallOpts, arg0, arg1)
}

// DepositRecords is a free data retrieval call binding the contract method 0x4402027f.
//
// Solidity: function _depositRecords(uint8 , uint64 ) view returns(uint8 _destinationChainID, address _depositer, bytes32 _resourceID, bytes _metaData)
func (_GenericHandler *GenericHandlerCallerSession) DepositRecords(arg0 uint8, arg1 uint64) (struct {
	DestinationChainID uint8
	Depositer          common.Address
	ResourceID         [32]byte
	MetaData           []byte
}, error) {
	return _GenericHandler.Contract.DepositRecords(&_GenericHandler.CallOpts, arg0, arg1)
}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_GenericHandler *GenericHandlerCaller) ResourceIDToContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "_resourceIDToContractAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_GenericHandler *GenericHandlerSession) ResourceIDToContractAddress(arg0 [32]byte) (common.Address, error) {
	return _GenericHandler.Contract.ResourceIDToContractAddress(&_GenericHandler.CallOpts, arg0)
}

// ResourceIDToContractAddress is a free data retrieval call binding the contract method 0xc54c2a11.
//
// Solidity: function _resourceIDToContractAddress(bytes32 ) view returns(address)
func (_GenericHandler *GenericHandlerCallerSession) ResourceIDToContractAddress(arg0 [32]byte) (common.Address, error) {
	return _GenericHandler.Contract.ResourceIDToContractAddress(&_GenericHandler.CallOpts, arg0)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xba484c09.
//
// Solidity: function getDepositRecord(uint64 depositNonce, uint8 destId) view returns((uint8,address,bytes32,bytes))
func (_GenericHandler *GenericHandlerCaller) GetDepositRecord(opts *bind.CallOpts, depositNonce uint64, destId uint8) (GenericHandlerDepositRecord, error) {
	var out []interface{}
	err := _GenericHandler.contract.Call(opts, &out, "getDepositRecord", depositNonce, destId)

	if err != nil {
		return *new(GenericHandlerDepositRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(GenericHandlerDepositRecord)).(*GenericHandlerDepositRecord)

	return out0, err

}

// GetDepositRecord is a free data retrieval call binding the contract method 0xba484c09.
//
// Solidity: function getDepositRecord(uint64 depositNonce, uint8 destId) view returns((uint8,address,bytes32,bytes))
func (_GenericHandler *GenericHandlerSession) GetDepositRecord(depositNonce uint64, destId uint8) (GenericHandlerDepositRecord, error) {
	return _GenericHandler.Contract.GetDepositRecord(&_GenericHandler.CallOpts, depositNonce, destId)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xba484c09.
//
// Solidity: function getDepositRecord(uint64 depositNonce, uint8 destId) view returns((uint8,address,bytes32,bytes))
func (_GenericHandler *GenericHandlerCallerSession) GetDepositRecord(depositNonce uint64, destId uint8) (GenericHandlerDepositRecord, error) {
	return _GenericHandler.Contract.GetDepositRecord(&_GenericHandler.CallOpts, depositNonce, destId)
}

// Deposit is a paid mutator transaction binding the contract method 0x38995da9.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes data) returns()
func (_GenericHandler *GenericHandlerTransactor) Deposit(opts *bind.TransactOpts, resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _GenericHandler.contract.Transact(opts, "deposit", resourceID, destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x38995da9.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes data) returns()
func (_GenericHandler *GenericHandlerSession) Deposit(resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.Deposit(&_GenericHandler.TransactOpts, resourceID, destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x38995da9.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes data) returns()
func (_GenericHandler *GenericHandlerTransactorSession) Deposit(resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.Deposit(&_GenericHandler.TransactOpts, resourceID, destinationChainID, depositNonce, depositer, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x138f0b8b.
//
// Solidity: function executeProposal(bytes32 resourceID, address[36] signers, bytes data) returns(bool, bool, address, uint256)
func (_GenericHandler *GenericHandlerTransactor) ExecuteProposal(opts *bind.TransactOpts, resourceID [32]byte, signers [36]common.Address, data []byte) (*types.Transaction, error) {
	return _GenericHandler.contract.Transact(opts, "executeProposal", resourceID, signers, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x138f0b8b.
//
// Solidity: function executeProposal(bytes32 resourceID, address[36] signers, bytes data) returns(bool, bool, address, uint256)
func (_GenericHandler *GenericHandlerSession) ExecuteProposal(resourceID [32]byte, signers [36]common.Address, data []byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.ExecuteProposal(&_GenericHandler.TransactOpts, resourceID, signers, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x138f0b8b.
//
// Solidity: function executeProposal(bytes32 resourceID, address[36] signers, bytes data) returns(bool, bool, address, uint256)
func (_GenericHandler *GenericHandlerTransactorSession) ExecuteProposal(resourceID [32]byte, signers [36]common.Address, data []byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.ExecuteProposal(&_GenericHandler.TransactOpts, resourceID, signers, data)
}

// SetResource is a paid mutator transaction binding the contract method 0xbba8185a.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress, bytes4 depositFunctionSig, bytes4 executeFunctionSig) returns()
func (_GenericHandler *GenericHandlerTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address, depositFunctionSig [4]byte, executeFunctionSig [4]byte) (*types.Transaction, error) {
	return _GenericHandler.contract.Transact(opts, "setResource", resourceID, contractAddress, depositFunctionSig, executeFunctionSig)
}

// SetResource is a paid mutator transaction binding the contract method 0xbba8185a.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress, bytes4 depositFunctionSig, bytes4 executeFunctionSig) returns()
func (_GenericHandler *GenericHandlerSession) SetResource(resourceID [32]byte, contractAddress common.Address, depositFunctionSig [4]byte, executeFunctionSig [4]byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.SetResource(&_GenericHandler.TransactOpts, resourceID, contractAddress, depositFunctionSig, executeFunctionSig)
}

// SetResource is a paid mutator transaction binding the contract method 0xbba8185a.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress, bytes4 depositFunctionSig, bytes4 executeFunctionSig) returns()
func (_GenericHandler *GenericHandlerTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address, depositFunctionSig [4]byte, executeFunctionSig [4]byte) (*types.Transaction, error) {
	return _GenericHandler.Contract.SetResource(&_GenericHandler.TransactOpts, resourceID, contractAddress, depositFunctionSig, executeFunctionSig)
}
