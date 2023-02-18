package crypto

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"testing"

	util "github.com/bulwarkid/virtual-fido/util"
)

func TestEncryptDecrypt(t *testing.T) {
	data := []byte("data")
	key := GenerateSymmetricKey()
	encryptedData, nonce, err := Encrypt(key, data)
	if err != nil {
		t.Fatal(err)
	}
	decryptedData, err := Decrypt(key, encryptedData, nonce)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(decryptedData, data) {
		t.Fatalf("'%s' does not match '%s'", string(decryptedData), string(data))
	}
}

func TestSignVerify(t *testing.T) {
	data := []byte("data")
	key := GenerateECDSAKey()
	signature := Sign(key, data)
	if !Verify(&key.PublicKey, data, signature) {
		t.Fatalf("Signature not correct: %#v", signature)
	}
}

func TestSealOpen(t *testing.T) {
	data := []byte("data")
	key := GenerateSymmetricKey()
	box := Seal(key, data)
	decryptedData := Open(key, box)
	if !bytes.Equal(data, decryptedData) {
		t.Fatalf("'%s' does not equal '%s'", decryptedData, data)
	}
}

func TestHashSHA256(t *testing.T) {
	data := []byte("test")
	target := "9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"
	hash := HashSHA256(data)
	encodedHash := hex.EncodeToString(hash)
	if encodedHash != target {
		t.Fatalf("'%s' does not equal '%s'", encodedHash, target)
	}
}

func TestEncryptDecryptAESCBC(t *testing.T) {
	data := util.Read(rand.Reader, 32)
	key := GenerateSymmetricKey()
	encryptedData := EncryptAESCBC(key, data)
	decryptedData := DecryptAESCBC(key, encryptedData)
	if !bytes.Equal(data, decryptedData) {
		t.Fatalf("'%s' does not equal '%s'", hex.EncodeToString(decryptedData), hex.EncodeToString(data))
	}
}