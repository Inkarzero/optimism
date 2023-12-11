// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const CrossL2InboxStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/interop/CrossL2Inbox.sol:CrossL2Inbox\",\"label\":\"roots\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_bytes32,t_mapping(t_bytes32,t_bool))\"},{\"astId\":1001,\"contract\":\"src/interop/CrossL2Inbox.sol:CrossL2Inbox\",\"label\":\"crossL2Sender\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_address\"},{\"astId\":1002,\"contract\":\"src/interop/CrossL2Inbox.sol:CrossL2Inbox\",\"label\":\"messageSourceChain\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_bytes32\"},{\"astId\":1003,\"contract\":\"src/interop/CrossL2Inbox.sol:CrossL2Inbox\",\"label\":\"consumedMessages\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_mapping(t_bytes32,t_bool)\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_mapping(t_bytes32,t_mapping(t_bytes32,t_bool))\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e mapping(bytes32 =\u003e bool))\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_mapping(t_bytes32,t_bool)\"}}}"

var CrossL2InboxStorageLayout = new(solc.StorageLayout)

var CrossL2InboxDeployedBin = "0x60806040526004361061007f5760003560e01c806360a687aa1161004e57806360a687aa1461019a578063c00157da146101be578063db10b9a9146101f1578063f2c5dcb81461022957600080fd5b806339bc3c811461008b5780633d6d0dd4146100d057806354fd4d50146101225780635d9eb6d01461017857600080fd5b3661008657005b600080fd5b34801561009757600080fd5b506100bb6100a6366004610a0d565b60036020526000908152604090205460ff1681565b60405190151581526020015b60405180910390f35b3480156100dc57600080fd5b506001546100fd9073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100c7565b34801561012e57600080fd5b5061016b6040518060400160405280600581526020017f302e302e3100000000000000000000000000000000000000000000000000000081525081565b6040516100c79190610a26565b34801561018457600080fd5b50610198610193366004610c0f565b610249565b005b3480156101a657600080fd5b506101b060025481565b6040519081526020016100c7565b3480156101ca57600080fd5b507f00000000000000000000000000000000000000000000000000000000000000006100fd565b3480156101fd57600080fd5b506100bb61020c366004610d02565b600060208181529281526040808220909352908152205460ff1681565b34801561023557600080fd5b50610198610244366004610d24565b6106a3565b60208085015160009081528082526040808220868352909252205460ff1661031e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604c60248201527f43726f73734c32496e626f783a206d7573742070726f6f6620616761696e737460448201527f206b6e6f776e206f757470757420726f6f742066726f6d206d6573736167652060648201527f736f7572636520636861696e0000000000000000000000000000000000000000608482015260a4015b60405180910390fd5b6000610329856107fd565b905080600052600060205260406000208060005250602060002060405180600182536001828101889052602060218401526041830184905260618301819052608183015260a190910190848683379084018181039250906103e88284836021620186a0fa92508261039a576103e882fd5b600052505060015473ffffffffffffffffffffffffffffffffffffffff1661dead14610448576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603c60248201527f43726f73734c32496e626f783a2063616e206f6e6c792074726967676572206f60448201527f6e652063616c6c207065722063726f7373204c32206d657373616765000000006064820152608401610315565b60008181526003602052604090205460ff16156104e7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602f60248201527f43726f73734c32496e626f783a206d6573736167652068617320616c7265616460448201527f79206265656e20636f6e73756d656400000000000000000000000000000000006064820152608401610315565b60008181526003602090815260408220805460017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff009091168117909155606088015181547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff909116179055860151600255608086015160c087015160a088015160e089015161058e93929190610991565b600180547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055600060025560405190915082907f608b51d991a28926c87c94dae8c72df6a62c5f22b359bb418bf204355b39fa7d906105f890841515815260200190565b60405180910390a28015801561060e5750326001145b1561069b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603460248201527f43726f73734c32496e626f783a2063726f7373204c32206d657373616765206360448201527f616c6c20657865637574696f6e206661696c65640000000000000000000000006064820152608401610315565b505050505050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610768576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f43726f73734c32496e626f783a206f6e6c7920706f737469652063616e20646560448201527f6c69766572206d61696c000000000000000000000000000000000000000000006064820152608401610315565b60005b818110156107f857600160008085858581811061078a5761078a610d99565b90506040020160000135815260200190815260200160002060008585858181106107b6576107b6610d99565b90506040020160200135815260200190815260200160002060006101000a81548160ff02191690831515021790555080806107f090610dc8565b91505061076b565b505050565b60e0810151805160209182018190206040805193840182905283019190915260009182906060016040516020818303038152906040528051906020012090506000846060015185608001518660a001518760c00151604051602001610896949392919073ffffffffffffffffffffffffffffffffffffffff94851681529290931660208301526040820152606081019190915260800190565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0818403018152828252805160209182012081840152828201949094528051808303820181526060830182528051908501208785015197820151608084019890985260a0808401989098528151808403909801885260c08301825287519785019790972060e0830152610100808301979097528051808303909701875261012090910190525083519301929092207effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01000000000000000000000000000000000000000000000000000000000000001792915050565b60008060006109a18660006109ef565b9050806109d7576308c379a06000526020805278185361666543616c6c3a204e6f7420656e6f756768206761736058526064601cfd5b600080855160208701888b5af1979650505050505050565b600080603f83619c4001026040850201603f5a021015949350505050565b600060208284031215610a1f57600080fd5b5035919050565b600060208083528351808285015260005b81811015610a5357858101830151858201604001528201610a37565b81811115610a65576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051610100810167ffffffffffffffff81118282101715610aec57610aec610a99565b60405290565b803573ffffffffffffffffffffffffffffffffffffffff81168114610b1657600080fd5b919050565b600082601f830112610b2c57600080fd5b813567ffffffffffffffff80821115610b4757610b47610a99565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715610b8d57610b8d610a99565b81604052838152866020858801011115610ba657600080fd5b836020870160208301376000602085830101528094505050505092915050565b60008083601f840112610bd857600080fd5b50813567ffffffffffffffff811115610bf057600080fd5b602083019150836020828501011115610c0857600080fd5b9250929050565b60008060008060608587031215610c2557600080fd5b843567ffffffffffffffff80821115610c3d57600080fd5b908601906101008289031215610c5257600080fd5b610c5a610ac8565b823581526020830135602082015260408301356040820152610c7e60608401610af2565b6060820152610c8f60808401610af2565b608082015260a083013560a082015260c083013560c082015260e083013582811115610cba57600080fd5b610cc68a828601610b1b565b60e0830152509550602087013594506040870135915080821115610ce957600080fd5b50610cf687828801610bc6565b95989497509550505050565b60008060408385031215610d1557600080fd5b50508035926020909101359150565b60008060208385031215610d3757600080fd5b823567ffffffffffffffff80821115610d4f57600080fd5b818501915085601f830112610d6357600080fd5b813581811115610d7257600080fd5b8660208260061b8501011115610d8757600080fd5b60209290920196919550909350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610e20577f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b506001019056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(CrossL2InboxStorageLayoutJSON), CrossL2InboxStorageLayout); err != nil {
		panic(err)
	}

	layouts["CrossL2Inbox"] = CrossL2InboxStorageLayout
	deployedBytecodes["CrossL2Inbox"] = CrossL2InboxDeployedBin
	immutableReferences["CrossL2Inbox"] = true
}
