package main

import (
    "fmt"
    "time"
    . "GoBlock/blockchain"
)

func main() {
    var bc *Blockchain = NewBlockchain()
    Mine_block(bc)
    Mine_block(bc)
    Mine_block(bc)
    Mine_block(bc)
    Mine_block(bc)
    Mine_block(bc)
    Mine_block(bc)
    time.Sleep(1 * time.Second)
    Mine_block(bc)
    Mine_block(bc)
    Mine_block(bc)
    Mine_block(bc)
    Mine_block(bc)
    bc.PrintChain()

    is_valid := bc.Is_chain_valid()
    fmt.Println("Is chain valid:", is_valid)
}
