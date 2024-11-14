package main

import (
	"bufio"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"

	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/ramonamorim/hex-to-wif-go/utils"
)

func main() {
	fmt.Print("Digite a chave privada em hexadecimal: ")
	reader := bufio.NewReader(os.Stdin)
	inputHex, _ := reader.ReadString('\n')

	inputHex = inputHex[:len(inputHex)-1]

	privKeyBytes, err := hex.DecodeString(inputHex)
	if err != nil {
		fmt.Println("Erro ao decodificar hexadecimal:", err)
		return
	}

	privKey, _ := btcec.PrivKeyFromBytes(privKeyBytes)

	privateKeyInt := new(big.Int).SetBytes(privKey.Serialize())
	fmt.Println("Wif: ", utils.GenerateWif(privateKeyInt))

}
