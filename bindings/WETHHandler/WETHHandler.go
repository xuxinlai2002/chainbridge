// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package WETHHandler

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

// WETHHandlerDepositRecord is an auto generated low-level Go binding around an user-defined struct.
type WETHHandlerDepositRecord struct {
	TokenAddress                   common.Address
	LenDestinationRecipientAddress uint8
	DestinationChainID             uint8
	ResourceID                     [32]byte
	DestinationRecipientAddress    []byte
	Depositer                      common.Address
	Amount                         *big.Int
}

// WETHHandlerABI is the input ABI used to generate the binding from.
const WETHHandlerABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"initialResourceIDs\",\"type\":\"bytes32[]\"},{\"internalType\":\"address[]\",\"name\":\"initialContractAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"burnableContractAddresses\",\"type\":\"address[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_burnList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_lenDestinationRecipientAddress\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_destinationRecipientAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"address\",\"name\":\"depositer\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"destId\",\"type\":\"uint8\"}],\"name\":\"getDepositRecord\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_lenDestinationRecipientAddress\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"_destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"_destinationRecipientAddress\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"_depositer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"internalType\":\"structWETHHandler.DepositRecord\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOrTokenID\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// WETHHandlerBin is the compiled bytecode used for deploying new contracts.
var WETHHandlerBin = "0x60806040523480156200001157600080fd5b5060405162001c2338038062001c238339818101604052810190620000379190620004c2565b81518351146200007e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200007590620005e1565b60405180910390fd5b836000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060005b8351811015620001135762000105848281518110620000dc57fe5b6020026020010151848381518110620000f157fe5b60200260200101516200015f60201b60201c565b8080600101915050620000c1565b5060005b81518110156200015457620001468282815181106200013257fe5b60200260200101516200025160201b60201c565b808060010191505062000117565b50505050506200071a565b806001600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16620002f5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602481526020018062001bff6024913960400191505060405180910390fd5b6001600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b600062000367620003618462000637565b62000603565b905080838252602082019050828560208602820111156200038757600080fd5b60005b85811015620003bb5781620003a088826200043a565b8452602084019350602083019250506001810190506200038a565b5050509392505050565b6000620003dc620003d68462000666565b62000603565b90508083825260208201905082856020860282011115620003fc57600080fd5b60005b85811015620004305781620004158882620004ab565b845260208401935060208301925050600181019050620003ff565b5050509392505050565b6000815190506200044b81620006e6565b92915050565b600082601f8301126200046357600080fd5b81516200047584826020860162000350565b91505092915050565b600082601f8301126200049057600080fd5b8151620004a2848260208601620003c5565b91505092915050565b600081519050620004bc8162000700565b92915050565b60008060008060808587031215620004d957600080fd5b6000620004e9878288016200043a565b945050602085015167ffffffffffffffff8111156200050757600080fd5b62000515878288016200047e565b935050604085015167ffffffffffffffff8111156200053357600080fd5b620005418782880162000451565b925050606085015167ffffffffffffffff8111156200055f57600080fd5b6200056d8782880162000451565b91505092959194509250565b600062000588603c8362000695565b91507f696e697469616c5265736f7572636549447320616e6420696e697469616c436f60008301527f6e7472616374416464726573736573206c656e206d69736d61746368000000006020830152604082019050919050565b60006020820190508181036000830152620005fc8162000579565b9050919050565b6000604051905081810181811067ffffffffffffffff821117156200062d576200062c620006e4565b5b8060405250919050565b600067ffffffffffffffff821115620006555762000654620006e4565b5b602082029050602081019050919050565b600067ffffffffffffffff821115620006845762000683620006e4565b5b602082029050602081019050919050565b600082825260208201905092915050565b6000620006b382620006c4565b9050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565bfe5b620006f181620006a6565b8114620006fd57600080fd5b50565b6200070b81620006ba565b81146200071757600080fd5b50565b6114d5806200072a6000396000f3fe608060405234801561001057600080fd5b50600436106100b45760003560e01c80637f79bea8116100715780637f79bea8146101a5578063b8fa3736146101d5578063ba484c09146101f1578063c8ba6c8714610221578063d9caed1214610251578063e248cff21461026d576100b4565b806307b7ed99146100b95780630a6d55d8146100d5578063318c136e1461010557806338995da9146101235780634402027f1461013f5780636a70d08114610175575b600080fd5b6100d360048036038101906100ce9190610dd0565b610289565b005b6100ef60048036038101906100ea9190610e48565b61029d565b6040516100fc919061120a565b60405180910390f35b61010d6102d0565b60405161011a919061120a565b60405180910390f35b61013d60048036038101906101389190610f05565b6102f4565b005b61015960048036038101906101549190610fd3565b610598565b60405161016c9796959493929190611225565b60405180910390f35b61018f600480360381019061018a9190610dd0565b6106d9565b60405161019c919061129b565b60405180910390f35b6101bf60048036038101906101ba9190610dd0565b6106f9565b6040516101cc919061129b565b60405180910390f35b6101ef60048036038101906101ea9190610e71565b610719565b005b61020b60048036038101906102069190610f97565b61072f565b60405161021891906112f1565b60405180910390f35b61023b60048036038101906102369190610dd0565b610924565b60405161024891906112b6565b60405180910390f35b61026b60048036038101906102669190610df9565b61093c565b005b61028760048036038101906102829190610ead565b610941565b005b61029161094e565b61029a81610a11565b50565b60016020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6102fc61094e565b606060008060c4359150604051925060e4359050808301602001604052600e360360e484376000600160008b815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff166103e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103dc906112d1565b60405180910390fd5b6040518060e001604052808273ffffffffffffffffffffffffffffffffffffffff1681526020018360ff1681526020018a60ff1681526020018b81526020018581526020018873ffffffffffffffffffffffffffffffffffffffff16815260200184815250600560008b60ff1660ff16815260200190815260200160002060008a67ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548160ff021916908360ff16021790555060408201518160000160156101000a81548160ff021916908360ff160217905550606082015181600101556080820151816002019080519060200190610537929190610c00565b5060a08201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c0820151816004015590505050505050505050505050565b6005602052816000526040600020602052806000526040600020600091509150508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060000160149054906101000a900460ff16908060000160159054906101000a900460ff1690806001015490806002018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106a35780601f10610678576101008083540402835291602001916106a3565b820191906000526020600020905b81548152906001019060200180831161068657829003601f168201915b5050505050908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040154905087565b60046020528060005260406000206000915054906101000a900460ff1681565b60036020528060005260406000206000915054906101000a900460ff1681565b61072161094e565b61072b8282610b0e565b5050565b610737610c8e565b600560008360ff1660ff16815260200190815260200160002060008467ffffffffffffffff1667ffffffffffffffff1681526020019081526020016000206040518060e00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460ff1660ff1660ff1681526020016000820160159054906101000a900460ff1660ff1660ff16815260200160018201548152602001600282018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156108b35780601f10610888576101008083540402835291602001916108b3565b820191906000526020600020905b81548152906001019060200180831161089657829003601f168201915b505050505081526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001600482015481525050905092915050565b60026020528060005260406000206000915090505481565b505050565b61094961094e565b505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610a0f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601e8152602001807f73656e646572206d7573742062652062726964676520636f6e7472616374000081525060200191505060405180910390fd5b565b600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff16610ab3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602481526020018061147c6024913960400191505060405180910390fd5b6001600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b806001600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282610c365760008555610c7d565b82601f10610c4f57805160ff1916838001178555610c7d565b82800160010185558215610c7d579182015b82811115610c7c578251825591602001919060010190610c61565b5b509050610c8a9190610d00565b5090565b6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600060ff168152602001600060ff1681526020016000801916815260200160608152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600081525090565b5b80821115610d19576000816000905550600101610d01565b5090565b600081359050610d2c81611408565b92915050565b600081359050610d418161141f565b92915050565b60008083601f840112610d5957600080fd5b8235905067ffffffffffffffff811115610d7257600080fd5b602083019150836001820283011115610d8a57600080fd5b9250929050565b600081359050610da081611436565b92915050565b600081359050610db58161144d565b92915050565b600081359050610dca81611464565b92915050565b600060208284031215610de257600080fd5b6000610df084828501610d1d565b91505092915050565b600080600060608486031215610e0e57600080fd5b6000610e1c86828701610d1d565b9350506020610e2d86828701610d1d565b9250506040610e3e86828701610d91565b9150509250925092565b600060208284031215610e5a57600080fd5b6000610e6884828501610d32565b91505092915050565b60008060408385031215610e8457600080fd5b6000610e9285828601610d32565b9250506020610ea385828601610d1d565b9150509250929050565b600080600060408486031215610ec257600080fd5b6000610ed086828701610d32565b935050602084013567ffffffffffffffff811115610eed57600080fd5b610ef986828701610d47565b92509250509250925092565b60008060008060008060a08789031215610f1e57600080fd5b6000610f2c89828a01610d32565b9650506020610f3d89828a01610dbb565b9550506040610f4e89828a01610da6565b9450506060610f5f89828a01610d1d565b935050608087013567ffffffffffffffff811115610f7c57600080fd5b610f8889828a01610d47565b92509250509295509295509295565b60008060408385031215610faa57600080fd5b6000610fb885828601610da6565b9250506020610fc985828601610dbb565b9150509250929050565b60008060408385031215610fe657600080fd5b6000610ff485828601610dbb565b925050602061100585828601610da6565b9150509250929050565b61101881611351565b82525050565b61102781611351565b82525050565b61103681611363565b82525050565b6110458161136f565b82525050565b6110548161136f565b82525050565b600061106582611313565b61106f818561131e565b935061107f8185602086016113c4565b611088816113f7565b840191505092915050565b600061109e82611313565b6110a8818561132f565b93506110b88185602086016113c4565b6110c1816113f7565b840191505092915050565b60006110d9602883611340565b91507f70726f766964656420746f6b656e41646472657373206973206e6f742077686960008301527f74656c69737465640000000000000000000000000000000000000000000000006020830152604082019050919050565b600060e08301600083015161114a600086018261100f565b50602083015161115d60208601826111ec565b50604083015161117060408601826111ec565b506060830151611183606086018261103c565b506080830151848203608086015261119b828261105a565b91505060a08301516111b060a086018261100f565b5060c08301516111c360c08601826111ce565b508091505092915050565b6111d781611399565b82525050565b6111e681611399565b82525050565b6111f5816113b7565b82525050565b611204816113b7565b82525050565b600060208201905061121f600083018461101e565b92915050565b600060e08201905061123a600083018a61101e565b61124760208301896111fb565b61125460408301886111fb565b611261606083018761104b565b81810360808301526112738186611093565b905061128260a083018561101e565b61128f60c08301846111dd565b98975050505050505050565b60006020820190506112b0600083018461102d565b92915050565b60006020820190506112cb600083018461104b565b92915050565b600060208201905081810360008301526112ea816110cc565b9050919050565b6000602082019050818103600083015261130b8184611132565b905092915050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600061135c82611379565b9050919050565b60008115159050919050565b6000819050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b600060ff82169050919050565b60005b838110156113e25780820151818401526020810190506113c7565b838111156113f1576000848401525b50505050565b6000601f19601f8301169050919050565b61141181611351565b811461141c57600080fd5b50565b6114288161136f565b811461143357600080fd5b50565b61143f81611399565b811461144a57600080fd5b50565b611456816113a3565b811461146157600080fd5b50565b61146d816113b7565b811461147857600080fd5b5056fe70726f766964656420636f6e7472616374206973206e6f742077686974656c6973746564a2646970667358221220d0d9216dd7b08b7ba42ddc5fef1d4e8582f1a0eaef383a471d773795cc60567b64736f6c6343000706003370726f766964656420636f6e7472616374206973206e6f742077686974656c6973746564"

// DeployWETHHandler deploys a new Ethereum contract, binding an instance of WETHHandler to it.
func DeployWETHHandler(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address, initialResourceIDs [][32]byte, initialContractAddresses []common.Address, burnableContractAddresses []common.Address) (common.Address, *types.Transaction, *WETHHandler, error) {
	parsed, err := abi.JSON(strings.NewReader(WETHHandlerABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(WETHHandlerBin), backend, bridgeAddress, initialResourceIDs, initialContractAddresses, burnableContractAddresses)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &WETHHandler{WETHHandlerCaller: WETHHandlerCaller{contract: contract}, WETHHandlerTransactor: WETHHandlerTransactor{contract: contract}, WETHHandlerFilterer: WETHHandlerFilterer{contract: contract}}, nil
}

// WETHHandler is an auto generated Go binding around an Ethereum contract.
type WETHHandler struct {
	WETHHandlerCaller     // Read-only binding to the contract
	WETHHandlerTransactor // Write-only binding to the contract
	WETHHandlerFilterer   // Log filterer for contract events
}

// WETHHandlerCaller is an auto generated read-only Go binding around an Ethereum contract.
type WETHHandlerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WETHHandlerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WETHHandlerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WETHHandlerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WETHHandlerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WETHHandlerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WETHHandlerSession struct {
	Contract     *WETHHandler      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WETHHandlerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WETHHandlerCallerSession struct {
	Contract *WETHHandlerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// WETHHandlerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WETHHandlerTransactorSession struct {
	Contract     *WETHHandlerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// WETHHandlerRaw is an auto generated low-level Go binding around an Ethereum contract.
type WETHHandlerRaw struct {
	Contract *WETHHandler // Generic contract binding to access the raw methods on
}

// WETHHandlerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WETHHandlerCallerRaw struct {
	Contract *WETHHandlerCaller // Generic read-only contract binding to access the raw methods on
}

// WETHHandlerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WETHHandlerTransactorRaw struct {
	Contract *WETHHandlerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWETHHandler creates a new instance of WETHHandler, bound to a specific deployed contract.
func NewWETHHandler(address common.Address, backend bind.ContractBackend) (*WETHHandler, error) {
	contract, err := bindWETHHandler(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WETHHandler{WETHHandlerCaller: WETHHandlerCaller{contract: contract}, WETHHandlerTransactor: WETHHandlerTransactor{contract: contract}, WETHHandlerFilterer: WETHHandlerFilterer{contract: contract}}, nil
}

// NewWETHHandlerCaller creates a new read-only instance of WETHHandler, bound to a specific deployed contract.
func NewWETHHandlerCaller(address common.Address, caller bind.ContractCaller) (*WETHHandlerCaller, error) {
	contract, err := bindWETHHandler(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WETHHandlerCaller{contract: contract}, nil
}

// NewWETHHandlerTransactor creates a new write-only instance of WETHHandler, bound to a specific deployed contract.
func NewWETHHandlerTransactor(address common.Address, transactor bind.ContractTransactor) (*WETHHandlerTransactor, error) {
	contract, err := bindWETHHandler(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WETHHandlerTransactor{contract: contract}, nil
}

// NewWETHHandlerFilterer creates a new log filterer instance of WETHHandler, bound to a specific deployed contract.
func NewWETHHandlerFilterer(address common.Address, filterer bind.ContractFilterer) (*WETHHandlerFilterer, error) {
	contract, err := bindWETHHandler(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WETHHandlerFilterer{contract: contract}, nil
}

// bindWETHHandler binds a generic wrapper to an already deployed contract.
func bindWETHHandler(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WETHHandlerABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WETHHandler *WETHHandlerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WETHHandler.Contract.WETHHandlerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WETHHandler *WETHHandlerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETHHandler.Contract.WETHHandlerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WETHHandler *WETHHandlerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WETHHandler.Contract.WETHHandlerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WETHHandler *WETHHandlerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WETHHandler.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WETHHandler *WETHHandlerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETHHandler.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WETHHandler *WETHHandlerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WETHHandler.Contract.contract.Transact(opts, method, params...)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_WETHHandler *WETHHandlerCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WETHHandler.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_WETHHandler *WETHHandlerSession) BridgeAddress() (common.Address, error) {
	return _WETHHandler.Contract.BridgeAddress(&_WETHHandler.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_WETHHandler *WETHHandlerCallerSession) BridgeAddress() (common.Address, error) {
	return _WETHHandler.Contract.BridgeAddress(&_WETHHandler.CallOpts)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_WETHHandler *WETHHandlerCaller) BurnList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _WETHHandler.contract.Call(opts, &out, "_burnList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_WETHHandler *WETHHandlerSession) BurnList(arg0 common.Address) (bool, error) {
	return _WETHHandler.Contract.BurnList(&_WETHHandler.CallOpts, arg0)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_WETHHandler *WETHHandlerCallerSession) BurnList(arg0 common.Address) (bool, error) {
	return _WETHHandler.Contract.BurnList(&_WETHHandler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_WETHHandler *WETHHandlerCaller) ContractWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _WETHHandler.contract.Call(opts, &out, "_contractWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_WETHHandler *WETHHandlerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _WETHHandler.Contract.ContractWhitelist(&_WETHHandler.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_WETHHandler *WETHHandlerCallerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _WETHHandler.Contract.ContractWhitelist(&_WETHHandler.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0x4402027f.
//
// Solidity: function _depositRecords(uint8 , uint64 ) view returns(address _tokenAddress, uint8 _lenDestinationRecipientAddress, uint8 _destinationChainID, bytes32 _resourceID, bytes _destinationRecipientAddress, address _depositer, uint256 _amount)
func (_WETHHandler *WETHHandlerCaller) DepositRecords(opts *bind.CallOpts, arg0 uint8, arg1 uint64) (struct {
	TokenAddress                   common.Address
	LenDestinationRecipientAddress uint8
	DestinationChainID             uint8
	ResourceID                     [32]byte
	DestinationRecipientAddress    []byte
	Depositer                      common.Address
	Amount                         *big.Int
}, error) {
	var out []interface{}
	err := _WETHHandler.contract.Call(opts, &out, "_depositRecords", arg0, arg1)

	outstruct := new(struct {
		TokenAddress                   common.Address
		LenDestinationRecipientAddress uint8
		DestinationChainID             uint8
		ResourceID                     [32]byte
		DestinationRecipientAddress    []byte
		Depositer                      common.Address
		Amount                         *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.LenDestinationRecipientAddress = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.DestinationChainID = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.ResourceID = *abi.ConvertType(out[3], new([32]byte)).(*[32]byte)
	outstruct.DestinationRecipientAddress = *abi.ConvertType(out[4], new([]byte)).(*[]byte)
	outstruct.Depositer = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DepositRecords is a free data retrieval call binding the contract method 0x4402027f.
//
// Solidity: function _depositRecords(uint8 , uint64 ) view returns(address _tokenAddress, uint8 _lenDestinationRecipientAddress, uint8 _destinationChainID, bytes32 _resourceID, bytes _destinationRecipientAddress, address _depositer, uint256 _amount)
func (_WETHHandler *WETHHandlerSession) DepositRecords(arg0 uint8, arg1 uint64) (struct {
	TokenAddress                   common.Address
	LenDestinationRecipientAddress uint8
	DestinationChainID             uint8
	ResourceID                     [32]byte
	DestinationRecipientAddress    []byte
	Depositer                      common.Address
	Amount                         *big.Int
}, error) {
	return _WETHHandler.Contract.DepositRecords(&_WETHHandler.CallOpts, arg0, arg1)
}

// DepositRecords is a free data retrieval call binding the contract method 0x4402027f.
//
// Solidity: function _depositRecords(uint8 , uint64 ) view returns(address _tokenAddress, uint8 _lenDestinationRecipientAddress, uint8 _destinationChainID, bytes32 _resourceID, bytes _destinationRecipientAddress, address _depositer, uint256 _amount)
func (_WETHHandler *WETHHandlerCallerSession) DepositRecords(arg0 uint8, arg1 uint64) (struct {
	TokenAddress                   common.Address
	LenDestinationRecipientAddress uint8
	DestinationChainID             uint8
	ResourceID                     [32]byte
	DestinationRecipientAddress    []byte
	Depositer                      common.Address
	Amount                         *big.Int
}, error) {
	return _WETHHandler.Contract.DepositRecords(&_WETHHandler.CallOpts, arg0, arg1)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_WETHHandler *WETHHandlerCaller) ResourceIDToTokenContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _WETHHandler.contract.Call(opts, &out, "_resourceIDToTokenContractAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_WETHHandler *WETHHandlerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _WETHHandler.Contract.ResourceIDToTokenContractAddress(&_WETHHandler.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_WETHHandler *WETHHandlerCallerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _WETHHandler.Contract.ResourceIDToTokenContractAddress(&_WETHHandler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_WETHHandler *WETHHandlerCaller) TokenContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _WETHHandler.contract.Call(opts, &out, "_tokenContractAddressToResourceID", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_WETHHandler *WETHHandlerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _WETHHandler.Contract.TokenContractAddressToResourceID(&_WETHHandler.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_WETHHandler *WETHHandlerCallerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _WETHHandler.Contract.TokenContractAddressToResourceID(&_WETHHandler.CallOpts, arg0)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xba484c09.
//
// Solidity: function getDepositRecord(uint64 depositNonce, uint8 destId) view returns((address,uint8,uint8,bytes32,bytes,address,uint256))
func (_WETHHandler *WETHHandlerCaller) GetDepositRecord(opts *bind.CallOpts, depositNonce uint64, destId uint8) (WETHHandlerDepositRecord, error) {
	var out []interface{}
	err := _WETHHandler.contract.Call(opts, &out, "getDepositRecord", depositNonce, destId)

	if err != nil {
		return *new(WETHHandlerDepositRecord), err
	}

	out0 := *abi.ConvertType(out[0], new(WETHHandlerDepositRecord)).(*WETHHandlerDepositRecord)

	return out0, err

}

// GetDepositRecord is a free data retrieval call binding the contract method 0xba484c09.
//
// Solidity: function getDepositRecord(uint64 depositNonce, uint8 destId) view returns((address,uint8,uint8,bytes32,bytes,address,uint256))
func (_WETHHandler *WETHHandlerSession) GetDepositRecord(depositNonce uint64, destId uint8) (WETHHandlerDepositRecord, error) {
	return _WETHHandler.Contract.GetDepositRecord(&_WETHHandler.CallOpts, depositNonce, destId)
}

// GetDepositRecord is a free data retrieval call binding the contract method 0xba484c09.
//
// Solidity: function getDepositRecord(uint64 depositNonce, uint8 destId) view returns((address,uint8,uint8,bytes32,bytes,address,uint256))
func (_WETHHandler *WETHHandlerCallerSession) GetDepositRecord(depositNonce uint64, destId uint8) (WETHHandlerDepositRecord, error) {
	return _WETHHandler.Contract.GetDepositRecord(&_WETHHandler.CallOpts, depositNonce, destId)
}

// Deposit is a paid mutator transaction binding the contract method 0x38995da9.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes data) returns()
func (_WETHHandler *WETHHandlerTransactor) Deposit(opts *bind.TransactOpts, resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _WETHHandler.contract.Transact(opts, "deposit", resourceID, destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x38995da9.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes data) returns()
func (_WETHHandler *WETHHandlerSession) Deposit(resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _WETHHandler.Contract.Deposit(&_WETHHandler.TransactOpts, resourceID, destinationChainID, depositNonce, depositer, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x38995da9.
//
// Solidity: function deposit(bytes32 resourceID, uint8 destinationChainID, uint64 depositNonce, address depositer, bytes data) returns()
func (_WETHHandler *WETHHandlerTransactorSession) Deposit(resourceID [32]byte, destinationChainID uint8, depositNonce uint64, depositer common.Address, data []byte) (*types.Transaction, error) {
	return _WETHHandler.Contract.Deposit(&_WETHHandler.TransactOpts, resourceID, destinationChainID, depositNonce, depositer, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_WETHHandler *WETHHandlerTransactor) ExecuteProposal(opts *bind.TransactOpts, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _WETHHandler.contract.Transact(opts, "executeProposal", resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_WETHHandler *WETHHandlerSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _WETHHandler.Contract.ExecuteProposal(&_WETHHandler.TransactOpts, resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0xe248cff2.
//
// Solidity: function executeProposal(bytes32 resourceID, bytes data) returns()
func (_WETHHandler *WETHHandlerTransactorSession) ExecuteProposal(resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _WETHHandler.Contract.ExecuteProposal(&_WETHHandler.TransactOpts, resourceID, data)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_WETHHandler *WETHHandlerTransactor) SetBurnable(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _WETHHandler.contract.Transact(opts, "setBurnable", contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_WETHHandler *WETHHandlerSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _WETHHandler.Contract.SetBurnable(&_WETHHandler.TransactOpts, contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_WETHHandler *WETHHandlerTransactorSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _WETHHandler.Contract.SetBurnable(&_WETHHandler.TransactOpts, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_WETHHandler *WETHHandlerTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _WETHHandler.contract.Transact(opts, "setResource", resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_WETHHandler *WETHHandlerSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _WETHHandler.Contract.SetResource(&_WETHHandler.TransactOpts, resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_WETHHandler *WETHHandlerTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _WETHHandler.Contract.SetResource(&_WETHHandler.TransactOpts, resourceID, contractAddress)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_WETHHandler *WETHHandlerTransactor) Withdraw(opts *bind.TransactOpts, tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _WETHHandler.contract.Transact(opts, "withdraw", tokenAddress, recipient, amountOrTokenID)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_WETHHandler *WETHHandlerSession) Withdraw(tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _WETHHandler.Contract.Withdraw(&_WETHHandler.TransactOpts, tokenAddress, recipient, amountOrTokenID)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_WETHHandler *WETHHandlerTransactorSession) Withdraw(tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _WETHHandler.Contract.Withdraw(&_WETHHandler.TransactOpts, tokenAddress, recipient, amountOrTokenID)
}
