package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"

	"github.com/decred/dcrd/dcrec/secp256k1/v4"
	"golang.org/x/crypto/ripemd160"
)

func GenerateWif(privKeyInt *big.Int) string {
	privKeyHex := fmt.Sprintf("%064x", privKeyInt)

	privKeyBytes, err := hex.DecodeString(privKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	extendedKey := append([]byte{byte(0x80)}, privKeyBytes...)
	extendedKey = append(extendedKey, byte(0x01))

	firstSHA := sha256.Sum256(extendedKey)
	secondSHA := sha256.Sum256(firstSHA[:])
	checksum := secondSHA[:4]

	finalKey := append(extendedKey, checksum...)

	wif := Encode(finalKey)

	return wif
}

func CreatePublicHash160(privKeyInt *big.Int) []byte {

	privKeyBytes := privKeyInt.Bytes()

	privKey := secp256k1.PrivKeyFromBytes(privKeyBytes)

	compressedPubKey := privKey.PubKey().SerializeCompressed()

	pubKeyHash := hash160(compressedPubKey)

	return pubKeyHash

}

func checksum(payload []byte) []byte {
	hash1 := sha256.Sum256(payload)
	hash2 := sha256.Sum256(hash1[:])
	return hash2[:4]
}

func Hash160ToAddress(hash160 []byte) string {
	versionedPayload := append([]byte{0x00}, hash160...)
	checksum := checksum(versionedPayload)
	fullPayload := append(versionedPayload, checksum...)
	return Encode(fullPayload)
}

func hash160(b []byte) []byte {
	h := sha256.New()
	h.Write(b)
	sha256Hash := h.Sum(nil)

	r := ripemd160.New()
	r.Write(sha256Hash)
	return r.Sum(nil)
}
