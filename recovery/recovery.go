package recovery

import (
	"github.com/hashicorp/vault/shamir"
)

func Split(secret []byte, parts int, threshold int) [][]byte {
	// Split the secret into shares using Shamir's Secret Sharing algorithm from Vault
	shares, err := shamir.Split(secret, parts, threshold)
	if err != nil {
		panic(err)
	}

	//TODO: save shares

	return shares
}

func Combine(shares [][]byte) []byte {
	secret, err := shamir.Combine(shares)
	if err != nil {
		panic(err)
	}

	return secret
}
