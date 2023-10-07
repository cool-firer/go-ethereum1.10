package rlp

import (
	"fmt"
	"testing"
	"bytes"
	"encoding/hex"
	"math/big"
)

func TestEcodeUint(t *testing.T) {
	b := new(bytes.Buffer)
	var val interface{}  = uint32(256)
	Encode(b, val)
	fmt.Println(hex.Dump(b.Bytes()))
}

func TestEncodeBigInt(t *testing.T) {
	s, _ := hex.DecodeString("102030405060708090A0B0C0D0E0F2")
	var val interface{}  = new(big.Int).SetBytes(s)

	b := new(bytes.Buffer)
	Encode(b, val)
	fmt.Println(hex.Dump(b.Bytes()))


	b.Reset()
	s, _ = hex.DecodeString("102030405060708090A0B0C0D0E0F2102030405060708090A0B0C0D0E0F2102030405060708090A0B0C0D0E0F2102030405060708090A0B0C0D0E0F2")
	val = new(big.Int).SetBytes(s)
	Encode(b, val)
	fmt.Println(hex.Dump(b.Bytes()))
}