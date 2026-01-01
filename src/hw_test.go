package main

import (
	"testing"

	"hw/pkg/encrypt"
)

func TestEncryptAndCopy(t *testing.T) {
	// Copy contents from test to val folder
	err := encrypt.CopyFolderContents("test", "val")
	if err != nil {
		t.Fatalf("Error copying folders: %v", err)
	}

	// Encrypt the contents of val folder
	key := []byte("a-32-byte-key-for-AES-256-!!!!!!") // 32 bytes
	err = encrypt.EncryptFolder("val", key)
	if err != nil {
		t.Fatalf("Error encrypting val folder: %v", err)
	}
}
