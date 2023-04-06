package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/matopenKW/go-contract/internal/pkg/contract"
)

var (
	privateKey      string
	contractAddress string
	blockChainHost  string
)

const (
	chainID = 80001
)

func init() {
	privateKey = os.Getenv("PRIVATE_KEY")
	blockChainHost = os.Getenv("BLOCK_CHAIN_HOST")
	contractAddress = os.Getenv("CONTRACT_ADDRESS")
}

func main() {
	ctx := context.Background()

	ethCli, err := newEthCli(blockChainHost)
	if err != nil {
		panic(err)
	}

	addr := common.HexToAddress(contractAddress)
	contract, err := contract.NewContract(addr, ethCli)
	if err != nil {
		panic(err)
	}

	opt, err := newTransactOpts(ctx, ethCli)
	if err != nil {
		panic(err)
	}
	_ = opt

	fmt.Println(contract.GetMessage(nil))

	//if _, err := contract.SetMessage(opt, "Hello Smart Contract!"); err != nil {
	//	panic(err)
	//}

	//fmt.Println(contract.GetMessage(nil))
}

func newTransactOpts(ctx context.Context, cli *ethclient.Client) (*bind.TransactOpts, error) {
	priKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		return nil, fmt.Errorf("HexToECDSA error, %v", err)
	}

	publicKey := priKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := cli.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, fmt.Errorf("PendingNonceAt error, %v", err)
	}

	gasPrice, err := cli.SuggestGasPrice(ctx)
	if err != nil {
		return nil, fmt.Errorf("SuggestGasPrice error, %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(priKey, big.NewInt(chainID))
	if err != nil {
		return nil, fmt.Errorf("NewKeyedTransactorWithChainID error, %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	return auth, nil
}

func newEthCli(networkHost string) (*ethclient.Client, error) {
	cli, err := ethclient.Dial(networkHost)
	if err != nil {
		return nil, fmt.Errorf("failed to connect chain")
	}
	return cli, nil
}
