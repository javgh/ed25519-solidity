package main

import (
	"context"
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"golang.org/x/crypto/curve25519"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	solidityCurve25519 "./contract"
)

// Start Ganache with:
// ganache-cli --account \
//   "0xa1d63a5f23ac9b62199e84d87fff196c603b61f6c42bddd0bcca9839d7449ba7,100000000000000000000"
const GanacheEndpoint = "http://127.0.0.1:8545"
const MainPrivateKey = "a1d63a5f23ac9b62199e84d87fff196c603b61f6c42bddd0bcca9839d7449ba7"

func toBigEndian(littleEndian []byte) []byte {
	bigEndian := make([]byte, len(littleEndian))
	for i := 0; i < len(littleEndian); i++ {
		bigEndian[i] = littleEndian[len(littleEndian)-1-i]
	}
	return bigEndian
}

func main() {
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
	auth.GasLimit = uint64(500000)
	auth.GasPrice = gasPrice

	_, _, instance, err := solidityCurve25519.DeployCurve25519(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	// Pick a random key and compare Go's curve25519
	// implementation with the Solidity code.
	var in, dst [32]byte
	_, err = rand.Read(in[:])
	if err != nil {
		log.Fatal(err)
	}
	curve25519.ScalarBaseMult(&dst, &in)

	solidityIn := new(big.Int)
	solidityIn.SetBytes(toBigEndian(in[:]))

	solidityDst, err := instance.PublicKey(nil, solidityIn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Tested %s (big endian)\n", solidityIn.Text(16))
	fmt.Printf("  Go code      : %s\n", hex.EncodeToString(dst[:]))
	fmt.Printf("  Go code      : %s\n", hex.EncodeToString(toBigEndian(dst[:])))
	fmt.Printf("  Solidity code: %s\n", solidityDst.Qx.Text(16))
	fmt.Printf("  Solidity code: %s\n", solidityDst.Qy.Text(16))
}
