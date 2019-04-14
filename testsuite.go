package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	outpost "./contract"
)

// Start Ganache with:
// ganache-cli --account \
//   "0xa1d63a5f23ac9b62199e84d87fff196c603b61f6c42bddd0bcca9839d7449ba7,100000000000000000000"
const GanacheEndpoint = "http://127.0.0.1:8545"
const AdminPrivateKey = "a1d63a5f23ac9b62199e84d87fff196c603b61f6c42bddd0bcca9839d7449ba7"

func main() {
	client, err := ethclient.Dial(GanacheEndpoint)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	fmt.Println("Connection established")

	privateKey, err := crypto.HexToECDSA(AdminPrivateKey)
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

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	_, _, instance, err := outpost.DeployOutpost(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	pings, err := instance.Pings(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(pings)
}
