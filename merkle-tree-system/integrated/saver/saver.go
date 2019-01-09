package saver

import (
	RootValidator "../contracts"
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"github.com/LimeChain/merkletree"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"sync"
	"time"
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

func (saver *EthereumRootSaver) TriggerSave() (string, error) {
	saver.mutex.Lock()
	defer saver.mutex.Unlock()

	address := common.HexToAddress(saver.contractAddress)
	contract, err := RootValidator.NewRootValidator(address, saver.client)
	if err != nil {
		return "", err
	}

	tx, err := contract.SetRoot(bind.NewKeyedTransactor(saver.privateKey), common.HexToHash(saver.tree.Root()))
	fmt.Printf("Broadcasted tx hash: %v\n", tx.Hash().Hex())
	if err != nil {
		return "", err
	}

	retries := 150
	for {
		if retries == 0 {
			return "", errors.New("Could not wait for transaction to be mined")
		}
		t, isPending, err := saver.client.TransactionByHash(context.Background(), tx.Hash())
		if err != nil {
			panic(err)
			// return "", err
		}
		if isPending || t == nil {
			time.Sleep(2 * time.Second)
			retries--
			continue
		}
		return t.Hash().Hex(), nil
	}

	return "", errors.New("Could not save the root hash")
}

func (saver *EthereumRootSaver) FetchRoot() (string, error) {
	address := common.HexToAddress(saver.contractAddress)
	contract, err := RootValidator.NewRootValidator(address, saver.client)
	if err != nil {
		return "", err
	}

	result, err := contract.LimeRoot(nil)
	if err != nil {
		return "", err
	}

	empty := [32]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	if bytes.Equal(result[:], empty[:]) {
		return "", nil
	}

	hexResult := common.Bytes2Hex(result[:])

	return "0x" + hexResult, nil
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
