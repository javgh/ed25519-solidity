package main

import (
	"context"
	"crypto/ecdsa"
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

	solidityEd25519 "./contract"
)

const GanacheEndpoint = "http://127.0.0.1:8545"
const MainPrivateKey = "a1d63a5f23ac9b62199e84d87fff196c603b61f6c42bddd0bcca9839d7449ba7"

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
	fmt.Printf("  ganache-cli --account \"0x%s,100000000000000000000\"\n", MainPrivateKey)

	// Connect to local Ethereum testnet and prepare smart contract deployment
	client, err := ethclient.Dial(GanacheEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	privateKey, err := crypto.HexToECDSA(MainPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(nil)
	}

	// Deploy smart contract
	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(6700000)
	auth.GasPrice = gasPrice

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

		adaptorBigInt := new(big.Int)
		adaptorBigInt.SetBytes(toBigEndian(adaptor[:]))
		adaptorPointBigInt := new(big.Int)
		adaptorPointBytes := toBigEndian(adaptorPoint[:])
		adaptorPointBytes[0] &= 127 // clear sign bit
		adaptorPointBigInt.SetBytes(adaptorPointBytes)

		_, adaptorPointSolidity, err := instance.Scalarmult(nil, adaptorBigInt)
		if err != nil {
			log.Fatal(err)
		}

		var estimateData []byte
		estimateData = append(estimateData, hexutil.MustDecode("0xe49cf911")...) // scalarmult
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
