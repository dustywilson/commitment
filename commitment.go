package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/crypto/sha3"
)

func main() {
	flag.Parse()

	value := flag.Arg(0)
	nonce := flag.Arg(1)

	if len(value) == 0 || len(nonce) == 0 {
		fmt.Printf("Usage: %s <value> <nonce>\n", os.Args[0])
		os.Exit(1)
	}

	shake := sha3.NewShake256()
	shake.Write([]byte(value))
	shake.Write([]byte("|"))
	shake.Write([]byte(nonce))

	hash := make([]byte, 16)
	_, err := shake.Read(hash)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("%X-%X-%X-%X\n", hash[0:3], hash[4:7], hash[8:11], hash[12:15])
}
