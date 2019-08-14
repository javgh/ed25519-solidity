package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"math/big"

	"github.com/HyperspaceApp/ed25519"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	solidityEd25519 "github.com/javgh/ed25519-solidity/contract"
)

const ganacheEndpoint = "http://127.0.0.1:8545"
const mainPrivateKey = "a1d63a5f23ac9b62199e84d87fff196c603b61f6c42bddd0bcca9839d7449ba7"

func toBigEndian(littleEndian []byte) []byte {
	bigEndian := make([]byte, len(littleEndian))
	for i := range littleEndian {
		bigEndian[i] = littleEndian[len(littleEndian)-1-i]
	}
	return bigEndian
}

func main() {
	var n = flag.Int("n", 3, "number of test cases to generate")
	flag.Parse()

	fmt.Printf("Please ensure Ganache has been started with this command:\n")
	fmt.Printf("  ganache-cli --account \"0x%s,100000000000000000000\"\n", mainPrivateKey)

	// Connect to local Ethereum testnet and prepare smart contract deployment
	client, err := ethclient.Dial(ganacheEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	privateKey, err := crypto.HexToECDSA(mainPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	// Deploy smart contract
	auth := bind.NewKeyedTransactor(privateKey)
	auth.GasLimit = uint64(6700000)

	address, _, instance, err := solidityEd25519.DeployEd25519(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < *n; i++ {
		// Create random keypair and compare results from Go and Solidity
		adaptor, adaptorPoint, err := ed25519.GenerateAdaptor(rand.Reader)
		if err != nil {
			log.Fatal(err)
		}

		adaptorBigInt := new(big.Int).SetBytes(toBigEndian(adaptor[:]))
		adaptorPointBytes := toBigEndian(adaptorPoint[:])
		adaptorPointBytes[0] &= 127 // clear sign bit
		adaptorPointBigInt := new(big.Int).SetBytes(adaptorPointBytes)

		_, adaptorPointSolidity, err := instance.ScalarMultBase(nil, adaptorBigInt)
		if err != nil {
			log.Fatal(err)
		}

		var estimateData []byte
		estimateData = append(estimateData, hexutil.MustDecode("0xc4f4912b")...) // scalarMultBase
		estimateData = append(estimateData, common.LeftPadBytes(adaptorBigInt.Bytes(), 32)...)

		gasEstimate, err := client.EstimateGas(context.Background(), ethereum.CallMsg{
			To:   &address,
			Data: estimateData,
		})
		if err != nil {
			log.Fatal(err)
		}

		wasSuccessful := adaptorPointBigInt.Cmp(adaptorPointSolidity) == 0
		if wasSuccessful {
			fmt.Printf("\nTest successful:\n")
		} else {
			fmt.Printf("\nTest failed:\n")
		}
		fmt.Printf("  Adaptor (little endian)                   : %s\n",
			hex.EncodeToString(adaptor[:]))
		fmt.Printf("  Adaptor (big int, hex)                    : %s\n",
			adaptorBigInt.Text(16))
		fmt.Printf("  Adaptor (big int, decimal)                : %s\n",
			adaptorBigInt.Text(10))
		fmt.Printf("\n")
		fmt.Printf("  Adaptor point (big int, decimal, Go)      : %s\n",
			adaptorPointBigInt.Text(10))
		fmt.Printf("  Adaptor point (big int, decimal, Solidity): %s\n",
			adaptorPointSolidity.Text(10))
		fmt.Printf("\n")
		fmt.Printf("  Gas estimate: %d\n", gasEstimate)

		if !wasSuccessful {
			log.Fatal("last test failed")
		}
	}
}
