# wallet-recovery
Utility to facilitate a recoverable crypto wallet

## Wallet Recovery CLI Tool

This command-line tool provides functionalities to split and combine a secret using the Shamir's Secret Sharing algorithm from Vault.

### Installation

To use this tool, you need to have Go installed on your machine.

Then, you can clone this repository and build the binary using the following command:

```
go build -o wallet-recovery
```

### Usage

The command-line tool takes the following flags:

* `action`: (required) The action to perform. It can be either "split" or "combine".
* `inputFile`: (required for "split") The path to the input file containing the secret to split.
* `parts`: (required for "split") The number of parts to split the secret into.
* `threshold`: (required for "split" and "combine") The minimum number of parts required to combine the secret.
* `[shareFiles]`: (required for "combine") The paths to the share files to combine.

To split a secret, use the following command:

```
./wallet-recovery -action=split -inputFile=/path/to/input/file -parts=3 -threshold=2
```

This command splits the secret in the input file into 3 parts, requiring at least 2 parts to combine.

To combine a secret, use the following command:

```
./wallet-recovery -action=combine -threshold=2 /path/to/share/file1 /path/to/share/file2 ...
```

This command combines the share files into the original secret, requiring at least 2 parts to be present.
