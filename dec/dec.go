package dec

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/chacha20poly1305"
)

func DecryptFile(path string, name string, key []byte) error {
	// -----------------------------
	// - 1 leggere il file cifrato
	// - 2 decifrarlo con la chiave simmetrica e il nonce nel file
	// - 3 salvare il file decifrato con il nome originale
	// note: utilizziamo xchacha20poly1305
	// -----------------------------

	// leggere il file cifrato
	noExtName := strings.Split(name, ".")[0]
	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return err
	}

	defer f.Close()

	// 2 decifrarlo con la chiave simmetrica
	ciphertext, err := ioutil.ReadAll(f)
	if err != nil {
		return err
	}

	// crea una cifratura XChaCha20-Poly1305
	cipher, err := chacha20poly1305.NewX(key)
	if err != nil {
		return err
	}

	// prendi il nonce dal file
	nonce := ciphertext[len(ciphertext)-chacha20poly1305.NonceSizeX:]
	ciphertext = ciphertext[:len(ciphertext)-chacha20poly1305.NonceSizeX]

	// mostra nonce come stringa
	nonceStr := hex.EncodeToString(nonce)
	fmt.Println("Nonce:", nonceStr)

	// decifra il file
	plaintext, err := cipher.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return err
	}

	// salvare il file decifrato con il nome originale + .txt
	// crea la cartella decifrati se non esiste

	if _, err := os.Stat("./decifrati"); os.IsNotExist(err) {
		os.Mkdir("./decifrati", 0755)
	}

	// Crea il file decifrato
	fdec, err := os.OpenFile("./decifrati/"+noExtName+".txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		return err
	}

	defer fdec.Close()

	if _, err := fdec.Write(plaintext); err != nil {
		panic(err)
	}

	return nil
}
