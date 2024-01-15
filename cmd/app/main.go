package main

import (
	"bingo/internal/blockchain"
	"bingo/internal/wallet"
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	walletM := wallet.NewWallet()
	walletA := wallet.NewWallet()
	walletB := wallet.NewWallet()

	t := wallet.NewTransaction(walletA.PrivateKey(), walletA.PublicKey(), walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0)

	// Blockchain
	blockchain := blockchain.NewBlockchain(walletM.BlockchainAddress())

	isAdded := blockchain.AddTransaction(walletA.BlockchainAddress(), walletB.BlockchainAddress(), 1.0, walletA.PublicKey(), t.GenerateSignature())

	fmt.Println("isAdded", isAdded)

	blockchain.Mining()
	blockchain.Print()
	fmt.Printf("B: %.1f \n", blockchain.CalculateTotalAmount(walletB.BlockchainAddress()))
	fmt.Printf("A: %.1f \n", blockchain.CalculateTotalAmount(walletA.BlockchainAddress()))
	fmt.Printf("M: %.1f \n", blockchain.CalculateTotalAmount(walletM.BlockchainAddress()))
	// fmt.Printf("signature: %s\n", t.GenerateSignature())
}
