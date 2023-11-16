package main

import (
	"context"
	"fmt"
	"log"
	"os"

	multibaas "github.com/curvegrid/multibaas-sdk-go"
)

// chainIDToERC20Addr maps chain IDs to random ERC20 contract addresses for the purpose of this example
var chainIDToERC20Addr = map[int64]string{
	1:          "0x6B175474E89094C44Da98b954EedeAC495271d0F", // Ethereum Mainnet
	5:          "0xB4FBF271143F4FBf7B91A5ded31805e42b2208d6", // Ethereum Goerli
	11155111:   "0x969D499507B4f437953Db24A4980FdEEDa6Db8a1", // Ethereum Sepolia
	137:        "0x7ceB23fD6bC0adD59E62ac25578270cFf1b9f619", // Polygon Mainnet
	80001:      "0x3C0d2375d317092F530caFb440337b4E88163f29", // Polygon Mumbai
	43114:      "0xd586E7F844cEa2F87f50152665BCbc2C279D8d70", // Avalanche Mainnet
	43113:      "0x12C135a68b7B3Cd006eDb785cB53398a5DA59613", // Avalanche Fuji
	2017072401: "0x8a58447d8AE49b6Ac6cf773B11a3981C46a6D89D", // Curvegrid Testnet
}

const chain = "ethereum"

func main() {
	/* Initializing the SDK */

	conf := multibaas.NewConfiguration()
	client := multibaas.NewAPIClient(conf)

	// Configure the SDK using environment variables
	ctx := context.Background()
	ctx = context.WithValue(ctx, multibaas.ContextServerVariables, map[string]string{
		"base_url": os.Getenv("MB_BASE_URL"),
	})
	ctx = context.WithValue(ctx, multibaas.ContextAccessToken, os.Getenv("MB_API_KEY"))

	/* Example 1: getting blockchain details */

	// First let's get the chain ID of the blockchain MultiBaas is connected to
	resp1, _, err := client.ChainsAPI.GetChainStatus(ctx, chain).Execute()
	if err != nil {
		log.Fatal(err)
	}

	chainId := resp1.Result.ChainID
	latestBlock := resp1.Result.BlockNumber

	fmt.Printf("Example 1: MultiBaas is connected to the chain ID #%d and the latest block number is #%d\n", chainId, latestBlock)

	/* Example 2: calling a smart contract function */

	// Now let's call an ERC20 contract to get the balance of an address
	walletAddr := "0xAb5801a7D398351b8bE11C439e05C5B3259aeC9B" // Vitalik Buterin's wallet
	contractLabel := "erc20interface"
	contractMethod := "balanceOf"
	contractOverride := true

	contractAddr, ok := chainIDToERC20Addr[chainId]
	if !ok {
		log.Fatalf("This code example is missing an ERC20 contract address for the chain ID %d", chainId)
	}

	payload := multibaas.PostMethodArgs{
		Args: []interface{}{
			walletAddr,
		},
		ContractOverride: &contractOverride,
	}

	resp2, _, err := client.ContractsAPI.CallContractFunction(ctx, chain, contractAddr, contractLabel, contractMethod).PostMethodArgs(payload).Execute()
	if err != nil {
		log.Fatal(err)
	}

	balance := resp2.Result.MethodCallResponse.Output
	fmt.Printf("Example 2: The wallet %s balance on the ERC20 contract %s is: %v\n", walletAddr, contractAddr, balance)

	/* Example 3: handling errors */

	// Intentionally calling a contract method that doesn't exist to trigger an error
	_, _, err = client.ContractsAPI.CallContractFunction(ctx, chain, contractAddr, contractLabel, "thisMethodDoNotExist").PostMethodArgs(payload).Execute()
	if err != nil {
		if mbErr, ok := multibaas.IsMultiBaasErr(err); ok {
			fmt.Printf("Example 3: MultiBaas returned an error with status '%d' and message: '%s'\n", mbErr.Status, mbErr.Message)
		}
	}
}
