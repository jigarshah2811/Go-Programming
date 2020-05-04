package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
)

func main()  {
	BlockChainExample()
}

type Transaction struct {
	user string
	amount int
}

type Block struct {
	previousBlock *Block
	previousHash string
	nonce int
	body []Transaction
}
var blockchain []Block

func CreateHash(b Block) string  {
	plainText := strconv.Itoa(b.nonce)

	for _, t := range b.body {
		plainText += t.user
		plainText += strconv.Itoa(t.amount)
	}

	data := sha256.Sum256([]byte(plainText))
	hash := hex.EncodeToString(data[:])
	return hash
}

func (thisBlock *Block) Add(previousBlock *Block)  {
	if previousBlock == nil {
		thisBlock.previousBlock = nil
		thisBlock.previousHash = ""
		return
	}

	thisBlock.previousBlock = previousBlock
	thisBlock.previousHash = CreateHash(*previousBlock)
}

func VerifyBlockChain() bool {

	for _, b := range blockchain {
		previousHash := CreateHash(*b.previousBlock)
		if strings.Compare(previousHash, b.previousHash) != 0 {
			return false
		}
	}
	return true
}

func (thisBlock *Block) Mine()  {
	blockchain = append(blockchain, *thisBlock)
}

func BlockChainExample()  {

	var block1, block2 Block

	t1b1 := Transaction{user: "Jigar", amount: 10}
	t1b2 := Transaction{user: "Krupa", amount: 20}
	t2b2 := Transaction{user: "Parshvi", amount: 30}

	block1.body = append(block1.body, t1b1)
	block1.Add(nil)

	block2.body = append(block2.body, t1b2)
	block2.body = append(block2.body, t2b2)
	block2.Add(&block1)


	/* Verification */
	fmt.Println(block1)
	fmt.Println(block2)
	if VerifyBlockChain() {
		fmt.Println("VerifyBlockChain succeed")
	} else {
		fmt.Println("VerifyBlockChain failed!")
	}


}
