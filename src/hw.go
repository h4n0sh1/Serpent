package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

func encryptFile(filePath string, key []byte) error {
	// Read the input file
	plaintext, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read input file: %w", err)
	}

	// Create AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		return fmt.Errorf("failed to create cipher: %w", err)
	}

	// Create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return fmt.Errorf("failed to create GCM: %w", err)
	}

	// Generate nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return fmt.Errorf("failed to generate nonce: %w", err)
	}

	// Encrypt the data
	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil)

	// Write the encrypted data back to the file
	return os.WriteFile(filePath, ciphertext, 0644)
}

func copyFolderContents(srcDir, dstDir string) error {
	// Ensure destination directory exists
	if err := os.MkdirAll(dstDir, 0755); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}

	// List files in source folder
	entries, err := os.ReadDir(srcDir)
	if err != nil {
		return fmt.Errorf("failed to read source directory: %w", err)
	}

	// Copy each file to destination folder
	for _, entry := range entries {
		if entry.IsDir() {
			continue // Skip directories for now
		}
		srcPath := filepath.Join(srcDir, entry.Name())
		dstPath := filepath.Join(dstDir, entry.Name())

		err := copyFile(srcPath, dstPath)
		if err != nil {
			return fmt.Errorf("failed to copy %s: %w", entry.Name(), err)
		}
	}
	return nil
}

func encryptFolder(folderPath string, key []byte) error {
	// List files in the folder
	entries, err := os.ReadDir(folderPath)
	if err != nil {
		return fmt.Errorf("failed to read folder: %w", err)
	}

	// Encrypt each file
	for _, entry := range entries {
		if entry.IsDir() {
			continue // Skip directories
		}
		filePath := filepath.Join(folderPath, entry.Name())
		err := encryptFile(filePath, key)
		if err != nil {
			return fmt.Errorf("failed to encrypt %s: %w", entry.Name(), err)
		}
	}
	return nil
}

func copyFile(src, dst string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dstFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	_, err = io.Copy(dstFile, srcFile)
	return err
}

func main() {
	fmt.Println("Hello, World!")
}
