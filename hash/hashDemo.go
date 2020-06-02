package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	fmt.Println(getHashString("hellolilosir"))
}

func getHashString(str string) string {
	hashInBytes := sha256.Sum256([]byte(str))
	return hex.EncodeToString(hashInBytes[:])
}
