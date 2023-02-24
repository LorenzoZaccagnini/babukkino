package dec

import (
	"crypto/rand"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/chacha20poly1305"
)

func DecryptFile(path string, name string, key []byte) error {
	// TODO
	// - 1 leggere il file cifrato
	// - 2 decifrarlo con la chiave simmetrica
	// - 3 salvare il file decifrato con il nome originale + .txt
	// note: utilizziamo xchacha20poly1305

	// 1 leggere il file cifrato
	noExtName := strings.Split(name, ".")[0]

	lockedFile, err := os.Open(path + noExtName + ".kino")
	if err != nil {
		log.Fatal(err)
		return err
	}

	defer lockedFile.Close()

	// 2 decifrarlo con la chiave simmetrica
	nonce := make([]byte, chacha20poly1305.NonceSizeX)
	if _, err := rand.Read(nonce); err != nil {
		log.Fatal(err)
		return err
	}

	// 3 salvare il file decifrato con il nome originale + .txt
	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		log.Fatal(err)
		return err
	}

	buf := make([]byte, 1024*32)
	ad_counter := 0

	// 4 crea la cartella se non esiste
	if _, err := os.Stat("./decifrati"); os.IsNotExist(err) {
		os.Mkdir("./decifrati", 0755)
	}

	// 5 apri il file decifrato
	decryptedFile, err := os.OpenFile("./decifrati/"+noExtName+".txt", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
		return err
	}

	for {
		n, err := lockedFile.Read(buf)
		if n > 0 {
			if err != nil {
				break
			}

			plainFile, err := aead.Open(nil, nonce, buf, []byte(string(ad_counter)))
			if err != nil {
				println("error plainFile")
				log.Fatal(err)
				return err
			}
			ad_counter++

			if _, err := decryptedFile.Write(plainFile); err != nil {
				log.Fatal(err)
				return err
			}

			//log plain text
			log.Println(string(plainFile))
		}

	}

	return nil
}
