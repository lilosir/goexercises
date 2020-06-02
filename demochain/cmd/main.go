package main

import "demochain/core"

func main() {
	bc := core.NewBlockChain()
	bc.SendData("send 1 btc to Oliver")
	bc.SendData("send 1 EOS to Oliver")
	bc.Print()
}
