// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Bridge

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

// BridgeProposal is an auto generated low-level Go binding around an user-defined struct.
type BridgeProposal struct {
	ResourceID    [32]byte
	DataHash      [32]byte
	YesVotes      []common.Address
	NoVotes       []common.Address
	Status        uint8
	ProposedBlock *big.Int
}

// BridgeABI is the input ABI used to generate the binding from.
const BridgeABI = "[{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"chainID\",\"type\":\"uint8\"},{\"internalType\":\"address[]\",\"name\":\"initialRelayers\",\"type\":\"address[]\"},{\"internalType\":\"uint256\",\"name\":\"initialRelayerThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiry\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"data\",\"type\":\"string\"}],\"name\":\"LogString\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"originChainID\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"enumBridge.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"ProposalEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"originChainID\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"indexed\":true,\"internalType\":\"enumBridge.ProposalStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"}],\"name\":\"ProposalVote\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"RelayerAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"RelayerRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"RelayerThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RELAYER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_bridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_burnList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_chainID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_contractWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"_depositCounts\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"name\":\"_depositRecords\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_expiry\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint72\",\"name\":\"\",\"type\":\"uint72\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_hasVotedOnProposal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint72\",\"name\":\"\",\"type\":\"uint72\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_proposals\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumBridge.ProposalStatus\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_proposedBlock\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"h\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"_recoverSigner2\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_relayerThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToHandlerAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"_resourceIDToTokenContractAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"_tokenContractAddressToResourceID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_totalProposals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"_totalRelayers\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"addressToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayerAddress\",\"type\":\"address\"}],\"name\":\"adminAddRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newThreshold\",\"type\":\"uint256\"}],\"name\":\"adminChangeRelayerThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminPauseTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayerAddress\",\"type\":\"address\"}],\"name\":\"adminRemoveRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"adminSetBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"depositFunctionSig\",\"type\":\"bytes4\"},{\"internalType\":\"bytes4\",\"name\":\"executeFunctionSig\",\"type\":\"bytes4\"}],\"name\":\"adminSetGenericResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"handlerAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"adminSetResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"adminUnpauseTransfers\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"bytes32ToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"bytesToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"destinationChainID\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"chainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"}],\"name\":\"executeProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"originChainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"getProposal\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"_resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"_dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[]\",\"name\":\"_yesVotes\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"_noVotes\",\"type\":\"address[]\"},{\"internalType\":\"enumBridge.ProposalStatus\",\"name\":\"_status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"_proposedBlock\",\"type\":\"uint256\"}],\"internalType\":\"structBridge.Proposal\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"relayer\",\"type\":\"address\"}],\"name\":\"isRelayer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"renounceAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"_pubKeyList\",\"type\":\"bytes[]\"}],\"name\":\"setAbiterList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setBurnable\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"contractAddress\",\"type\":\"address\"}],\"name\":\"setResource\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"sub\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"toStringBase\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable[]\",\"name\":\"addrs\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"transferFunds\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"uintToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"chainID\",\"type\":\"uint8\"},{\"internalType\":\"uint64\",\"name\":\"depositNonce\",\"type\":\"uint64\"},{\"internalType\":\"bytes32\",\"name\":\"resourceID\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"}],\"name\":\"voteProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOrTokenID\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// BridgeBin is the compiled bytecode used for deploying new contracts.
var BridgeBin = "0x6080604052604051620065053803806200650583398181016040528101906200002991906200050a565b6000600560006101000a81548160ff02191690831515021790555084600760006101000a81548160ff021916908360ff1602179055508260088190555081600b8190555080600c81905550620000896000801b336200013d60201b60201c565b620000be7fe2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc46000801b6200015360201b60201c565b60005b84518110156200013157620001117fe2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc4868381518110620000fd57fe5b6020026020010151620001b760201b60201c565b6009600081548092919060010191905055508080600101915050620000c1565b50505050505062000730565b6200014f82826200024660201b60201c565b5050565b806006600084815260200190815260200160002060020154837fbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff60405160405180910390a48060066000848152602001908152602001600020600201819055505050565b620001ee6006600084815260200190815260200160002060020154620001e2620002ea60201b60201c565b620002f260201b60201c565b62000230576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000227906200060d565b60405180910390fd5b6200024282826200024660201b60201c565b5050565b6200027581600660008581526020019081526020016000206000016200032b60201b62002f3c1790919060201c565b15620002e6576200028b620002ea60201b60201c565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b600033905090565b60006200032382600660008681526020019081526020016000206000016200036360201b62002f6c1790919060201c565b905092915050565b60006200035b836000018373ffffffffffffffffffffffffffffffffffffffff1660001b6200039b60201b60201c565b905092915050565b600062000393836000018373ffffffffffffffffffffffffffffffffffffffff1660001b6200041560201b60201c565b905092915050565b6000620003af83836200041560201b60201c565b6200040a5782600001829080600181540180825580915050600190039060005260206000200160009091909190915055826000018054905083600101600084815260200190815260200160002081905550600190506200040f565b600090505b92915050565b600080836001016000848152602001908152602001600020541415905092915050565b6000815190506200044981620006e2565b92915050565b600082601f8301126200046157600080fd5b81516200047862000472826200065d565b6200062f565b915081818352602084019350602081019050838560208402820111156200049e57600080fd5b60005b83811015620004d25781620004b7888262000438565b845260208401935060208301925050600181019050620004a1565b5050505092915050565b600081519050620004ed81620006fc565b92915050565b600081519050620005048162000716565b92915050565b600080600080600060a086880312156200052357600080fd5b60006200053388828901620004f3565b955050602086015167ffffffffffffffff8111156200055157600080fd5b6200055f888289016200044f565b94505060406200057288828901620004dc565b93505060606200058588828901620004dc565b92505060806200059888828901620004dc565b9150509295509295909350565b6000620005b4602f8362000686565b91507f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60008301527f2061646d696e20746f206772616e7400000000000000000000000000000000006020830152604082019050919050565b600060208201905081810360008301526200062881620005a5565b9050919050565b6000604051905081810181811067ffffffffffffffff821117156200065357600080fd5b8060405250919050565b600067ffffffffffffffff8211156200067557600080fd5b602082029050602081019050919050565b600082825260208201905092915050565b6000620006a482620006ab565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600060ff82169050919050565b620006ed8162000697565b8114620006f957600080fd5b50565b6200070781620006cb565b81146200071357600080fd5b50565b6200072181620006d5565b81146200072d57600080fd5b50565b615dc580620007406000396000f3fe6080604052600436106103035760003560e01c80638c0c263111610190578063beab7131116100dc578063cdb0f73a11610095578063d9caed121161006f578063d9caed1214610c5b578063e8437ee714610c84578063e939567914610cad578063ffaac0eb14610cea5761030a565b8063cdb0f73a14610bde578063d547741f14610c07578063d7a9cd7914610c305761030a565b8063beab713114610aba578063c5b37c2214610ae5578063c5ec897014610b10578063c8ba6c8714610b3b578063ca15c87314610b78578063cb10f21514610bb55761030a565b8063926d7d7f11610149578063a217fddf11610123578063a217fddf146109ec578063a9cf69fa14610a17578063b67d77c514610a54578063b8fa373614610a915761030a565b8063926d7d7f1461096d5780639d5773e0146109985780639d82dd63146109c35761030a565b80638c0c2631146108135780638d1b75e41461083c5780638f667892146108795780639010d07c146108b657806391d14854146108f35780639201de55146109305761030a565b80634e0560051161024f5780636a70d081116102085780637febe63f116101e25780637febe63f14610757578063802aabe81461079457806380ae1c28146107bf57806384db809f146107d65761030a565b80636a70d081146106b45780637263cd76146106f15780637f79bea81461071a5761030a565b80634e0560051461057d57806350598719146105a6578063541d5548146105e65780635c975abb146106235780635e1fab0f1461064e5780635e57966d146106775761030a565b8063318c136e116102bc5780633ee7094a116102965780633ee7094a146104b15780634454b20d146104ee5780634603ae38146105175780634b0b919d146105405761030a565b8063318c136e1461042057806336568abe1461044b57806339614e4f146104745761030a565b806305e2ca171461030f57806307b7ed991461032b5780630a6d55d8146103545780631ff013f114610391578063248a9ca3146103ba5780632f2ff15d146103f75761030a565b3661030a57005b600080fd5b610329600480360381019061032491906142d7565b610d01565b005b34801561033757600080fd5b50610352600480360381019061034d9190613d81565b610fc3565b005b34801561036057600080fd5b5061037b6004803603810190610376919061403d565b610fd7565b6040516103889190615234565b60405180910390f35b34801561039d57600080fd5b506103b860048036038101906103b39190614392565b61100a565b005b3480156103c657600080fd5b506103e160048036038101906103dc919061403d565b611a10565b6040516103ee9190615285565b60405180910390f35b34801561040357600080fd5b5061041e60048036038101906104199190614066565b611a30565b005b34801561042c57600080fd5b50610435611aa4565b6040516104429190615234565b60405180910390f35b34801561045757600080fd5b50610472600480360381019061046d9190614066565b611ac8565b005b34801561048057600080fd5b5061049b60048036038101906104969190614141565b611b4b565b6040516104a89190615481565b60405180910390f35b3480156104bd57600080fd5b506104d860048036038101906104d391906141e7565b611b7c565b6040516104e5919061545f565b60405180910390f35b3480156104fa57600080fd5b50610515600480360381019061051091906143f5565b611c39565b005b34801561052357600080fd5b5061053e60048036038101906105399190613f24565b612039565b005b34801561054c57600080fd5b50610567600480360381019061056291906142ae565b612111565b6040516105749190615940565b60405180910390f35b34801561058957600080fd5b506105a4600480360381019061059f9190614182565b612138565b005b3480156105b257600080fd5b506105cd60048036038101906105c89190614223565b612177565b6040516105dd9493929190615379565b60405180910390f35b3480156105f257600080fd5b5061060d60048036038101906106089190613d81565b6121c1565b60405161061a919061526a565b60405180910390f35b34801561062f57600080fd5b506106386121f4565b604051610645919061526a565b60405180910390f35b34801561065a57600080fd5b5061067560048036038101906106709190613d81565b61220b565b005b34801561068357600080fd5b5061069e60048036038101906106999190613d81565b612230565b6040516106ab9190615481565b60405180910390f35b3480156106c057600080fd5b506106db60048036038101906106d69190613d81565b612261565b6040516106e8919061526a565b60405180910390f35b3480156106fd57600080fd5b5061071860048036038101906107139190613f99565b612281565b005b34801561072657600080fd5b50610741600480360381019061073c9190613d81565b612312565b60405161074e919061526a565b60405180910390f35b34801561076357600080fd5b5061077e6004803603810190610779919061425f565b612332565b60405161078b919061526a565b60405180910390f35b3480156107a057600080fd5b506107a961236e565b6040516107b69190615925565b60405180910390f35b3480156107cb57600080fd5b506107d4612374565b005b3480156107e257600080fd5b506107fd60048036038101906107f8919061403d565b612386565b60405161080a9190615234565b60405180910390f35b34801561081f57600080fd5b5061083a60048036038101906108359190613dd3565b6123b9565b005b34801561084857600080fd5b50610863600480360381019061085e9190614141565b612436565b6040516108709190615481565b60405180910390f35b34801561088557600080fd5b506108a0600480360381019061089b91906140de565b6126b9565b6040516108ad9190615234565b60405180910390f35b3480156108c257600080fd5b506108dd60048036038101906108d891906140a2565b612783565b6040516108ea9190615234565b60405180910390f35b3480156108ff57600080fd5b5061091a60048036038101906109159190614066565b6127b5565b604051610927919061526a565b60405180910390f35b34801561093c57600080fd5b506109576004803603810190610952919061403d565b6127e7565b6040516109649190615481565b60405180910390f35b34801561097957600080fd5b50610982612818565b60405161098f9190615285565b60405180910390f35b3480156109a457600080fd5b506109ad61283c565b6040516109ba9190615925565b60405180910390f35b3480156109cf57600080fd5b506109ea60048036038101906109e59190613d81565b612842565b005b3480156109f857600080fd5b50610a01612936565b604051610a0e9190615285565b60405180910390f35b348015610a2357600080fd5b50610a3e6004803603810190610a399190614343565b61293d565b604051610a4b9190615903565b60405180910390f35b348015610a6057600080fd5b50610a7b6004803603810190610a7691906141ab565b612b1e565b604051610a889190615925565b60405180910390f35b348015610a9d57600080fd5b50610ab86004803603810190610ab39190614066565b612b68565b005b348015610ac657600080fd5b50610acf612b7e565b604051610adc919061595b565b60405180910390f35b348015610af157600080fd5b50610afa612b91565b604051610b079190615925565b60405180910390f35b348015610b1c57600080fd5b50610b25612b97565b604051610b329190615925565b60405180910390f35b348015610b4757600080fd5b50610b626004803603810190610b5d9190613d81565b612b9d565b604051610b6f9190615285565b60405180910390f35b348015610b8457600080fd5b50610b9f6004803603810190610b9a919061403d565b612bb5565b604051610bac9190615925565b60405180910390f35b348015610bc157600080fd5b50610bdc6004803603810190610bd79190613e5e565b612bdc565b005b348015610bea57600080fd5b50610c056004803603810190610c009190613d81565b612cae565b005b348015610c1357600080fd5b50610c2e6004803603810190610c299190614066565b612da2565b005b348015610c3c57600080fd5b50610c45612e16565b604051610c529190615925565b60405180910390f35b348015610c6757600080fd5b50610c826004803603810190610c7d9190613e0f565b612e1c565b005b348015610c9057600080fd5b50610cab6004803603810190610ca69190613ead565b612e21565b005b348015610cb957600080fd5b50610cd46004803603810190610ccf9190614182565b612ef9565b604051610ce19190615481565b60405180910390f35b348015610cf657600080fd5b50610cff612f2a565b005b610d09612f9c565b7fa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f604051610d3690615663565b60405180910390a16000600e600085815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1690507fa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f610da082612230565b604051610dad9190615481565b60405180910390a1600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610e25576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e1c906156e3565b60405180910390fd5b6000600d60008760ff1660ff168152602001908152602001600020600081819054906101000a900467ffffffffffffffff1660010191906101000a81548167ffffffffffffffff021916908367ffffffffffffffff160217905590508383600f60008467ffffffffffffffff1667ffffffffffffffff16815260200190815260200160002060008960ff1660ff1681526020019081526020016000209190610ece92919061394b565b5060008290508073ffffffffffffffffffffffffffffffffffffffff166338995da9878985338a8a6040518763ffffffff1660e01b8152600401610f1796959493929190615403565b600060405180830381600087803b158015610f3157600080fd5b505af1158015610f45573d6000803e3d6000fd5b505050507fa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f604051610f76906157e3565b60405180910390a18167ffffffffffffffff16868860ff167fdbb69440df8433824a026ef190652f29929eb64b4d1d5d2a69be8afe3e6eaed860405160405180910390a450505050505050565b610fcb612fee565b610fd48161307e565b50565b60016020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b611012613165565b61101a612f9c565b7fa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f60405161104790615763565b60405180910390a160008460ff1660088567ffffffffffffffff1668ffffffffffffffffff16901b1790506000601060008368ffffffffffffffffff1668ffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000209050600073ffffffffffffffffffffffffffffffffffffffff16600e600086815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415611153576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161114a906158c3565b60405180910390fd5b60018160040160009054906101000a900460ff16600481111561117257fe5b11156111b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111aa90615803565b60405180910390fd5b601160008368ffffffffffffffffff1668ffffffffffffffffff168152602001908152602001600020600084815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1615611278576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161126f90615523565b60405180910390fd5b7fa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f6040516112a590615783565b60405180910390a160008160040160009054906101000a900460ff1660048111156112cc57fe5b141561169d57600a60008154600101919050819055506040518060c00160405280858152602001848152602001600167ffffffffffffffff8111801561131157600080fd5b506040519080825280602002602001820160405280156113405781602001602082028036833780820191505090505b508152602001600067ffffffffffffffff8111801561135e57600080fd5b5060405190808252806020026020018201604052801561138d5781602001602082028036833780820191505090505b508152602001600160048111156113a057fe5b815260200143815250601060008468ffffffffffffffffff1668ffffffffffffffffff1681526020019081526020016000206000858152602001908152602001600020600082015181600001556020820151816001015560408201518160020190805190602001906114139291906139cb565b5060608201518160030190805190602001906114309291906139cb565b5060808201518160040160006101000a81548160ff0219169083600481111561145557fe5b021790555060a08201518160050155905050338160020160008154811061147857fe5b9060005260206000200160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f6040516114ed906155e3565b60405180910390a17fa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f61152a8368ffffffffffffffffff16612ef9565b6040516115379190615481565b60405180910390a17fa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f60405161156c90615683565b60405180910390a17fa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f61159e846127e7565b6040516115ab9190615481565b60405180910390a17fa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f6040516115e090615863565b60405180910390a17fa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f61162e8260040160009054906101000a900460ff16600481111561162957fe5b612ef9565b60405161163b9190615481565b60405180910390a16001600481111561165057fe5b8567ffffffffffffffff168760ff167f803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab6504178787604051611690929190615350565b60405180910390a4611815565b7fa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f6040516116ca90615603565b60405180910390a1600c546116e3438360050154612b1e565b11156117685760048160040160006101000a81548160ff0219169083600481111561170a57fe5b021790555060048081111561171b57fe5b8567ffffffffffffffff168760ff167f803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab650417878760405161175b929190615350565b60405180910390a4611814565b806001015483146117ae576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117a5906158a3565b60405180910390fd5b80600201339080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5b7fa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f604051611842906156c3565b60405180910390a160048081111561185657fe5b8160040160009054906101000a900460ff16600481111561187357fe5b14611a08576001601160008468ffffffffffffffffff1668ffffffffffffffffff168152602001908152602001600020600085815260200190815260200160002060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508060040160009054906101000a900460ff16600481111561192557fe5b8567ffffffffffffffff168760ff167f25f8daaa4635a7729927ba3f5b3d59cc3320aca7c32c9db4e7ca7b9574343640876040516119639190615285565b60405180910390a460016008541115806119865750600854816002018054905010155b15611a075760028160040160006101000a81548160ff021916908360048111156119ac57fe5b0217905550600260048111156119be57fe5b8567ffffffffffffffff168760ff167f803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab65041787876040516119fe929190615350565b60405180910390a45b5b505050505050565b600060066000838152602001908152602001600020600201549050919050565b611a576006600084815260200190815260200160002060020154611a526131d0565b6127b5565b611a96576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a8d90615503565b60405180910390fd5b611aa082826131d8565b5050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b611ad06131d0565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611b3d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b34906158e3565b60405180910390fd5b611b47828261326c565b5050565b6060611b7582604051602001611b6191906151da565b604051602081830303815290604052612436565b9050919050565b600f602052816000526040600020602052806000526040600020600091509150508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611c315780601f10611c0657610100808354040283529160200191611c31565b820191906000526020600020905b815481529060010190602001808311611c1457829003601f168201915b505050505081565b611c41613165565b611c49612f9c565b7fa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f604051611c7690615583565b60405180910390a16000806000806000600e600087815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008a60ff1660088b67ffffffffffffffff1668ffffffffffffffffff16901b1790506000828a8a604051602001611cf693929190615195565b6040516020818303038152906040528051906020012090506000601060008468ffffffffffffffffff1668ffffffffffffffffff1681526020019081526020016000206000838152602001908152602001600020905060006004811115611d5957fe5b8160040160009054906101000a900460ff166004811115611d7657fe5b1415611db7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611dae906154c3565b60405180910390fd5b60026004811115611dc457fe5b8160040160009054906101000a900460ff166004811115611de157fe5b14611e21576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e18906154a3565b60405180910390fd5b80600101548214611e67576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e5e906156a3565b60405180910390fd5b60038160040160006101000a81548160ff02191690836004811115611e8857fe5b021790555060008490508073ffffffffffffffffffffffffffffffffffffffff1663138f0b8b836000015460128f8f6040518563ffffffff1660e01b8152600401611ed6949392919061530e565b608060405180830381600087803b158015611ef057600080fd5b505af1158015611f04573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611f289190613fda565b809950819a50829b50839c50505050507fa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f604051611f6590615623565b60405180910390a1888015611f775750875b15611fbc577fa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f604051611fa9906157a3565b60405180910390a1611fbb8787613300565b5b8160040160009054906101000a900460ff166004811115611fd957fe5b8d67ffffffffffffffff168f60ff167f803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab65041785600001548660010154604051612021929190615350565b60405180910390a45050505050505050505050505050565b612041613433565b7fa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f60405161206e90615883565b60405180910390a160005b8484905081101561210a5784848281811061209057fe5b90506020020160208101906120a59190613daa565b73ffffffffffffffffffffffffffffffffffffffff166108fc8484848181106120ca57fe5b905060200201359081150290604051600060405180830381858888f193505050501580156120fc573d6000803e3d6000fd5b508080600101915050612079565b5050505050565b600d6020528060005260406000206000915054906101000a900467ffffffffffffffff1681565b612140613433565b80600881905550807fa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c860405160405180910390a250565b6010602052816000526040600020602052806000526040600020600091509150508060000154908060010154908060040160009054906101000a900460ff16908060050154905084565b60006121ed7fe2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc4836127b5565b9050919050565b6000600560009054906101000a900460ff16905090565b612213613433565b6122206000801b82611a30565b61222d6000801b33611ac8565b50565b606061225a82604051602001612246919061517a565b604051602081830303815290604052612436565b9050919050565b60046020528060005260406000206000915054906101000a900460ff1681565b60005b602460ff168160ff16101561230e576122b2828260ff16815181106122a557fe5b6020026020010151613481565b60128260ff16602481106122c257fe5b0160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508080600101915050612284565b5050565b60036020528060005260406000206000915054906101000a900460ff1681565b6011602052826000526040600020602052816000526040600020602052806000526040600020600092509250509054906101000a900460ff1681565b60095481565b61237c613433565b61238461349d565b565b600e6020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6123c1613433565b60008290508073ffffffffffffffffffffffffffffffffffffffff166307b7ed99836040518263ffffffff1660e01b81526004016123ff9190615234565b600060405180830381600087803b15801561241957600080fd5b505af115801561242d573d6000803e3d6000fd5b50505050505050565b6060806040518060400160405280601081526020017f303132333435363738396162636465660000000000000000000000000000000081525090506060600284510260020167ffffffffffffffff8111801561249157600080fd5b506040519080825280601f01601f1916602001820160405280156124c45781602001600182028036833780820191505090505b5090507f3000000000000000000000000000000000000000000000000000000000000000816000815181106124f557fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f78000000000000000000000000000000000000000000000000000000000000008160018151811061255257fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060005b84518110156126ae5782600486838151811061259c57fe5b602001015160f81c60f81b7effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916901c60f81c60ff16815181106125db57fe5b602001015160f81c60f81b8260028302600201815181106125f857fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535082600f60f81b86838151811061263957fe5b602001015160f81c60f81b1660f81c60ff168151811061265557fe5b602001015160f81c60f81b82600283026003018151811061267257fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080600101915050612584565b508092505050919050565b600060606040518060400160405280601c81526020017f19457468657265756d205369676e6564204d6573736167653a0a33320000000081525090506000818760405160200161270a9291906151f1565b60405160208183030381529060405280519060200120905060006001828888886040516000815260200160405260405161274794939291906153be565b6020604051602081039080840390855afa158015612769573d6000803e3d6000fd5b505050602060405103519050809350505050949350505050565b60006127ad82600660008681526020019081526020016000206000016134f990919063ffffffff16565b905092915050565b60006127df8260066000868152602001908152602001600020600001612f6c90919063ffffffff16565b905092915050565b6060612811826040516020016127fd91906151bf565b604051602081830303815290604052612436565b9050919050565b7fe2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc481565b600a5481565b61284a613433565b6128747fe2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc4826127b5565b6128b3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016128aa906155a3565b60405180910390fd5b6128dd7fe2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc482612da2565b8073ffffffffffffffffffffffffffffffffffffffff167f10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b60405160405180910390a26009600081548092919060019003919050555050565b6000801b81565b612945613a55565b60008460ff1660088567ffffffffffffffff1668ffffffffffffffffff16901b179050601060008268ffffffffffffffffff1668ffffffffffffffffff16815260200190815260200160002060008481526020019081526020016000206040518060c0016040529081600082015481526020016001820154815260200160028201805480602002602001604051908101604052809291908181526020018280548015612a4657602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116129fc575b5050505050815260200160038201805480602002602001604051908101604052809291908181526020018280548015612ad457602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311612a8a575b505050505081526020016004820160009054906101000a900460ff166004811115612afb57fe5b6004811115612b0657fe5b81526020016005820154815250509150509392505050565b6000612b6083836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250613513565b905092915050565b612b70612fee565b612b7a828261356e565b5050565b600760009054906101000a900460ff1681565b600b5481565b600c5481565b60026020528060005260406000206000915090505481565b6000612bd560066000848152602001908152602001600020600001613660565b9050919050565b612be4613433565b82600e600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008390508073ffffffffffffffffffffffffffffffffffffffff1663b8fa373684846040518363ffffffff1660e01b8152600401612c769291906152a0565b600060405180830381600087803b158015612c9057600080fd5b505af1158015612ca4573d6000803e3d6000fd5b5050505050505050565b612cb6613433565b612ce07fe2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc4826127b5565b15612d20576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d17906157c3565b60405180910390fd5b612d4a7fe2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc482611a30565b8073ffffffffffffffffffffffffffffffffffffffff167f03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c560405160405180910390a260096000815480929190600101919050555050565b612dc96006600084815260200190815260200160002060020154612dc46131d0565b6127b5565b612e08576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612dff90615703565b60405180910390fd5b612e12828261326c565b5050565b60085481565b505050565b612e29613433565b84600e600086815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060008590508073ffffffffffffffffffffffffffffffffffffffff1663bba8185a868686866040518563ffffffff1660e01b8152600401612ebf94939291906152c9565b600060405180830381600087803b158015612ed957600080fd5b505af1158015612eed573d6000803e3d6000fd5b50505050505050505050565b6060612f2382604051602001612f0f9190615219565b604051602081830303815290604052612436565b9050919050565b612f32613433565b612f3a613675565b565b6000612f64836000018373ffffffffffffffffffffffffffffffffffffffff1660001b6136d1565b905092915050565b6000612f94836000018373ffffffffffffffffffffffffffffffffffffffff1660001b613741565b905092915050565b600560009054906101000a900460ff1615612fec576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612fe390615743565b60405180910390fd5b565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461307c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161307390615563565b60405180910390fd5b565b600360008273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff1661310a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161310190615643565b60405180910390fd5b6001600460008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff02191690831515021790555050565b61318f7fe2b7fb3b832174769106daebcfd6d1970523240dda11281102db9363b83b0dc4336127b5565b6131ce576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016131c590615823565b60405180910390fd5b565b600033905090565b6132008160066000858152602001908152602001600020600001612f3c90919063ffffffff16565b156132685761320d6131d0565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a45b5050565b613294816006600085815260200190815260200160002060000161376490919063ffffffff16565b156132fc576132a16131d0565b73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16837ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b60405160405180910390a45b5050565b60008273ffffffffffffffffffffffffffffffffffffffff1682600067ffffffffffffffff8111801561333257600080fd5b506040519080825280601f01601f1916602001820160405280156133655781602001600182028036833780820191505090505b5060405161337391906151da565b60006040518083038185875af1925050503d80600081146133b0576040519150601f19603f3d011682016040523d82523d6000602084013e6133b5565b606091505b50509050806133f8577fa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f6040516133eb90615723565b60405180910390a161342e565b7fa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f604051613425906155c3565b60405180910390a15b505050565b6134406000801b336127b5565b61347f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161347690615843565b60405180910390fd5b565b6000808280519060200120905080600052600051915050919050565b6134a5612f9c565b6001600560006101000a81548160ff0219169083151502179055507f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258336040516134ef919061524f565b60405180910390a1565b60006135088360000183613794565b60001c905092915050565b600083831115829061355b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016135529190615481565b60405180910390fd5b5060008385039050809150509392505050565b806001600084815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055506001600360008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055505050565b600061366e82600001613801565b9050919050565b61367d613812565b6000600560006101000a81548160ff0219169083151502179055507f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa336040516136c7919061524f565b60405180910390a1565b60006136dd8383613741565b61373657826000018290806001815401808255809150506001900390600052602060002001600090919091909150558260000180549050836001016000848152602001908152602001600020819055506001905061373b565b600090505b92915050565b600080836001016000848152602001908152602001600020541415905092915050565b600061378c836000018373ffffffffffffffffffffffffffffffffffffffff1660001b613863565b905092915050565b6000818360000180549050116137df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016137d6906154e3565b60405180910390fd5b8260000182815481106137ee57fe5b9060005260206000200154905092915050565b600081600001805490509050919050565b600560009054906101000a900460ff16613861576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161385890615543565b60405180910390fd5b565b6000808360010160008481526020019081526020016000205490506000811461393f57600060018203905060006001866000018054905003905060008660000182815481106138ae57fe5b90600052602060002001549050808760000184815481106138cb57fe5b906000526020600020018190555060018301876001016000838152602001908152602001600020819055508660000180548061390357fe5b60019003818190600052602060002001600090559055866001016000878152602001908152602001600020600090556001945050505050613945565b60009150505b92915050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061398c57803560ff19168380011785556139ba565b828001600101855582156139ba579182015b828111156139b957823582559160200191906001019061399e565b5b5090506139c79190613a9c565b5090565b828054828255906000526020600020908101928215613a44579160200282015b82811115613a435782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550916020019190600101906139eb565b5b509050613a519190613ab9565b5090565b6040518060c001604052806000801916815260200160008019168152602001606081526020016060815260200160006004811115613a8f57fe5b8152602001600081525090565b5b80821115613ab5576000816000905550600101613a9d565b5090565b5b80821115613af057600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101613aba565b5090565b600081359050613b0381615cc0565b92915050565b600081519050613b1881615cc0565b92915050565b600081359050613b2d81615cd7565b92915050565b60008083601f840112613b4557600080fd5b8235905067ffffffffffffffff811115613b5e57600080fd5b602083019150836020820283011115613b7657600080fd5b9250929050565b600082601f830112613b8e57600080fd5b8135613ba1613b9c826159a3565b615976565b9150818183526020840193506020810190508360005b83811015613be75781358601613bcd8882613cc4565b845260208401935060208301925050600181019050613bb7565b5050505092915050565b60008083601f840112613c0357600080fd5b8235905067ffffffffffffffff811115613c1c57600080fd5b602083019150836020820283011115613c3457600080fd5b9250929050565b600081519050613c4a81615cee565b92915050565b600081359050613c5f81615d05565b92915050565b600081359050613c7481615d1c565b92915050565b60008083601f840112613c8c57600080fd5b8235905067ffffffffffffffff811115613ca557600080fd5b602083019150836001820283011115613cbd57600080fd5b9250929050565b600082601f830112613cd557600080fd5b8135613ce8613ce3826159cb565b615976565b91508082526020830160208301858383011115613d0457600080fd5b613d0f838284615be1565b50505092915050565b600081359050613d2781615d33565b92915050565b600081519050613d3c81615d33565b92915050565b600081359050613d5181615d4a565b92915050565b600081359050613d6681615d61565b92915050565b600081359050613d7b81615d78565b92915050565b600060208284031215613d9357600080fd5b6000613da184828501613af4565b91505092915050565b600060208284031215613dbc57600080fd5b6000613dca84828501613b1e565b91505092915050565b60008060408385031215613de657600080fd5b6000613df485828601613af4565b9250506020613e0585828601613af4565b9150509250929050565b600080600060608486031215613e2457600080fd5b6000613e3286828701613af4565b9350506020613e4386828701613af4565b9250506040613e5486828701613d18565b9150509250925092565b600080600060608486031215613e7357600080fd5b6000613e8186828701613af4565b9350506020613e9286828701613c50565b9250506040613ea386828701613af4565b9150509250925092565b600080600080600060a08688031215613ec557600080fd5b6000613ed388828901613af4565b9550506020613ee488828901613c50565b9450506040613ef588828901613af4565b9350506060613f0688828901613c65565b9250506080613f1788828901613c65565b9150509295509295909350565b60008060008060408587031215613f3a57600080fd5b600085013567ffffffffffffffff811115613f5457600080fd5b613f6087828801613b33565b9450945050602085013567ffffffffffffffff811115613f7f57600080fd5b613f8b87828801613bf1565b925092505092959194509250565b600060208284031215613fab57600080fd5b600082013567ffffffffffffffff811115613fc557600080fd5b613fd184828501613b7d565b91505092915050565b60008060008060808587031215613ff057600080fd5b6000613ffe87828801613c3b565b945050602061400f87828801613c3b565b935050604061402087828801613b09565b925050606061403187828801613d2d565b91505092959194509250565b60006020828403121561404f57600080fd5b600061405d84828501613c50565b91505092915050565b6000806040838503121561407957600080fd5b600061408785828601613c50565b925050602061409885828601613af4565b9150509250929050565b600080604083850312156140b557600080fd5b60006140c385828601613c50565b92505060206140d485828601613d18565b9150509250929050565b600080600080608085870312156140f457600080fd5b600061410287828801613c50565b945050602061411387828801613d6c565b935050604061412487828801613c50565b925050606061413587828801613c50565b91505092959194509250565b60006020828403121561415357600080fd5b600082013567ffffffffffffffff81111561416d57600080fd5b61417984828501613cc4565b91505092915050565b60006020828403121561419457600080fd5b60006141a284828501613d18565b91505092915050565b600080604083850312156141be57600080fd5b60006141cc85828601613d18565b92505060206141dd85828601613d18565b9150509250929050565b600080604083850312156141fa57600080fd5b600061420885828601613d42565b925050602061421985828601613d6c565b9150509250929050565b6000806040838503121561423657600080fd5b600061424485828601613d57565b925050602061425585828601613c50565b9150509250929050565b60008060006060848603121561427457600080fd5b600061428286828701613d57565b935050602061429386828701613c50565b92505060406142a486828701613af4565b9150509250925092565b6000602082840312156142c057600080fd5b60006142ce84828501613d6c565b91505092915050565b600080600080606085870312156142ed57600080fd5b60006142fb87828801613d6c565b945050602061430c87828801613c50565b935050604085013567ffffffffffffffff81111561432957600080fd5b61433587828801613c7a565b925092505092959194509250565b60008060006060848603121561435857600080fd5b600061436686828701613d6c565b935050602061437786828701613d42565b925050604061438886828701613c50565b9150509250925092565b600080600080608085870312156143a857600080fd5b60006143b687828801613d6c565b94505060206143c787828801613d42565b93505060406143d887828801613c50565b92505060606143e987828801613c50565b91505092959194509250565b60008060008060006080868803121561440d57600080fd5b600061441b88828901613d6c565b955050602061442c88828901613d42565b945050604086013567ffffffffffffffff81111561444957600080fd5b61445588828901613c7a565b9350935050606061446888828901613c50565b9150509295509295909350565b6000614481838361449c565b60208301905092915050565b61449681615b99565b82525050565b6144a581615ac0565b82525050565b6144b481615ac0565b82525050565b6144cb6144c682615ac0565b615c3d565b82525050565b6144da81615a11565b6144e48184615a57565b92506144ef826159f7565b8060005b838110156145275761450482615c75565b61450e8782614475565b965061451983615a3d565b9250506001810190506144f3565b505050505050565b600061453a82615a1c565b6145448185615a62565b935061454f83615a01565b8060005b838110156145805781516145678882614475565b975061457283615a4a565b925050600181019050614553565b5085935050505092915050565b61459681615ae4565b82525050565b6145a581615af0565b82525050565b6145b481615af0565b82525050565b6145cb6145c682615af0565b615c4f565b82525050565b6145da81615afa565b82525050565b60006145ec8385615a73565b93506145f9838584615be1565b61460283615c88565b840190509392505050565b60006146198385615a84565b9350614626838584615be1565b82840190509392505050565b600061463d82615a27565b6146478185615a73565b9350614657818560208601615bf0565b61466081615c88565b840191505092915050565b600061467682615a27565b6146808185615a84565b9350614690818560208601615bf0565b80840191505092915050565b6146a581615bab565b82525050565b6146b481615bab565b82525050565b60006146c582615a32565b6146cf8185615a8f565b93506146df818560208601615bf0565b6146e881615c88565b840191505092915050565b6000614700601c83615a8f565b91507f70726f706f73616c20616c7265616479207472616e73666572726564000000006000830152602082019050919050565b6000614740601683615a8f565b91507f70726f706f73616c206973206e6f7420616374697665000000000000000000006000830152602082019050919050565b6000614780602283615a8f565b91507f456e756d657261626c655365743a20696e646578206f7574206f6620626f756e60008301527f64730000000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006147e6602f83615a8f565b91507f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60008301527f2061646d696e20746f206772616e7400000000000000000000000000000000006020830152604082019050919050565b600061484c601583615a8f565b91507f72656c6179657220616c726561647920766f74656400000000000000000000006000830152602082019050919050565b600061488c601483615a8f565b91507f5061757361626c653a206e6f74207061757365640000000000000000000000006000830152602082019050919050565b60006148cc601e83615a8f565b91507f73656e646572206d7573742062652062726964676520636f6e747261637400006000830152602082019050919050565b600061490c601f83615a8f565b91507f636f6d6520746f20627269646765206578656375746550726f706f73616c20006000830152602082019050919050565b600061494c601f83615a8f565b91507f6164647220646f65736e277420686176652072656c6179657220726f6c6521006000830152602082019050919050565b600061498c600e83615a8f565b91507f746f6b656e2063616c6c204f4b200000000000000000000000000000000000006000830152602082019050919050565b60006149cc601683615a8f565b91507f70726f706f73616c206e6f6e6365416e644944203a20000000000000000000006000830152602082019050919050565b6000614a0c601b83615a8f565b91507f636f6d652070726f706f73616c2e5f737461747573203d3d20312000000000006000830152602082019050919050565b6000614a4c602083615a8f565b91507f78786c206578656375746550726f706f73616c20766572696669656420656e646000830152602082019050919050565b6000614a8c602483615a8f565b91507f70726f766964656420636f6e7472616374206973206e6f742077686974656c6960008301527f73746564000000000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000614af2601683615a8f565b91507f636f6d6520746f20627269646765206465706f736974000000000000000000006000830152602082019050919050565b6000614b32601483615a8f565b91507f70726f706f73616c206461746148617368203a200000000000000000000000006000830152602082019050919050565b6000614b72601b83615a8f565b91507f6461746120646f65736e2774206d6174636820646174616861736800000000006000830152602082019050919050565b6000614bb2601d83615a8f565b91507f636f6d6520746f2062726964676520766f746550726f706f73616c20320000006000830152602082019050919050565b6000614bf2602083615a8f565b91507f7265736f757263654944206e6f74206d617070656420746f2068616e646c65726000830152602082019050919050565b6000614c32603083615a8f565b91507f416363657373436f6e74726f6c3a2073656e646572206d75737420626520616e60008301527f2061646d696e20746f207265766f6b65000000000000000000000000000000006020830152604082019050919050565b6000614c98601283615a8f565b91507f746f6b656e2063616c6c206661696c65642000000000000000000000000000006000830152602082019050919050565b6000614cd8601083615a8f565b91507f5061757361626c653a20706175736564000000000000000000000000000000006000830152602082019050919050565b6000614d18601b83615a8f565b91507f636f6d6520746f2062726964676520766f746550726f706f73616c00000000006000830152602082019050919050565b6000614d58601d83615a8f565b91507f636f6d6520746f2062726964676520766f746550726f706f73616c20310000006000830152602082019050919050565b6000614d98601c83615a8f565b91507f78786c20636f6d6520746f205f736166655472616e73666572455448000000006000830152602082019050919050565b6000614dd8601e83615a8f565b91507f6164647220616c7265616479206861732072656c6179657220726f6c652100006000830152602082019050919050565b6000614e18601883615a8f565b91507f627269646765206465706f736974206461746120656e642000000000000000006000830152602082019050919050565b6000614e58602a83615a8f565b91507f70726f706f73616c20616c7265616479207061737365642f657865637574656460008301527f2f63616e63656c6c6564000000000000000000000000000000000000000000006020830152604082019050919050565b6000614ebe602083615a8f565b91507f73656e64657220646f65736e277420686176652072656c6179657220726f6c656000830152602082019050919050565b6000614efe601e83615a8f565b91507f73656e64657220646f65736e277420686176652061646d696e20726f6c6500006000830152602082019050919050565b6000614f3e601283615a8f565b91507f70726f706f73616c20737461747573203a2000000000000000000000000000006000830152602082019050919050565b6000614f7e601c83615a8f565b91507f636f6d6520746f20627269646765207472616e7366657246756e6473000000006000830152602082019050919050565b6000614fbe601183615a8f565b91507f6461746168617368206d69736d617463680000000000000000000000000000006000830152602082019050919050565b6000614ffe601983615a8f565b91507f6e6f2068616e646c657220666f72207265736f757263654944000000000000006000830152602082019050919050565b600061503e602f83615a8f565b91507f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560008301527f20726f6c657320666f722073656c6600000000000000000000000000000000006020830152604082019050919050565b600060c0830160008301516150af600086018261459c565b5060208301516150c2602086018261459c565b50604083015184820360408601526150da828261452f565b915050606083015184820360608601526150f4828261452f565b9150506080830151615109608086018261469c565b5060a083015161511c60a0860182615127565b508091505092915050565b61513081615b59565b82525050565b61513f81615b59565b82525050565b61515661515182615b59565b615c6b565b82525050565b61516581615b63565b82525050565b61517481615b8c565b82525050565b600061518682846144ba565b60148201915081905092915050565b60006151a182866144ba565b6014820191506151b282848661460d565b9150819050949350505050565b60006151cb82846145ba565b60208201915081905092915050565b60006151e6828461466b565b915081905092915050565b60006151fd828561466b565b915061520982846145ba565b6020820191508190509392505050565b60006152258284615145565b60208201915081905092915050565b600060208201905061524960008301846144ab565b92915050565b6000602082019050615264600083018461448d565b92915050565b600060208201905061527f600083018461458d565b92915050565b600060208201905061529a60008301846145ab565b92915050565b60006040820190506152b560008301856145ab565b6152c260208301846144ab565b9392505050565b60006080820190506152de60008301876145ab565b6152eb60208301866144ab565b6152f860408301856145d1565b61530560608301846145d1565b95945050505050565b60006104c08201905061532460008301876145ab565b61533160208301866144d1565b8181036104a08301526153458184866145e0565b905095945050505050565b600060408201905061536560008301856145ab565b61537260208301846145ab565b9392505050565b600060808201905061538e60008301876145ab565b61539b60208301866145ab565b6153a860408301856146ab565b6153b56060830184615136565b95945050505050565b60006080820190506153d360008301876145ab565b6153e0602083018661516b565b6153ed60408301856145ab565b6153fa60608301846145ab565b95945050505050565b600060a08201905061541860008301896145ab565b615425602083018861516b565b615432604083018761515c565b61543f606083018661448d565b81810360808301526154528184866145e0565b9050979650505050505050565b600060208201905081810360008301526154798184614632565b905092915050565b6000602082019050818103600083015261549b81846146ba565b905092915050565b600060208201905081810360008301526154bc816146f3565b9050919050565b600060208201905081810360008301526154dc81614733565b9050919050565b600060208201905081810360008301526154fc81614773565b9050919050565b6000602082019050818103600083015261551c816147d9565b9050919050565b6000602082019050818103600083015261553c8161483f565b9050919050565b6000602082019050818103600083015261555c8161487f565b9050919050565b6000602082019050818103600083015261557c816148bf565b9050919050565b6000602082019050818103600083015261559c816148ff565b9050919050565b600060208201905081810360008301526155bc8161493f565b9050919050565b600060208201905081810360008301526155dc8161497f565b9050919050565b600060208201905081810360008301526155fc816149bf565b9050919050565b6000602082019050818103600083015261561c816149ff565b9050919050565b6000602082019050818103600083015261563c81614a3f565b9050919050565b6000602082019050818103600083015261565c81614a7f565b9050919050565b6000602082019050818103600083015261567c81614ae5565b9050919050565b6000602082019050818103600083015261569c81614b25565b9050919050565b600060208201905081810360008301526156bc81614b65565b9050919050565b600060208201905081810360008301526156dc81614ba5565b9050919050565b600060208201905081810360008301526156fc81614be5565b9050919050565b6000602082019050818103600083015261571c81614c25565b9050919050565b6000602082019050818103600083015261573c81614c8b565b9050919050565b6000602082019050818103600083015261575c81614ccb565b9050919050565b6000602082019050818103600083015261577c81614d0b565b9050919050565b6000602082019050818103600083015261579c81614d4b565b9050919050565b600060208201905081810360008301526157bc81614d8b565b9050919050565b600060208201905081810360008301526157dc81614dcb565b9050919050565b600060208201905081810360008301526157fc81614e0b565b9050919050565b6000602082019050818103600083015261581c81614e4b565b9050919050565b6000602082019050818103600083015261583c81614eb1565b9050919050565b6000602082019050818103600083015261585c81614ef1565b9050919050565b6000602082019050818103600083015261587c81614f31565b9050919050565b6000602082019050818103600083015261589c81614f71565b9050919050565b600060208201905081810360008301526158bc81614fb1565b9050919050565b600060208201905081810360008301526158dc81614ff1565b9050919050565b600060208201905081810360008301526158fc81615031565b9050919050565b6000602082019050818103600083015261591d8184615097565b905092915050565b600060208201905061593a6000830184615136565b92915050565b6000602082019050615955600083018461515c565b92915050565b6000602082019050615970600083018461516b565b92915050565b6000604051905081810181811067ffffffffffffffff8211171561599957600080fd5b8060405250919050565b600067ffffffffffffffff8211156159ba57600080fd5b602082029050602081019050919050565b600067ffffffffffffffff8211156159e257600080fd5b601f19601f8301169050602081019050919050565b6000819050919050565b6000819050602082019050919050565b600060249050919050565b600081519050919050565b600081519050919050565b600081519050919050565b6000600182019050919050565b6000602082019050919050565b600081905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600082825260208201905092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000615acb82615b39565b9050919050565b6000615add82615b39565b9050919050565b60008115159050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b6000819050615b3482615cb3565b919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600067ffffffffffffffff82169050919050565b600068ffffffffffffffffff82169050919050565b600060ff82169050919050565b6000615ba482615bbd565b9050919050565b6000615bb682615b26565b9050919050565b6000615bc882615bcf565b9050919050565b6000615bda82615b39565b9050919050565b82818337600083830152505050565b60005b83811015615c0e578082015181840152602081019050615bf3565b83811115615c1d576000848401525b50505050565b6000615c36615c3183615ca6565b615aa0565b9050919050565b6000615c4882615c59565b9050919050565b6000819050919050565b6000615c6482615c99565b9050919050565b6000819050919050565b6000615c818254615c23565b9050919050565b6000601f19601f8301169050919050565b60008160601b9050919050565b60008160001c9050919050565b60058110615cbd57fe5b50565b615cc981615ac0565b8114615cd457600080fd5b50565b615ce081615ad2565b8114615ceb57600080fd5b50565b615cf781615ae4565b8114615d0257600080fd5b50565b615d0e81615af0565b8114615d1957600080fd5b50565b615d2581615afa565b8114615d3057600080fd5b50565b615d3c81615b59565b8114615d4757600080fd5b50565b615d5381615b63565b8114615d5e57600080fd5b50565b615d6a81615b77565b8114615d7557600080fd5b50565b615d8181615b8c565b8114615d8c57600080fd5b5056fea2646970667358221220d1ef541bece1e03ce022258269af60aa0d5cb023df58595c805d5e9ed155e1cc64736f6c634300060c0033"

// DeployBridge deploys a new Ethereum contract, binding an instance of Bridge to it.
func DeployBridge(auth *bind.TransactOpts, backend bind.ContractBackend, chainID uint8, initialRelayers []common.Address, initialRelayerThreshold *big.Int, fee *big.Int, expiry *big.Int) (common.Address, *types.Transaction, *Bridge, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BridgeBin), backend, chainID, initialRelayers, initialRelayerThreshold, fee, expiry)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BridgeABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bridge *BridgeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bridge *BridgeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bridge.Contract.DEFAULTADMINROLE(&_Bridge.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bridge *BridgeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bridge.Contract.DEFAULTADMINROLE(&_Bridge.CallOpts)
}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_Bridge *BridgeCaller) RELAYERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "RELAYER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_Bridge *BridgeSession) RELAYERROLE() ([32]byte, error) {
	return _Bridge.Contract.RELAYERROLE(&_Bridge.CallOpts)
}

// RELAYERROLE is a free data retrieval call binding the contract method 0x926d7d7f.
//
// Solidity: function RELAYER_ROLE() view returns(bytes32)
func (_Bridge *BridgeCallerSession) RELAYERROLE() ([32]byte, error) {
	return _Bridge.Contract.RELAYERROLE(&_Bridge.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_Bridge *BridgeCaller) BridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_bridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_Bridge *BridgeSession) BridgeAddress() (common.Address, error) {
	return _Bridge.Contract.BridgeAddress(&_Bridge.CallOpts)
}

// BridgeAddress is a free data retrieval call binding the contract method 0x318c136e.
//
// Solidity: function _bridgeAddress() view returns(address)
func (_Bridge *BridgeCallerSession) BridgeAddress() (common.Address, error) {
	return _Bridge.Contract.BridgeAddress(&_Bridge.CallOpts)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_Bridge *BridgeCaller) BurnList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_burnList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_Bridge *BridgeSession) BurnList(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.BurnList(&_Bridge.CallOpts, arg0)
}

// BurnList is a free data retrieval call binding the contract method 0x6a70d081.
//
// Solidity: function _burnList(address ) view returns(bool)
func (_Bridge *BridgeCallerSession) BurnList(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.BurnList(&_Bridge.CallOpts, arg0)
}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() view returns(uint8)
func (_Bridge *BridgeCaller) ChainID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_chainID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() view returns(uint8)
func (_Bridge *BridgeSession) ChainID() (uint8, error) {
	return _Bridge.Contract.ChainID(&_Bridge.CallOpts)
}

// ChainID is a free data retrieval call binding the contract method 0xbeab7131.
//
// Solidity: function _chainID() view returns(uint8)
func (_Bridge *BridgeCallerSession) ChainID() (uint8, error) {
	return _Bridge.Contract.ChainID(&_Bridge.CallOpts)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_Bridge *BridgeCaller) ContractWhitelist(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_contractWhitelist", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_Bridge *BridgeSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.ContractWhitelist(&_Bridge.CallOpts, arg0)
}

// ContractWhitelist is a free data retrieval call binding the contract method 0x7f79bea8.
//
// Solidity: function _contractWhitelist(address ) view returns(bool)
func (_Bridge *BridgeCallerSession) ContractWhitelist(arg0 common.Address) (bool, error) {
	return _Bridge.Contract.ContractWhitelist(&_Bridge.CallOpts, arg0)
}

// DepositCounts is a free data retrieval call binding the contract method 0x4b0b919d.
//
// Solidity: function _depositCounts(uint8 ) view returns(uint64)
func (_Bridge *BridgeCaller) DepositCounts(opts *bind.CallOpts, arg0 uint8) (uint64, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_depositCounts", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DepositCounts is a free data retrieval call binding the contract method 0x4b0b919d.
//
// Solidity: function _depositCounts(uint8 ) view returns(uint64)
func (_Bridge *BridgeSession) DepositCounts(arg0 uint8) (uint64, error) {
	return _Bridge.Contract.DepositCounts(&_Bridge.CallOpts, arg0)
}

// DepositCounts is a free data retrieval call binding the contract method 0x4b0b919d.
//
// Solidity: function _depositCounts(uint8 ) view returns(uint64)
func (_Bridge *BridgeCallerSession) DepositCounts(arg0 uint8) (uint64, error) {
	return _Bridge.Contract.DepositCounts(&_Bridge.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0x3ee7094a.
//
// Solidity: function _depositRecords(uint64 , uint8 ) view returns(bytes)
func (_Bridge *BridgeCaller) DepositRecords(opts *bind.CallOpts, arg0 uint64, arg1 uint8) ([]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_depositRecords", arg0, arg1)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// DepositRecords is a free data retrieval call binding the contract method 0x3ee7094a.
//
// Solidity: function _depositRecords(uint64 , uint8 ) view returns(bytes)
func (_Bridge *BridgeSession) DepositRecords(arg0 uint64, arg1 uint8) ([]byte, error) {
	return _Bridge.Contract.DepositRecords(&_Bridge.CallOpts, arg0, arg1)
}

// DepositRecords is a free data retrieval call binding the contract method 0x3ee7094a.
//
// Solidity: function _depositRecords(uint64 , uint8 ) view returns(bytes)
func (_Bridge *BridgeCallerSession) DepositRecords(arg0 uint64, arg1 uint8) ([]byte, error) {
	return _Bridge.Contract.DepositRecords(&_Bridge.CallOpts, arg0, arg1)
}

// Expiry is a free data retrieval call binding the contract method 0xc5ec8970.
//
// Solidity: function _expiry() view returns(uint256)
func (_Bridge *BridgeCaller) Expiry(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_expiry")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Expiry is a free data retrieval call binding the contract method 0xc5ec8970.
//
// Solidity: function _expiry() view returns(uint256)
func (_Bridge *BridgeSession) Expiry() (*big.Int, error) {
	return _Bridge.Contract.Expiry(&_Bridge.CallOpts)
}

// Expiry is a free data retrieval call binding the contract method 0xc5ec8970.
//
// Solidity: function _expiry() view returns(uint256)
func (_Bridge *BridgeCallerSession) Expiry() (*big.Int, error) {
	return _Bridge.Contract.Expiry(&_Bridge.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_Bridge *BridgeCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_Bridge *BridgeSession) Fee() (*big.Int, error) {
	return _Bridge.Contract.Fee(&_Bridge.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xc5b37c22.
//
// Solidity: function _fee() view returns(uint256)
func (_Bridge *BridgeCallerSession) Fee() (*big.Int, error) {
	return _Bridge.Contract.Fee(&_Bridge.CallOpts)
}

// HasVotedOnProposal is a free data retrieval call binding the contract method 0x7febe63f.
//
// Solidity: function _hasVotedOnProposal(uint72 , bytes32 , address ) view returns(bool)
func (_Bridge *BridgeCaller) HasVotedOnProposal(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte, arg2 common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_hasVotedOnProposal", arg0, arg1, arg2)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVotedOnProposal is a free data retrieval call binding the contract method 0x7febe63f.
//
// Solidity: function _hasVotedOnProposal(uint72 , bytes32 , address ) view returns(bool)
func (_Bridge *BridgeSession) HasVotedOnProposal(arg0 *big.Int, arg1 [32]byte, arg2 common.Address) (bool, error) {
	return _Bridge.Contract.HasVotedOnProposal(&_Bridge.CallOpts, arg0, arg1, arg2)
}

// HasVotedOnProposal is a free data retrieval call binding the contract method 0x7febe63f.
//
// Solidity: function _hasVotedOnProposal(uint72 , bytes32 , address ) view returns(bool)
func (_Bridge *BridgeCallerSession) HasVotedOnProposal(arg0 *big.Int, arg1 [32]byte, arg2 common.Address) (bool, error) {
	return _Bridge.Contract.HasVotedOnProposal(&_Bridge.CallOpts, arg0, arg1, arg2)
}

// Proposals is a free data retrieval call binding the contract method 0x50598719.
//
// Solidity: function _proposals(uint72 , bytes32 ) view returns(bytes32 _resourceID, bytes32 _dataHash, uint8 _status, uint256 _proposedBlock)
func (_Bridge *BridgeCaller) Proposals(opts *bind.CallOpts, arg0 *big.Int, arg1 [32]byte) (struct {
	ResourceID    [32]byte
	DataHash      [32]byte
	Status        uint8
	ProposedBlock *big.Int
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_proposals", arg0, arg1)

	outstruct := new(struct {
		ResourceID    [32]byte
		DataHash      [32]byte
		Status        uint8
		ProposedBlock *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ResourceID = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.DataHash = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)
	outstruct.Status = *abi.ConvertType(out[2], new(uint8)).(*uint8)
	outstruct.ProposedBlock = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Proposals is a free data retrieval call binding the contract method 0x50598719.
//
// Solidity: function _proposals(uint72 , bytes32 ) view returns(bytes32 _resourceID, bytes32 _dataHash, uint8 _status, uint256 _proposedBlock)
func (_Bridge *BridgeSession) Proposals(arg0 *big.Int, arg1 [32]byte) (struct {
	ResourceID    [32]byte
	DataHash      [32]byte
	Status        uint8
	ProposedBlock *big.Int
}, error) {
	return _Bridge.Contract.Proposals(&_Bridge.CallOpts, arg0, arg1)
}

// Proposals is a free data retrieval call binding the contract method 0x50598719.
//
// Solidity: function _proposals(uint72 , bytes32 ) view returns(bytes32 _resourceID, bytes32 _dataHash, uint8 _status, uint256 _proposedBlock)
func (_Bridge *BridgeCallerSession) Proposals(arg0 *big.Int, arg1 [32]byte) (struct {
	ResourceID    [32]byte
	DataHash      [32]byte
	Status        uint8
	ProposedBlock *big.Int
}, error) {
	return _Bridge.Contract.Proposals(&_Bridge.CallOpts, arg0, arg1)
}

// RecoverSigner2 is a free data retrieval call binding the contract method 0x8f667892.
//
// Solidity: function _recoverSigner2(bytes32 h, uint8 v, bytes32 r, bytes32 s) pure returns(address)
func (_Bridge *BridgeCaller) RecoverSigner2(opts *bind.CallOpts, h [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_recoverSigner2", h, v, r, s)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RecoverSigner2 is a free data retrieval call binding the contract method 0x8f667892.
//
// Solidity: function _recoverSigner2(bytes32 h, uint8 v, bytes32 r, bytes32 s) pure returns(address)
func (_Bridge *BridgeSession) RecoverSigner2(h [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _Bridge.Contract.RecoverSigner2(&_Bridge.CallOpts, h, v, r, s)
}

// RecoverSigner2 is a free data retrieval call binding the contract method 0x8f667892.
//
// Solidity: function _recoverSigner2(bytes32 h, uint8 v, bytes32 r, bytes32 s) pure returns(address)
func (_Bridge *BridgeCallerSession) RecoverSigner2(h [32]byte, v uint8, r [32]byte, s [32]byte) (common.Address, error) {
	return _Bridge.Contract.RecoverSigner2(&_Bridge.CallOpts, h, v, r, s)
}

// RelayerThreshold is a free data retrieval call binding the contract method 0xd7a9cd79.
//
// Solidity: function _relayerThreshold() view returns(uint256)
func (_Bridge *BridgeCaller) RelayerThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_relayerThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RelayerThreshold is a free data retrieval call binding the contract method 0xd7a9cd79.
//
// Solidity: function _relayerThreshold() view returns(uint256)
func (_Bridge *BridgeSession) RelayerThreshold() (*big.Int, error) {
	return _Bridge.Contract.RelayerThreshold(&_Bridge.CallOpts)
}

// RelayerThreshold is a free data retrieval call binding the contract method 0xd7a9cd79.
//
// Solidity: function _relayerThreshold() view returns(uint256)
func (_Bridge *BridgeCallerSession) RelayerThreshold() (*big.Int, error) {
	return _Bridge.Contract.RelayerThreshold(&_Bridge.CallOpts)
}

// ResourceIDToHandlerAddress is a free data retrieval call binding the contract method 0x84db809f.
//
// Solidity: function _resourceIDToHandlerAddress(bytes32 ) view returns(address)
func (_Bridge *BridgeCaller) ResourceIDToHandlerAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_resourceIDToHandlerAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToHandlerAddress is a free data retrieval call binding the contract method 0x84db809f.
//
// Solidity: function _resourceIDToHandlerAddress(bytes32 ) view returns(address)
func (_Bridge *BridgeSession) ResourceIDToHandlerAddress(arg0 [32]byte) (common.Address, error) {
	return _Bridge.Contract.ResourceIDToHandlerAddress(&_Bridge.CallOpts, arg0)
}

// ResourceIDToHandlerAddress is a free data retrieval call binding the contract method 0x84db809f.
//
// Solidity: function _resourceIDToHandlerAddress(bytes32 ) view returns(address)
func (_Bridge *BridgeCallerSession) ResourceIDToHandlerAddress(arg0 [32]byte) (common.Address, error) {
	return _Bridge.Contract.ResourceIDToHandlerAddress(&_Bridge.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_Bridge *BridgeCaller) ResourceIDToTokenContractAddress(opts *bind.CallOpts, arg0 [32]byte) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_resourceIDToTokenContractAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_Bridge *BridgeSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _Bridge.Contract.ResourceIDToTokenContractAddress(&_Bridge.CallOpts, arg0)
}

// ResourceIDToTokenContractAddress is a free data retrieval call binding the contract method 0x0a6d55d8.
//
// Solidity: function _resourceIDToTokenContractAddress(bytes32 ) view returns(address)
func (_Bridge *BridgeCallerSession) ResourceIDToTokenContractAddress(arg0 [32]byte) (common.Address, error) {
	return _Bridge.Contract.ResourceIDToTokenContractAddress(&_Bridge.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_Bridge *BridgeCaller) TokenContractAddressToResourceID(opts *bind.CallOpts, arg0 common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_tokenContractAddressToResourceID", arg0)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_Bridge *BridgeSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _Bridge.Contract.TokenContractAddressToResourceID(&_Bridge.CallOpts, arg0)
}

// TokenContractAddressToResourceID is a free data retrieval call binding the contract method 0xc8ba6c87.
//
// Solidity: function _tokenContractAddressToResourceID(address ) view returns(bytes32)
func (_Bridge *BridgeCallerSession) TokenContractAddressToResourceID(arg0 common.Address) ([32]byte, error) {
	return _Bridge.Contract.TokenContractAddressToResourceID(&_Bridge.CallOpts, arg0)
}

// TotalProposals is a free data retrieval call binding the contract method 0x9d5773e0.
//
// Solidity: function _totalProposals() view returns(uint256)
func (_Bridge *BridgeCaller) TotalProposals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_totalProposals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalProposals is a free data retrieval call binding the contract method 0x9d5773e0.
//
// Solidity: function _totalProposals() view returns(uint256)
func (_Bridge *BridgeSession) TotalProposals() (*big.Int, error) {
	return _Bridge.Contract.TotalProposals(&_Bridge.CallOpts)
}

// TotalProposals is a free data retrieval call binding the contract method 0x9d5773e0.
//
// Solidity: function _totalProposals() view returns(uint256)
func (_Bridge *BridgeCallerSession) TotalProposals() (*big.Int, error) {
	return _Bridge.Contract.TotalProposals(&_Bridge.CallOpts)
}

// TotalRelayers is a free data retrieval call binding the contract method 0x802aabe8.
//
// Solidity: function _totalRelayers() view returns(uint256)
func (_Bridge *BridgeCaller) TotalRelayers(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "_totalRelayers")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalRelayers is a free data retrieval call binding the contract method 0x802aabe8.
//
// Solidity: function _totalRelayers() view returns(uint256)
func (_Bridge *BridgeSession) TotalRelayers() (*big.Int, error) {
	return _Bridge.Contract.TotalRelayers(&_Bridge.CallOpts)
}

// TotalRelayers is a free data retrieval call binding the contract method 0x802aabe8.
//
// Solidity: function _totalRelayers() view returns(uint256)
func (_Bridge *BridgeCallerSession) TotalRelayers() (*big.Int, error) {
	return _Bridge.Contract.TotalRelayers(&_Bridge.CallOpts)
}

// AddressToString is a free data retrieval call binding the contract method 0x5e57966d.
//
// Solidity: function addressToString(address account) pure returns(string)
func (_Bridge *BridgeCaller) AddressToString(opts *bind.CallOpts, account common.Address) (string, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "addressToString", account)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// AddressToString is a free data retrieval call binding the contract method 0x5e57966d.
//
// Solidity: function addressToString(address account) pure returns(string)
func (_Bridge *BridgeSession) AddressToString(account common.Address) (string, error) {
	return _Bridge.Contract.AddressToString(&_Bridge.CallOpts, account)
}

// AddressToString is a free data retrieval call binding the contract method 0x5e57966d.
//
// Solidity: function addressToString(address account) pure returns(string)
func (_Bridge *BridgeCallerSession) AddressToString(account common.Address) (string, error) {
	return _Bridge.Contract.AddressToString(&_Bridge.CallOpts, account)
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 value) pure returns(string)
func (_Bridge *BridgeCaller) Bytes32ToString(opts *bind.CallOpts, value [32]byte) (string, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "bytes32ToString", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 value) pure returns(string)
func (_Bridge *BridgeSession) Bytes32ToString(value [32]byte) (string, error) {
	return _Bridge.Contract.Bytes32ToString(&_Bridge.CallOpts, value)
}

// Bytes32ToString is a free data retrieval call binding the contract method 0x9201de55.
//
// Solidity: function bytes32ToString(bytes32 value) pure returns(string)
func (_Bridge *BridgeCallerSession) Bytes32ToString(value [32]byte) (string, error) {
	return _Bridge.Contract.Bytes32ToString(&_Bridge.CallOpts, value)
}

// BytesToString is a free data retrieval call binding the contract method 0x39614e4f.
//
// Solidity: function bytesToString(bytes value) pure returns(string)
func (_Bridge *BridgeCaller) BytesToString(opts *bind.CallOpts, value []byte) (string, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "bytesToString", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BytesToString is a free data retrieval call binding the contract method 0x39614e4f.
//
// Solidity: function bytesToString(bytes value) pure returns(string)
func (_Bridge *BridgeSession) BytesToString(value []byte) (string, error) {
	return _Bridge.Contract.BytesToString(&_Bridge.CallOpts, value)
}

// BytesToString is a free data retrieval call binding the contract method 0x39614e4f.
//
// Solidity: function bytesToString(bytes value) pure returns(string)
func (_Bridge *BridgeCallerSession) BytesToString(value []byte) (string, error) {
	return _Bridge.Contract.BytesToString(&_Bridge.CallOpts, value)
}

// GetProposal is a free data retrieval call binding the contract method 0xa9cf69fa.
//
// Solidity: function getProposal(uint8 originChainID, uint64 depositNonce, bytes32 dataHash) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_Bridge *BridgeCaller) GetProposal(opts *bind.CallOpts, originChainID uint8, depositNonce uint64, dataHash [32]byte) (BridgeProposal, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getProposal", originChainID, depositNonce, dataHash)

	if err != nil {
		return *new(BridgeProposal), err
	}

	out0 := *abi.ConvertType(out[0], new(BridgeProposal)).(*BridgeProposal)

	return out0, err

}

// GetProposal is a free data retrieval call binding the contract method 0xa9cf69fa.
//
// Solidity: function getProposal(uint8 originChainID, uint64 depositNonce, bytes32 dataHash) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_Bridge *BridgeSession) GetProposal(originChainID uint8, depositNonce uint64, dataHash [32]byte) (BridgeProposal, error) {
	return _Bridge.Contract.GetProposal(&_Bridge.CallOpts, originChainID, depositNonce, dataHash)
}

// GetProposal is a free data retrieval call binding the contract method 0xa9cf69fa.
//
// Solidity: function getProposal(uint8 originChainID, uint64 depositNonce, bytes32 dataHash) view returns((bytes32,bytes32,address[],address[],uint8,uint256))
func (_Bridge *BridgeCallerSession) GetProposal(originChainID uint8, depositNonce uint64, dataHash [32]byte) (BridgeProposal, error) {
	return _Bridge.Contract.GetProposal(&_Bridge.CallOpts, originChainID, depositNonce, dataHash)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bridge *BridgeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bridge *BridgeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bridge.Contract.GetRoleAdmin(&_Bridge.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bridge *BridgeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bridge.Contract.GetRoleAdmin(&_Bridge.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Bridge *BridgeCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Bridge *BridgeSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Bridge.Contract.GetRoleMember(&_Bridge.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Bridge *BridgeCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Bridge.Contract.GetRoleMember(&_Bridge.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Bridge *BridgeCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Bridge *BridgeSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Bridge.Contract.GetRoleMemberCount(&_Bridge.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Bridge *BridgeCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Bridge.Contract.GetRoleMemberCount(&_Bridge.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bridge *BridgeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bridge *BridgeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bridge.Contract.HasRole(&_Bridge.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bridge *BridgeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bridge.Contract.HasRole(&_Bridge.CallOpts, role, account)
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address relayer) view returns(bool)
func (_Bridge *BridgeCaller) IsRelayer(opts *bind.CallOpts, relayer common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "isRelayer", relayer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address relayer) view returns(bool)
func (_Bridge *BridgeSession) IsRelayer(relayer common.Address) (bool, error) {
	return _Bridge.Contract.IsRelayer(&_Bridge.CallOpts, relayer)
}

// IsRelayer is a free data retrieval call binding the contract method 0x541d5548.
//
// Solidity: function isRelayer(address relayer) view returns(bool)
func (_Bridge *BridgeCallerSession) IsRelayer(relayer common.Address) (bool, error) {
	return _Bridge.Contract.IsRelayer(&_Bridge.CallOpts, relayer)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeSession) Paused() (bool, error) {
	return _Bridge.Contract.Paused(&_Bridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeCallerSession) Paused() (bool, error) {
	return _Bridge.Contract.Paused(&_Bridge.CallOpts)
}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) pure returns(uint256)
func (_Bridge *BridgeCaller) Sub(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "sub", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) pure returns(uint256)
func (_Bridge *BridgeSession) Sub(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Bridge.Contract.Sub(&_Bridge.CallOpts, a, b)
}

// Sub is a free data retrieval call binding the contract method 0xb67d77c5.
//
// Solidity: function sub(uint256 a, uint256 b) pure returns(uint256)
func (_Bridge *BridgeCallerSession) Sub(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Bridge.Contract.Sub(&_Bridge.CallOpts, a, b)
}

// ToStringBase is a free data retrieval call binding the contract method 0x8d1b75e4.
//
// Solidity: function toStringBase(bytes data) pure returns(string)
func (_Bridge *BridgeCaller) ToStringBase(opts *bind.CallOpts, data []byte) (string, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "toStringBase", data)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ToStringBase is a free data retrieval call binding the contract method 0x8d1b75e4.
//
// Solidity: function toStringBase(bytes data) pure returns(string)
func (_Bridge *BridgeSession) ToStringBase(data []byte) (string, error) {
	return _Bridge.Contract.ToStringBase(&_Bridge.CallOpts, data)
}

// ToStringBase is a free data retrieval call binding the contract method 0x8d1b75e4.
//
// Solidity: function toStringBase(bytes data) pure returns(string)
func (_Bridge *BridgeCallerSession) ToStringBase(data []byte) (string, error) {
	return _Bridge.Contract.ToStringBase(&_Bridge.CallOpts, data)
}

// UintToString is a free data retrieval call binding the contract method 0xe9395679.
//
// Solidity: function uintToString(uint256 value) pure returns(string)
func (_Bridge *BridgeCaller) UintToString(opts *bind.CallOpts, value *big.Int) (string, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "uintToString", value)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UintToString is a free data retrieval call binding the contract method 0xe9395679.
//
// Solidity: function uintToString(uint256 value) pure returns(string)
func (_Bridge *BridgeSession) UintToString(value *big.Int) (string, error) {
	return _Bridge.Contract.UintToString(&_Bridge.CallOpts, value)
}

// UintToString is a free data retrieval call binding the contract method 0xe9395679.
//
// Solidity: function uintToString(uint256 value) pure returns(string)
func (_Bridge *BridgeCallerSession) UintToString(value *big.Int) (string, error) {
	return _Bridge.Contract.UintToString(&_Bridge.CallOpts, value)
}

// AdminAddRelayer is a paid mutator transaction binding the contract method 0xcdb0f73a.
//
// Solidity: function adminAddRelayer(address relayerAddress) returns()
func (_Bridge *BridgeTransactor) AdminAddRelayer(opts *bind.TransactOpts, relayerAddress common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminAddRelayer", relayerAddress)
}

// AdminAddRelayer is a paid mutator transaction binding the contract method 0xcdb0f73a.
//
// Solidity: function adminAddRelayer(address relayerAddress) returns()
func (_Bridge *BridgeSession) AdminAddRelayer(relayerAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminAddRelayer(&_Bridge.TransactOpts, relayerAddress)
}

// AdminAddRelayer is a paid mutator transaction binding the contract method 0xcdb0f73a.
//
// Solidity: function adminAddRelayer(address relayerAddress) returns()
func (_Bridge *BridgeTransactorSession) AdminAddRelayer(relayerAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminAddRelayer(&_Bridge.TransactOpts, relayerAddress)
}

// AdminChangeRelayerThreshold is a paid mutator transaction binding the contract method 0x4e056005.
//
// Solidity: function adminChangeRelayerThreshold(uint256 newThreshold) returns()
func (_Bridge *BridgeTransactor) AdminChangeRelayerThreshold(opts *bind.TransactOpts, newThreshold *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminChangeRelayerThreshold", newThreshold)
}

// AdminChangeRelayerThreshold is a paid mutator transaction binding the contract method 0x4e056005.
//
// Solidity: function adminChangeRelayerThreshold(uint256 newThreshold) returns()
func (_Bridge *BridgeSession) AdminChangeRelayerThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AdminChangeRelayerThreshold(&_Bridge.TransactOpts, newThreshold)
}

// AdminChangeRelayerThreshold is a paid mutator transaction binding the contract method 0x4e056005.
//
// Solidity: function adminChangeRelayerThreshold(uint256 newThreshold) returns()
func (_Bridge *BridgeTransactorSession) AdminChangeRelayerThreshold(newThreshold *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.AdminChangeRelayerThreshold(&_Bridge.TransactOpts, newThreshold)
}

// AdminPauseTransfers is a paid mutator transaction binding the contract method 0x80ae1c28.
//
// Solidity: function adminPauseTransfers() returns()
func (_Bridge *BridgeTransactor) AdminPauseTransfers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminPauseTransfers")
}

// AdminPauseTransfers is a paid mutator transaction binding the contract method 0x80ae1c28.
//
// Solidity: function adminPauseTransfers() returns()
func (_Bridge *BridgeSession) AdminPauseTransfers() (*types.Transaction, error) {
	return _Bridge.Contract.AdminPauseTransfers(&_Bridge.TransactOpts)
}

// AdminPauseTransfers is a paid mutator transaction binding the contract method 0x80ae1c28.
//
// Solidity: function adminPauseTransfers() returns()
func (_Bridge *BridgeTransactorSession) AdminPauseTransfers() (*types.Transaction, error) {
	return _Bridge.Contract.AdminPauseTransfers(&_Bridge.TransactOpts)
}

// AdminRemoveRelayer is a paid mutator transaction binding the contract method 0x9d82dd63.
//
// Solidity: function adminRemoveRelayer(address relayerAddress) returns()
func (_Bridge *BridgeTransactor) AdminRemoveRelayer(opts *bind.TransactOpts, relayerAddress common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminRemoveRelayer", relayerAddress)
}

// AdminRemoveRelayer is a paid mutator transaction binding the contract method 0x9d82dd63.
//
// Solidity: function adminRemoveRelayer(address relayerAddress) returns()
func (_Bridge *BridgeSession) AdminRemoveRelayer(relayerAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminRemoveRelayer(&_Bridge.TransactOpts, relayerAddress)
}

// AdminRemoveRelayer is a paid mutator transaction binding the contract method 0x9d82dd63.
//
// Solidity: function adminRemoveRelayer(address relayerAddress) returns()
func (_Bridge *BridgeTransactorSession) AdminRemoveRelayer(relayerAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminRemoveRelayer(&_Bridge.TransactOpts, relayerAddress)
}

// AdminSetBurnable is a paid mutator transaction binding the contract method 0x8c0c2631.
//
// Solidity: function adminSetBurnable(address handlerAddress, address tokenAddress) returns()
func (_Bridge *BridgeTransactor) AdminSetBurnable(opts *bind.TransactOpts, handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminSetBurnable", handlerAddress, tokenAddress)
}

// AdminSetBurnable is a paid mutator transaction binding the contract method 0x8c0c2631.
//
// Solidity: function adminSetBurnable(address handlerAddress, address tokenAddress) returns()
func (_Bridge *BridgeSession) AdminSetBurnable(handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetBurnable(&_Bridge.TransactOpts, handlerAddress, tokenAddress)
}

// AdminSetBurnable is a paid mutator transaction binding the contract method 0x8c0c2631.
//
// Solidity: function adminSetBurnable(address handlerAddress, address tokenAddress) returns()
func (_Bridge *BridgeTransactorSession) AdminSetBurnable(handlerAddress common.Address, tokenAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetBurnable(&_Bridge.TransactOpts, handlerAddress, tokenAddress)
}

// AdminSetGenericResource is a paid mutator transaction binding the contract method 0xe8437ee7.
//
// Solidity: function adminSetGenericResource(address handlerAddress, bytes32 resourceID, address contractAddress, bytes4 depositFunctionSig, bytes4 executeFunctionSig) returns()
func (_Bridge *BridgeTransactor) AdminSetGenericResource(opts *bind.TransactOpts, handlerAddress common.Address, resourceID [32]byte, contractAddress common.Address, depositFunctionSig [4]byte, executeFunctionSig [4]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminSetGenericResource", handlerAddress, resourceID, contractAddress, depositFunctionSig, executeFunctionSig)
}

// AdminSetGenericResource is a paid mutator transaction binding the contract method 0xe8437ee7.
//
// Solidity: function adminSetGenericResource(address handlerAddress, bytes32 resourceID, address contractAddress, bytes4 depositFunctionSig, bytes4 executeFunctionSig) returns()
func (_Bridge *BridgeSession) AdminSetGenericResource(handlerAddress common.Address, resourceID [32]byte, contractAddress common.Address, depositFunctionSig [4]byte, executeFunctionSig [4]byte) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetGenericResource(&_Bridge.TransactOpts, handlerAddress, resourceID, contractAddress, depositFunctionSig, executeFunctionSig)
}

// AdminSetGenericResource is a paid mutator transaction binding the contract method 0xe8437ee7.
//
// Solidity: function adminSetGenericResource(address handlerAddress, bytes32 resourceID, address contractAddress, bytes4 depositFunctionSig, bytes4 executeFunctionSig) returns()
func (_Bridge *BridgeTransactorSession) AdminSetGenericResource(handlerAddress common.Address, resourceID [32]byte, contractAddress common.Address, depositFunctionSig [4]byte, executeFunctionSig [4]byte) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetGenericResource(&_Bridge.TransactOpts, handlerAddress, resourceID, contractAddress, depositFunctionSig, executeFunctionSig)
}

// AdminSetResource is a paid mutator transaction binding the contract method 0xcb10f215.
//
// Solidity: function adminSetResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_Bridge *BridgeTransactor) AdminSetResource(opts *bind.TransactOpts, handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminSetResource", handlerAddress, resourceID, tokenAddress)
}

// AdminSetResource is a paid mutator transaction binding the contract method 0xcb10f215.
//
// Solidity: function adminSetResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_Bridge *BridgeSession) AdminSetResource(handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetResource(&_Bridge.TransactOpts, handlerAddress, resourceID, tokenAddress)
}

// AdminSetResource is a paid mutator transaction binding the contract method 0xcb10f215.
//
// Solidity: function adminSetResource(address handlerAddress, bytes32 resourceID, address tokenAddress) returns()
func (_Bridge *BridgeTransactorSession) AdminSetResource(handlerAddress common.Address, resourceID [32]byte, tokenAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.AdminSetResource(&_Bridge.TransactOpts, handlerAddress, resourceID, tokenAddress)
}

// AdminUnpauseTransfers is a paid mutator transaction binding the contract method 0xffaac0eb.
//
// Solidity: function adminUnpauseTransfers() returns()
func (_Bridge *BridgeTransactor) AdminUnpauseTransfers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "adminUnpauseTransfers")
}

// AdminUnpauseTransfers is a paid mutator transaction binding the contract method 0xffaac0eb.
//
// Solidity: function adminUnpauseTransfers() returns()
func (_Bridge *BridgeSession) AdminUnpauseTransfers() (*types.Transaction, error) {
	return _Bridge.Contract.AdminUnpauseTransfers(&_Bridge.TransactOpts)
}

// AdminUnpauseTransfers is a paid mutator transaction binding the contract method 0xffaac0eb.
//
// Solidity: function adminUnpauseTransfers() returns()
func (_Bridge *BridgeTransactorSession) AdminUnpauseTransfers() (*types.Transaction, error) {
	return _Bridge.Contract.AdminUnpauseTransfers(&_Bridge.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0x05e2ca17.
//
// Solidity: function deposit(uint8 destinationChainID, bytes32 resourceID, bytes data) payable returns()
func (_Bridge *BridgeTransactor) Deposit(opts *bind.TransactOpts, destinationChainID uint8, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "deposit", destinationChainID, resourceID, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x05e2ca17.
//
// Solidity: function deposit(uint8 destinationChainID, bytes32 resourceID, bytes data) payable returns()
func (_Bridge *BridgeSession) Deposit(destinationChainID uint8, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.Deposit(&_Bridge.TransactOpts, destinationChainID, resourceID, data)
}

// Deposit is a paid mutator transaction binding the contract method 0x05e2ca17.
//
// Solidity: function deposit(uint8 destinationChainID, bytes32 resourceID, bytes data) payable returns()
func (_Bridge *BridgeTransactorSession) Deposit(destinationChainID uint8, resourceID [32]byte, data []byte) (*types.Transaction, error) {
	return _Bridge.Contract.Deposit(&_Bridge.TransactOpts, destinationChainID, resourceID, data)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x4454b20d.
//
// Solidity: function executeProposal(uint8 chainID, uint64 depositNonce, bytes data, bytes32 resourceID) returns()
func (_Bridge *BridgeTransactor) ExecuteProposal(opts *bind.TransactOpts, chainID uint8, depositNonce uint64, data []byte, resourceID [32]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "executeProposal", chainID, depositNonce, data, resourceID)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x4454b20d.
//
// Solidity: function executeProposal(uint8 chainID, uint64 depositNonce, bytes data, bytes32 resourceID) returns()
func (_Bridge *BridgeSession) ExecuteProposal(chainID uint8, depositNonce uint64, data []byte, resourceID [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.ExecuteProposal(&_Bridge.TransactOpts, chainID, depositNonce, data, resourceID)
}

// ExecuteProposal is a paid mutator transaction binding the contract method 0x4454b20d.
//
// Solidity: function executeProposal(uint8 chainID, uint64 depositNonce, bytes data, bytes32 resourceID) returns()
func (_Bridge *BridgeTransactorSession) ExecuteProposal(chainID uint8, depositNonce uint64, data []byte, resourceID [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.ExecuteProposal(&_Bridge.TransactOpts, chainID, depositNonce, data, resourceID)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bridge *BridgeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.GrantRole(&_Bridge.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.GrantRole(&_Bridge.TransactOpts, role, account)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x5e1fab0f.
//
// Solidity: function renounceAdmin(address newAdmin) returns()
func (_Bridge *BridgeTransactor) RenounceAdmin(opts *bind.TransactOpts, newAdmin common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceAdmin", newAdmin)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x5e1fab0f.
//
// Solidity: function renounceAdmin(address newAdmin) returns()
func (_Bridge *BridgeSession) RenounceAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RenounceAdmin(&_Bridge.TransactOpts, newAdmin)
}

// RenounceAdmin is a paid mutator transaction binding the contract method 0x5e1fab0f.
//
// Solidity: function renounceAdmin(address newAdmin) returns()
func (_Bridge *BridgeTransactorSession) RenounceAdmin(newAdmin common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RenounceAdmin(&_Bridge.TransactOpts, newAdmin)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bridge *BridgeSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RenounceRole(&_Bridge.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RenounceRole(&_Bridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bridge *BridgeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RevokeRole(&_Bridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RevokeRole(&_Bridge.TransactOpts, role, account)
}

// SetAbiterList is a paid mutator transaction binding the contract method 0x7263cd76.
//
// Solidity: function setAbiterList(bytes[] _pubKeyList) returns()
func (_Bridge *BridgeTransactor) SetAbiterList(opts *bind.TransactOpts, _pubKeyList [][]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setAbiterList", _pubKeyList)
}

// SetAbiterList is a paid mutator transaction binding the contract method 0x7263cd76.
//
// Solidity: function setAbiterList(bytes[] _pubKeyList) returns()
func (_Bridge *BridgeSession) SetAbiterList(_pubKeyList [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.SetAbiterList(&_Bridge.TransactOpts, _pubKeyList)
}

// SetAbiterList is a paid mutator transaction binding the contract method 0x7263cd76.
//
// Solidity: function setAbiterList(bytes[] _pubKeyList) returns()
func (_Bridge *BridgeTransactorSession) SetAbiterList(_pubKeyList [][]byte) (*types.Transaction, error) {
	return _Bridge.Contract.SetAbiterList(&_Bridge.TransactOpts, _pubKeyList)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_Bridge *BridgeTransactor) SetBurnable(opts *bind.TransactOpts, contractAddress common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setBurnable", contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_Bridge *BridgeSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetBurnable(&_Bridge.TransactOpts, contractAddress)
}

// SetBurnable is a paid mutator transaction binding the contract method 0x07b7ed99.
//
// Solidity: function setBurnable(address contractAddress) returns()
func (_Bridge *BridgeTransactorSession) SetBurnable(contractAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetBurnable(&_Bridge.TransactOpts, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_Bridge *BridgeTransactor) SetResource(opts *bind.TransactOpts, resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setResource", resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_Bridge *BridgeSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetResource(&_Bridge.TransactOpts, resourceID, contractAddress)
}

// SetResource is a paid mutator transaction binding the contract method 0xb8fa3736.
//
// Solidity: function setResource(bytes32 resourceID, address contractAddress) returns()
func (_Bridge *BridgeTransactorSession) SetResource(resourceID [32]byte, contractAddress common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SetResource(&_Bridge.TransactOpts, resourceID, contractAddress)
}

// TransferFunds is a paid mutator transaction binding the contract method 0x4603ae38.
//
// Solidity: function transferFunds(address[] addrs, uint256[] amounts) returns()
func (_Bridge *BridgeTransactor) TransferFunds(opts *bind.TransactOpts, addrs []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "transferFunds", addrs, amounts)
}

// TransferFunds is a paid mutator transaction binding the contract method 0x4603ae38.
//
// Solidity: function transferFunds(address[] addrs, uint256[] amounts) returns()
func (_Bridge *BridgeSession) TransferFunds(addrs []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.TransferFunds(&_Bridge.TransactOpts, addrs, amounts)
}

// TransferFunds is a paid mutator transaction binding the contract method 0x4603ae38.
//
// Solidity: function transferFunds(address[] addrs, uint256[] amounts) returns()
func (_Bridge *BridgeTransactorSession) TransferFunds(addrs []common.Address, amounts []*big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.TransferFunds(&_Bridge.TransactOpts, addrs, amounts)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x1ff013f1.
//
// Solidity: function voteProposal(uint8 chainID, uint64 depositNonce, bytes32 resourceID, bytes32 dataHash) returns()
func (_Bridge *BridgeTransactor) VoteProposal(opts *bind.TransactOpts, chainID uint8, depositNonce uint64, resourceID [32]byte, dataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "voteProposal", chainID, depositNonce, resourceID, dataHash)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x1ff013f1.
//
// Solidity: function voteProposal(uint8 chainID, uint64 depositNonce, bytes32 resourceID, bytes32 dataHash) returns()
func (_Bridge *BridgeSession) VoteProposal(chainID uint8, depositNonce uint64, resourceID [32]byte, dataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.VoteProposal(&_Bridge.TransactOpts, chainID, depositNonce, resourceID, dataHash)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x1ff013f1.
//
// Solidity: function voteProposal(uint8 chainID, uint64 depositNonce, bytes32 resourceID, bytes32 dataHash) returns()
func (_Bridge *BridgeTransactorSession) VoteProposal(chainID uint8, depositNonce uint64, resourceID [32]byte, dataHash [32]byte) (*types.Transaction, error) {
	return _Bridge.Contract.VoteProposal(&_Bridge.TransactOpts, chainID, depositNonce, resourceID, dataHash)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_Bridge *BridgeTransactor) Withdraw(opts *bind.TransactOpts, tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "withdraw", tokenAddress, recipient, amountOrTokenID)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_Bridge *BridgeSession) Withdraw(tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Withdraw(&_Bridge.TransactOpts, tokenAddress, recipient, amountOrTokenID)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address tokenAddress, address recipient, uint256 amountOrTokenID) returns()
func (_Bridge *BridgeTransactorSession) Withdraw(tokenAddress common.Address, recipient common.Address, amountOrTokenID *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.Withdraw(&_Bridge.TransactOpts, tokenAddress, recipient, amountOrTokenID)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeSession) Receive() (*types.Transaction, error) {
	return _Bridge.Contract.Receive(&_Bridge.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Bridge *BridgeTransactorSession) Receive() (*types.Transaction, error) {
	return _Bridge.Contract.Receive(&_Bridge.TransactOpts)
}

// BridgeDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Bridge contract.
type BridgeDepositIterator struct {
	Event *BridgeDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDeposit represents a Deposit event raised by the Bridge contract.
type BridgeDeposit struct {
	DestinationChainID uint8
	ResourceID         [32]byte
	DepositNonce       uint64
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xdbb69440df8433824a026ef190652f29929eb64b4d1d5d2a69be8afe3e6eaed8.
//
// Solidity: event Deposit(uint8 indexed destinationChainID, bytes32 indexed resourceID, uint64 indexed depositNonce)
func (_Bridge *BridgeFilterer) FilterDeposit(opts *bind.FilterOpts, destinationChainID []uint8, resourceID [][32]byte, depositNonce []uint64) (*BridgeDepositIterator, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var resourceIDRule []interface{}
	for _, resourceIDItem := range resourceID {
		resourceIDRule = append(resourceIDRule, resourceIDItem)
	}
	var depositNonceRule []interface{}
	for _, depositNonceItem := range depositNonce {
		depositNonceRule = append(depositNonceRule, depositNonceItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Deposit", destinationChainIDRule, resourceIDRule, depositNonceRule)
	if err != nil {
		return nil, err
	}
	return &BridgeDepositIterator{contract: _Bridge.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xdbb69440df8433824a026ef190652f29929eb64b4d1d5d2a69be8afe3e6eaed8.
//
// Solidity: event Deposit(uint8 indexed destinationChainID, bytes32 indexed resourceID, uint64 indexed depositNonce)
func (_Bridge *BridgeFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *BridgeDeposit, destinationChainID []uint8, resourceID [][32]byte, depositNonce []uint64) (event.Subscription, error) {

	var destinationChainIDRule []interface{}
	for _, destinationChainIDItem := range destinationChainID {
		destinationChainIDRule = append(destinationChainIDRule, destinationChainIDItem)
	}
	var resourceIDRule []interface{}
	for _, resourceIDItem := range resourceID {
		resourceIDRule = append(resourceIDRule, resourceIDItem)
	}
	var depositNonceRule []interface{}
	for _, depositNonceItem := range depositNonce {
		depositNonceRule = append(depositNonceRule, depositNonceItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Deposit", destinationChainIDRule, resourceIDRule, depositNonceRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDeposit)
				if err := _Bridge.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0xdbb69440df8433824a026ef190652f29929eb64b4d1d5d2a69be8afe3e6eaed8.
//
// Solidity: event Deposit(uint8 indexed destinationChainID, bytes32 indexed resourceID, uint64 indexed depositNonce)
func (_Bridge *BridgeFilterer) ParseDeposit(log types.Log) (*BridgeDeposit, error) {
	event := new(BridgeDeposit)
	if err := _Bridge.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeLogStringIterator is returned from FilterLogString and is used to iterate over the raw logs and unpacked data for LogString events raised by the Bridge contract.
type BridgeLogStringIterator struct {
	Event *BridgeLogString // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeLogStringIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeLogString)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeLogString)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeLogStringIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeLogStringIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeLogString represents a LogString event raised by the Bridge contract.
type BridgeLogString struct {
	Data string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterLogString is a free log retrieval operation binding the contract event 0xa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f.
//
// Solidity: event LogString(string data)
func (_Bridge *BridgeFilterer) FilterLogString(opts *bind.FilterOpts) (*BridgeLogStringIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "LogString")
	if err != nil {
		return nil, err
	}
	return &BridgeLogStringIterator{contract: _Bridge.contract, event: "LogString", logs: logs, sub: sub}, nil
}

// WatchLogString is a free log subscription operation binding the contract event 0xa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f.
//
// Solidity: event LogString(string data)
func (_Bridge *BridgeFilterer) WatchLogString(opts *bind.WatchOpts, sink chan<- *BridgeLogString) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "LogString")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeLogString)
				if err := _Bridge.contract.UnpackLog(event, "LogString", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLogString is a log parse operation binding the contract event 0xa95e6e2a182411e7a6f9ed114a85c3761d87f9b8f453d842c71235aa64fff99f.
//
// Solidity: event LogString(string data)
func (_Bridge *BridgeFilterer) ParseLogString(log types.Log) (*BridgeLogString, error) {
	event := new(BridgeLogString)
	if err := _Bridge.contract.UnpackLog(event, "LogString", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Bridge contract.
type BridgePausedIterator struct {
	Event *BridgePaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgePaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgePaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgePaused represents a Paused event raised by the Bridge contract.
type BridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bridge *BridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*BridgePausedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BridgePausedIterator{contract: _Bridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bridge *BridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BridgePaused) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgePaused)
				if err := _Bridge.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bridge *BridgeFilterer) ParsePaused(log types.Log) (*BridgePaused, error) {
	event := new(BridgePaused)
	if err := _Bridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeProposalEventIterator is returned from FilterProposalEvent and is used to iterate over the raw logs and unpacked data for ProposalEvent events raised by the Bridge contract.
type BridgeProposalEventIterator struct {
	Event *BridgeProposalEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeProposalEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeProposalEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeProposalEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeProposalEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeProposalEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeProposalEvent represents a ProposalEvent event raised by the Bridge contract.
type BridgeProposalEvent struct {
	OriginChainID uint8
	DepositNonce  uint64
	Status        uint8
	ResourceID    [32]byte
	DataHash      [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterProposalEvent is a free log retrieval operation binding the contract event 0x803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab650417.
//
// Solidity: event ProposalEvent(uint8 indexed originChainID, uint64 indexed depositNonce, uint8 indexed status, bytes32 resourceID, bytes32 dataHash)
func (_Bridge *BridgeFilterer) FilterProposalEvent(opts *bind.FilterOpts, originChainID []uint8, depositNonce []uint64, status []uint8) (*BridgeProposalEventIterator, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var depositNonceRule []interface{}
	for _, depositNonceItem := range depositNonce {
		depositNonceRule = append(depositNonceRule, depositNonceItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ProposalEvent", originChainIDRule, depositNonceRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &BridgeProposalEventIterator{contract: _Bridge.contract, event: "ProposalEvent", logs: logs, sub: sub}, nil
}

// WatchProposalEvent is a free log subscription operation binding the contract event 0x803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab650417.
//
// Solidity: event ProposalEvent(uint8 indexed originChainID, uint64 indexed depositNonce, uint8 indexed status, bytes32 resourceID, bytes32 dataHash)
func (_Bridge *BridgeFilterer) WatchProposalEvent(opts *bind.WatchOpts, sink chan<- *BridgeProposalEvent, originChainID []uint8, depositNonce []uint64, status []uint8) (event.Subscription, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var depositNonceRule []interface{}
	for _, depositNonceItem := range depositNonce {
		depositNonceRule = append(depositNonceRule, depositNonceItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ProposalEvent", originChainIDRule, depositNonceRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeProposalEvent)
				if err := _Bridge.contract.UnpackLog(event, "ProposalEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalEvent is a log parse operation binding the contract event 0x803c5a12f6bde629cea32e63d4b92d1b560816a6fb72e939d3c89e1cab650417.
//
// Solidity: event ProposalEvent(uint8 indexed originChainID, uint64 indexed depositNonce, uint8 indexed status, bytes32 resourceID, bytes32 dataHash)
func (_Bridge *BridgeFilterer) ParseProposalEvent(log types.Log) (*BridgeProposalEvent, error) {
	event := new(BridgeProposalEvent)
	if err := _Bridge.contract.UnpackLog(event, "ProposalEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeProposalVoteIterator is returned from FilterProposalVote and is used to iterate over the raw logs and unpacked data for ProposalVote events raised by the Bridge contract.
type BridgeProposalVoteIterator struct {
	Event *BridgeProposalVote // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeProposalVoteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeProposalVote)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeProposalVote)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeProposalVoteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeProposalVoteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeProposalVote represents a ProposalVote event raised by the Bridge contract.
type BridgeProposalVote struct {
	OriginChainID uint8
	DepositNonce  uint64
	Status        uint8
	ResourceID    [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterProposalVote is a free log retrieval operation binding the contract event 0x25f8daaa4635a7729927ba3f5b3d59cc3320aca7c32c9db4e7ca7b9574343640.
//
// Solidity: event ProposalVote(uint8 indexed originChainID, uint64 indexed depositNonce, uint8 indexed status, bytes32 resourceID)
func (_Bridge *BridgeFilterer) FilterProposalVote(opts *bind.FilterOpts, originChainID []uint8, depositNonce []uint64, status []uint8) (*BridgeProposalVoteIterator, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var depositNonceRule []interface{}
	for _, depositNonceItem := range depositNonce {
		depositNonceRule = append(depositNonceRule, depositNonceItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "ProposalVote", originChainIDRule, depositNonceRule, statusRule)
	if err != nil {
		return nil, err
	}
	return &BridgeProposalVoteIterator{contract: _Bridge.contract, event: "ProposalVote", logs: logs, sub: sub}, nil
}

// WatchProposalVote is a free log subscription operation binding the contract event 0x25f8daaa4635a7729927ba3f5b3d59cc3320aca7c32c9db4e7ca7b9574343640.
//
// Solidity: event ProposalVote(uint8 indexed originChainID, uint64 indexed depositNonce, uint8 indexed status, bytes32 resourceID)
func (_Bridge *BridgeFilterer) WatchProposalVote(opts *bind.WatchOpts, sink chan<- *BridgeProposalVote, originChainID []uint8, depositNonce []uint64, status []uint8) (event.Subscription, error) {

	var originChainIDRule []interface{}
	for _, originChainIDItem := range originChainID {
		originChainIDRule = append(originChainIDRule, originChainIDItem)
	}
	var depositNonceRule []interface{}
	for _, depositNonceItem := range depositNonce {
		depositNonceRule = append(depositNonceRule, depositNonceItem)
	}
	var statusRule []interface{}
	for _, statusItem := range status {
		statusRule = append(statusRule, statusItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "ProposalVote", originChainIDRule, depositNonceRule, statusRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeProposalVote)
				if err := _Bridge.contract.UnpackLog(event, "ProposalVote", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalVote is a log parse operation binding the contract event 0x25f8daaa4635a7729927ba3f5b3d59cc3320aca7c32c9db4e7ca7b9574343640.
//
// Solidity: event ProposalVote(uint8 indexed originChainID, uint64 indexed depositNonce, uint8 indexed status, bytes32 resourceID)
func (_Bridge *BridgeFilterer) ParseProposalVote(log types.Log) (*BridgeProposalVote, error) {
	event := new(BridgeProposalVote)
	if err := _Bridge.contract.UnpackLog(event, "ProposalVote", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRelayerAddedIterator is returned from FilterRelayerAdded and is used to iterate over the raw logs and unpacked data for RelayerAdded events raised by the Bridge contract.
type BridgeRelayerAddedIterator struct {
	Event *BridgeRelayerAdded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeRelayerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRelayerAdded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeRelayerAdded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeRelayerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRelayerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRelayerAdded represents a RelayerAdded event raised by the Bridge contract.
type BridgeRelayerAdded struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerAdded is a free log retrieval operation binding the contract event 0x03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c5.
//
// Solidity: event RelayerAdded(address indexed relayer)
func (_Bridge *BridgeFilterer) FilterRelayerAdded(opts *bind.FilterOpts, relayer []common.Address) (*BridgeRelayerAddedIterator, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RelayerAdded", relayerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRelayerAddedIterator{contract: _Bridge.contract, event: "RelayerAdded", logs: logs, sub: sub}, nil
}

// WatchRelayerAdded is a free log subscription operation binding the contract event 0x03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c5.
//
// Solidity: event RelayerAdded(address indexed relayer)
func (_Bridge *BridgeFilterer) WatchRelayerAdded(opts *bind.WatchOpts, sink chan<- *BridgeRelayerAdded, relayer []common.Address) (event.Subscription, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RelayerAdded", relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRelayerAdded)
				if err := _Bridge.contract.UnpackLog(event, "RelayerAdded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRelayerAdded is a log parse operation binding the contract event 0x03580ee9f53a62b7cb409a2cb56f9be87747dd15017afc5cef6eef321e4fb2c5.
//
// Solidity: event RelayerAdded(address indexed relayer)
func (_Bridge *BridgeFilterer) ParseRelayerAdded(log types.Log) (*BridgeRelayerAdded, error) {
	event := new(BridgeRelayerAdded)
	if err := _Bridge.contract.UnpackLog(event, "RelayerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRelayerRemovedIterator is returned from FilterRelayerRemoved and is used to iterate over the raw logs and unpacked data for RelayerRemoved events raised by the Bridge contract.
type BridgeRelayerRemovedIterator struct {
	Event *BridgeRelayerRemoved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeRelayerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRelayerRemoved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeRelayerRemoved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeRelayerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRelayerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRelayerRemoved represents a RelayerRemoved event raised by the Bridge contract.
type BridgeRelayerRemoved struct {
	Relayer common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRelayerRemoved is a free log retrieval operation binding the contract event 0x10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b.
//
// Solidity: event RelayerRemoved(address indexed relayer)
func (_Bridge *BridgeFilterer) FilterRelayerRemoved(opts *bind.FilterOpts, relayer []common.Address) (*BridgeRelayerRemovedIterator, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RelayerRemoved", relayerRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRelayerRemovedIterator{contract: _Bridge.contract, event: "RelayerRemoved", logs: logs, sub: sub}, nil
}

// WatchRelayerRemoved is a free log subscription operation binding the contract event 0x10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b.
//
// Solidity: event RelayerRemoved(address indexed relayer)
func (_Bridge *BridgeFilterer) WatchRelayerRemoved(opts *bind.WatchOpts, sink chan<- *BridgeRelayerRemoved, relayer []common.Address) (event.Subscription, error) {

	var relayerRule []interface{}
	for _, relayerItem := range relayer {
		relayerRule = append(relayerRule, relayerItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RelayerRemoved", relayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRelayerRemoved)
				if err := _Bridge.contract.UnpackLog(event, "RelayerRemoved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRelayerRemoved is a log parse operation binding the contract event 0x10e1f7ce9fd7d1b90a66d13a2ab3cb8dd7f29f3f8d520b143b063ccfbab6906b.
//
// Solidity: event RelayerRemoved(address indexed relayer)
func (_Bridge *BridgeFilterer) ParseRelayerRemoved(log types.Log) (*BridgeRelayerRemoved, error) {
	event := new(BridgeRelayerRemoved)
	if err := _Bridge.contract.UnpackLog(event, "RelayerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRelayerThresholdChangedIterator is returned from FilterRelayerThresholdChanged and is used to iterate over the raw logs and unpacked data for RelayerThresholdChanged events raised by the Bridge contract.
type BridgeRelayerThresholdChangedIterator struct {
	Event *BridgeRelayerThresholdChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeRelayerThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRelayerThresholdChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeRelayerThresholdChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeRelayerThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRelayerThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRelayerThresholdChanged represents a RelayerThresholdChanged event raised by the Bridge contract.
type BridgeRelayerThresholdChanged struct {
	NewThreshold *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRelayerThresholdChanged is a free log retrieval operation binding the contract event 0xa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c8.
//
// Solidity: event RelayerThresholdChanged(uint256 indexed newThreshold)
func (_Bridge *BridgeFilterer) FilterRelayerThresholdChanged(opts *bind.FilterOpts, newThreshold []*big.Int) (*BridgeRelayerThresholdChangedIterator, error) {

	var newThresholdRule []interface{}
	for _, newThresholdItem := range newThreshold {
		newThresholdRule = append(newThresholdRule, newThresholdItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RelayerThresholdChanged", newThresholdRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRelayerThresholdChangedIterator{contract: _Bridge.contract, event: "RelayerThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchRelayerThresholdChanged is a free log subscription operation binding the contract event 0xa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c8.
//
// Solidity: event RelayerThresholdChanged(uint256 indexed newThreshold)
func (_Bridge *BridgeFilterer) WatchRelayerThresholdChanged(opts *bind.WatchOpts, sink chan<- *BridgeRelayerThresholdChanged, newThreshold []*big.Int) (event.Subscription, error) {

	var newThresholdRule []interface{}
	for _, newThresholdItem := range newThreshold {
		newThresholdRule = append(newThresholdRule, newThresholdItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RelayerThresholdChanged", newThresholdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRelayerThresholdChanged)
				if err := _Bridge.contract.UnpackLog(event, "RelayerThresholdChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRelayerThresholdChanged is a log parse operation binding the contract event 0xa20d6b84cd798a24038be305eff8a45ca82ef54a2aa2082005d8e14c0a4746c8.
//
// Solidity: event RelayerThresholdChanged(uint256 indexed newThreshold)
func (_Bridge *BridgeFilterer) ParseRelayerThresholdChanged(log types.Log) (*BridgeRelayerThresholdChanged, error) {
	event := new(BridgeRelayerThresholdChanged)
	if err := _Bridge.contract.UnpackLog(event, "RelayerThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Bridge contract.
type BridgeRoleAdminChangedIterator struct {
	Event *BridgeRoleAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRoleAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeRoleAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRoleAdminChanged represents a RoleAdminChanged event raised by the Bridge contract.
type BridgeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bridge *BridgeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BridgeRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRoleAdminChangedIterator{contract: _Bridge.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bridge *BridgeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BridgeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRoleAdminChanged)
				if err := _Bridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bridge *BridgeFilterer) ParseRoleAdminChanged(log types.Log) (*BridgeRoleAdminChanged, error) {
	event := new(BridgeRoleAdminChanged)
	if err := _Bridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Bridge contract.
type BridgeRoleGrantedIterator struct {
	Event *BridgeRoleGranted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRoleGranted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeRoleGranted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRoleGranted represents a RoleGranted event raised by the Bridge contract.
type BridgeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRoleGrantedIterator{contract: _Bridge.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BridgeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRoleGranted)
				if err := _Bridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) ParseRoleGranted(log types.Log) (*BridgeRoleGranted, error) {
	event := new(BridgeRoleGranted)
	if err := _Bridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Bridge contract.
type BridgeRoleRevokedIterator struct {
	Event *BridgeRoleRevoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRoleRevoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeRoleRevoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRoleRevoked represents a RoleRevoked event raised by the Bridge contract.
type BridgeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRoleRevokedIterator{contract: _Bridge.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BridgeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRoleRevoked)
				if err := _Bridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) ParseRoleRevoked(log types.Log) (*BridgeRoleRevoked, error) {
	event := new(BridgeRoleRevoked)
	if err := _Bridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Bridge contract.
type BridgeUnpausedIterator struct {
	Event *BridgeUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BridgeUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeUnpaused represents a Unpaused event raised by the Bridge contract.
type BridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bridge *BridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BridgeUnpausedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BridgeUnpausedIterator{contract: _Bridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bridge *BridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeUnpaused)
				if err := _Bridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bridge *BridgeFilterer) ParseUnpaused(log types.Log) (*BridgeUnpaused, error) {
	event := new(BridgeUnpaused)
	if err := _Bridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
