package md5

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"testing"
)

func Benchmark_md5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := md5.New()
		io.WriteString(h, "message digest")
		h.Sum(nil)
		// io.WriteString(h, randString())
		// id := hex.EncodeToString(h.Sum(nil))
	}
}

func Benchmark_sha1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := sha1.New()
		io.WriteString(h, "message digest")
		h.Sum(nil)
	}
}
func Benchmark_sha256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := sha256.New()
		io.WriteString(h, "message digest")
		h.Sum(nil)
	}
}
func Benchmark_sha512(b *testing.B) {
	for i := 0; i < b.N; i++ {
		h := sha512.New()
		io.WriteString(h, "message digest")
		h.Sum(nil)
	}
}
func EncryptAESCFB(dst, src, key, iv []byte) error {
	block, err := aes.NewCipher(key)
	if err != nil {
		return err
	}
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(dst, src)
	return nil
}

func DecryptAESCFB(dst, src, key, iv []byte) error {
	aesBlockDecrypter, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil
	}
	aesDecrypter := cipher.NewCFBDecrypter(aesBlockDecrypter, iv)
	aesDecrypter.XORKeyStream(dst, src)
	return nil
}
func Benchmark_aes(b *testing.B) {
	key, _ := hex.DecodeString("000102030405060708090a0b0c0d0e0f")
	var msg = "Message"
	fmt.Println([]byte(msg))
	iv, _ := hex.DecodeString("101112131415161718191a1b1c1d1e1f") //  []byte("101112131415161718191a1b1c1d1e1f")[:aes.BlockSize]
	var err error

	// Encrypt
	encrypted := make([]byte, len(msg))
	err = EncryptAESCFB(encrypted, []byte(msg), []byte(key), []byte(iv))
	if err != nil {
		panic(err)
	}
	fmt.Println(hex.EncodeToString(encrypted))

	for i := 0; i < b.N; i++ {
		const key16 = "1234567890123456"
		const key24 = "123456789012345678901234"
		const key32 = "12345678901234567890123456789012"
		var key = key16
		var msg = "message"
		var iv = []byte(key)[:aes.BlockSize] // Using IV same as key is probably bad
		var err error

		// Encrypt
		encrypted := make([]byte, len(msg))
		err = EncryptAESCFB(encrypted, []byte(msg), []byte(key), iv)
		if err != nil {
			panic(err)
		}
		// fmt.Printf("Encrypting %v %s -> %v\n", []byte(msg), msg, encrypted)

		// Decrypt
		decrypted := make([]byte, len(msg))
		err = DecryptAESCFB(decrypted, encrypted, []byte(key), iv)
		if err != nil {
			panic(err)
		}
	}
}
