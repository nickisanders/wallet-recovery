package test

import (
	"testing"

	"wallet-recovery/recovery"
)

func TestSplitMessage(t *testing.T) {
	secret := "never gonna give you up never gonna let you down"
	msg := recovery.Split([]byte(secret), 3, 2)

	shareOne := []byte{155, 91, 183, 188, 40, 197, 254, 48, 241, 72, 251, 215, 90, 138, 186, 43, 83, 210, 56, 88, 218, 90, 158, 66}
	shareTwo := []byte{122, 52, 85, 206, 9, 196, 33, 39, 10, 183, 54, 42, 39, 175, 30, 210, 215, 42, 95, 197, 97, 219, 253, 243}
	shareThree := []byte{41, 215, 227, 112, 63, 110, 86, 18, 127, 92, 64, 142, 197, 7, 90, 232, 15, 186, 155, 222, 55, 55, 167, 205}
	shares := [][]byte{shareOne, shareTwo, shareThree}

	t.Fatalf(`Split(%v) = %v, want %v, error`, secret, msg, shares)
	t.Fatalf(`Split("never gonna give you up never gonna let you down") = %v, want %v, error`, msg, shares)
}
