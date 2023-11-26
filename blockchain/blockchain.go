package blockchain

import (
    "fmt"
    "crypto/sha256"
    "strconv"
    "encoding/hex"
    "time"
)

type Block struct {
    index int
    timestamp string
    proof int
    previous_hash string

}

func NewBlock(index int, timestamp string, proof int, previous_hash string) Block {
    return Block { index, timestamp, proof, previous_hash }
}

func (b Block) get_hash() string {
    // h := sha256.New()
    // s := fmt.Sprintf("%v", b)
    // sum := h.Sum([]byte(s))
    // // fmt.Println(sum)
    blockBytes := []byte( fmt.Sprintf("%d%s%d%s", b. index, b.timestamp, b.proof, b.previous_hash))
    hashBytes := sha256.Sum256(blockBytes)
    hashHex := hex.EncodeToString(hashBytes[:])
    return hashHex
}

type Blockchain struct {
    chain []Block
}

func NewBlockchain() *Blockchain {
    bc := Blockchain { make([]Block, 0) }
    _ = bc.create_block(1, "0")
    return &bc
}

func (bc *Blockchain) create_block(proof int, previous_hash string) Block {
    // time.Now().Format("20060102150405")
    block := NewBlock(len(bc.chain) + 1, time.Now().Format("20060102150405"), proof, previous_hash)
    bc.chain = append(bc.chain, block)
    return block
}

func (b Blockchain) get_previous_block() Block {
    return b.chain[len(b.chain) - 1]
}

func IntPow(base, exp int) int {
    result := 1
    for {
        if exp & 1 == 1 {
            result *= base
        }
        exp >>= 1
        if exp == 0 {
            break
        }
        base *= base
    }

    return result
}

func (b Blockchain) proof_of_work(previous_proof int) int {
    var new_proof int = 1
    var check_proof bool = false

    for check_proof == false {
        var operation = IntPow(new_proof, 2) - IntPow(previous_proof, 2)
        var hash_operation = sha256.Sum256([]byte(strconv.Itoa(operation)))

        var encoded_hash = hex.EncodeToString(hash_operation[:])

        if encoded_hash[:4] == "0000" {
            check_proof = true
        } else {
            new_proof += 1
        }
    }
    return new_proof
}

func (b Blockchain) Is_chain_valid() bool {
    var previous_block = b.chain[0]
    var block_index = 1

    for block_index < len(b.chain) {
        var block = b.chain[block_index]

        if block.previous_hash != previous_block.get_hash() {
            return false
        }

        var previous_proof = previous_block.proof
        var proof = block.proof

        var operation = IntPow(proof, 2) - IntPow(previous_proof, 2)
        var hash_operation = sha256.Sum256([]byte(strconv.Itoa(operation)))

        var encoded_hash = hex.EncodeToString(hash_operation[:])

        if encoded_hash[:4] != "0000" {
            return false
        }

        previous_block = block
        block_index += 1
    }

    return true
}

func (bc Blockchain) PrintChain() {
    index := 0
    fmt.Println("Chain: {")
    for index < len(bc.chain) {
        block := bc.chain[index]
        fmt.Printf("\t[index: %d, timestamp: %s, proof: %d, prev_hash: %s]\n", block.index, block.timestamp, block.proof, block.previous_hash)
        index += 1
    }
    fmt.Println("}")
}

func Mine_block(bc *Blockchain) {
   var previous_block = bc.get_previous_block()
   var previous_proof = previous_block.proof
   var proof = bc.proof_of_work(previous_proof)
   var previous_hash = previous_block.get_hash()

   _ = bc.create_block(proof, previous_hash)
}

