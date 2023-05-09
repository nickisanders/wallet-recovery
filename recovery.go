package main

import (
	"fmt"

	"github.com/hashicorp/vault/shamir"
)

func main() {
	secret := []byte("never gonna give you up never gonna let you down")
	parts := 4
	threshold := 3

	fmt.Println("Secret:", secret)

	shares := split(secret, parts, threshold)

	// Print out the shares
	fmt.Println("Shares:")
	for i, share := range shares {
		fmt.Printf("%d: %x\n", i, share)
	}

	shareArray := [][]byte{shares[0], shares[1], shares[2]}

	newSecret := string(combine(shareArray))

	fmt.Println("New secret: ", newSecret)
}

func split(secret []byte, parts int, threshold int) [][]byte {
	// Split the secret into shares using Shamir's Secret Sharing algorithm from Vault
	shares, err := shamir.Split(secret, parts, threshold)
	if err != nil {
		panic(err)
	}

	return shares
}

func combine(shares [][]byte) []byte {
	secret, err := shamir.Combine(shares)
	if err != nil {
		panic(err)
	}

	return secret
}
