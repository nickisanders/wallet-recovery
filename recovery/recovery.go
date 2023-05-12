package recovery

import (
	"fmt"
	"os"

	"github.com/hashicorp/vault/shamir"
)

func Split(secret []byte, parts int, threshold int) [][]byte {
	// Split the secret into shares using Shamir's Secret Sharing algorithm from Vault
	shares, err := shamir.Split(secret, parts, threshold)
	if err != nil {
		panic(err)
	}

	// save shares
	for index, share := range shares {
		filename := fmt.Sprintf("testFiles/share%d", index)
		err := saveToFile(share, filename)
		if err != nil {
			fmt.Printf("Error saving file: %v\n", err)
		} else {
			fmt.Printf("File saved successfully!\n")
		}
	}

	return shares
}

func Combine(shares [][]byte) []byte {
	secret, err := shamir.Combine(shares)
	if err != nil {
		panic(err)
	}

	return secret
}

func saveToFile(data []byte, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	return nil
}
