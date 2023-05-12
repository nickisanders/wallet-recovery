package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"wallet-recovery/recovery"
)

func main() {
	// Define command-line flags
	var (
		action    = flag.String("action", "", "Action to perform: split or combine")
		inputFile = flag.String("inputFile", "", "Input file path")
		parts     = flag.Int("parts", 0, "Number of parts to split the secret into")
		threshold = flag.Int("threshold", 0, "Minimum number of parts required to combine the secret")
	)
	flag.Parse()

	// Validate flags
	if *action == "" {
		fmt.Fprintln(os.Stderr, "Error: action is required")
		os.Exit(1)
	}
	if *action != "split" && *action != "combine" {
		fmt.Fprintln(os.Stderr, "Error: action must be 'split' or 'combine'")
		os.Exit(1)
	}
	if *action == "split" {
		if *inputFile == "" {
			fmt.Fprintln(os.Stderr, "Error: inputFile is required")
			os.Exit(1)
		}
		if *parts <= 0 {
			fmt.Fprintln(os.Stderr, "Error: parts must be a positive integer")
			os.Exit(1)
		}
		if *threshold <= 0 {
			fmt.Fprintln(os.Stderr, "Error: threshold must be a positive integer")
			os.Exit(1)
		}
	}

	// Perform action
	switch *action {
	case "split":
		// Read input file
		secret, err := ioutil.ReadFile(*inputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: failed to read input file: %v\n", err)
			os.Exit(1)
		}

		output := recovery.Split(secret, *parts, *threshold)
		fmt.Println(output)
	case "combine":
		var shares [][]byte
		for _, arg := range flag.Args() {
			share, err := ioutil.ReadFile(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: failed to read share file: %v\n", err)
				os.Exit(1)
			}
			shares = append(shares, share)
		}
		output := recovery.Combine(shares)
		fmt.Println(string(output))
	default:
		fmt.Println("No command given")
	}
}
