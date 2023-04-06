package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/matopenKW/go-contract/internal/pkg/history_contract"
	"github.com/spf13/cobra"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	privateKey      string
	contractAddress string
	blockChainHost  string
)

const (
	chainID      = 80001
	ownerAddress = "0x111b5dbd7f34ecf02a8857261a9dbb459a4b7fb2"
)

func init() {
	privateKey = os.Getenv("PRIVATE_KEY")
	blockChainHost = os.Getenv("BLOCK_CHAIN_HOST")
	contractAddress = os.Getenv("CONTRACT_ADDRESS")
}

func main() {

	var rootCmd = &cobra.Command{
		Use: "go-contract",
	}

	var cmd1 = &cobra.Command{
		Use: "store",
		Run: func(cmd *cobra.Command, args []string) {
			if err := store(); err != nil {
				panic(err)
			}
		},
	}

	var cmd2 = &cobra.Command{
		Use: "get-history",
		Run: func(cmd *cobra.Command, args []string) {
			if err := getHistory(); err != nil {
				panic(err)
			}
		},
	}

	rootCmd.AddCommand(cmd1, cmd2)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func store() error {
	ethCli, err := newEthCli(blockChainHost)
	if err != nil {
		return err
	}

	addr := common.HexToAddress(contractAddress)
	contract, err := history_contract.NewHistoryContract(addr, ethCli)
	if err != nil {
		return err
	}

	opt, err := newTransactOpts(context.Background(), ethCli)
	if err != nil {
		return err
	}

	if _, err := contract.StoreHistory(opt, "test data"); err != nil {
		return err
	}

	fmt.Println("success")
	return nil
}

func getHistory() error {
	ethCli, err := newEthCli(blockChainHost)
	if err != nil {
		return err
	}

	addr := common.HexToAddress(contractAddress)
	contract, err := history_contract.NewHistoryContract(addr, ethCli)
	if err != nil {
		return err
	}

	res, err := contract.GetHistory(&bind.CallOpts{
		From: common.HexToAddress(ownerAddress),
	}, big.NewInt(1))
	if err != nil {
		return err
	}
	fmt.Println(res)

	fmt.Println("success")
	return nil
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
