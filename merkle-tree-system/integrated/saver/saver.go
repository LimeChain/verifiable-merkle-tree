package saver

import (
	MerkleLimeContract "../contracts"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/LimeChain/merkletree"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"sync"
)

type EthereumRootSaver struct {
	client          *ethclient.Client
	privateKey      *ecdsa.PrivateKey
	contractAddress string
	tree            merkletree.MerkleTree
	mutex           sync.RWMutex
}

func (saver *EthereumRootSaver) GetContractAddress() string {
	return saver.contractAddress
}
func (saver *EthereumRootSaver) GetSaverWalletAddress() (string, error) {
	publicKey := saver.privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return "", errors.New("Could not derive public key from the saved private key")
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return address, nil
}

func (saver *EthereumRootSaver) getDefaultOptions(value, gasLimit uint) (*bind.TransactOpts, error) {
	fromAddress, err := saver.GetSaverWalletAddress()
	if err != nil {
		return nil, err
	}
	nonce, err := saver.client.PendingNonceAt(context.Background(), common.HexToAddress(fromAddress))
	if err != nil {
		return nil, err
	}

	gasPrice, err := saver.client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(saver.privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(int64(value)) // in wei
	auth.GasLimit = uint64(gasLimit)      // in units
	auth.GasPrice = gasPrice
	return auth, nil
}

func (saver *EthereumRootSaver) TriggerSave() (string, error) {
	saver.mutex.Lock()
	defer saver.mutex.Unlock()

	address := common.HexToAddress(saver.contractAddress)
	contract, err := MerkleLimeContract.NewMerklelimecontract(address, saver.client)
	if err != nil {
		return "", err
	}

	opts, err := saver.getDefaultOptions(0, 600000)
	if err != nil {
		return "", err
	}

	newRoot := [32]byte{}
	copy(newRoot[:], []byte(saver.tree.Root()))

	tx, err := contract.SetRoot(opts, newRoot)
	if err != nil {
		return "", err
	}

	return tx.Hash().Hex(), nil
}

func (saver *EthereumRootSaver) FetchRoot() (string, error) {
	address := common.HexToAddress(saver.contractAddress)
	contract, err := MerkleLimeContract.NewMerklelimecontract(address, saver.client)
	if err != nil {
		return "", err
	}

	result, err := contract.LimeRoot(nil)
	if err != nil {
		return "", err
	}

	hexResult := common.Bytes2Hex(result[:])

	fmt.Println(hexResult)

	return hexResult, nil
}

func NewSaver(host, privateKeyHex, contractAddressHex string, tree merkletree.MerkleTree) (saver *EthereumRootSaver, err error) {
	client, err := ethclient.Dial(host)
	if err != nil {
		return nil, err
	}

	privateKey, err := crypto.HexToECDSA(privateKeyHex)

	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	_, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("The private key could not produce a public key")
	}

	saver = &EthereumRootSaver{}
	saver.client = client
	saver.privateKey = privateKey
	saver.contractAddress = contractAddressHex
	saver.tree = tree

	return saver, nil
}
