package enc

import (
	"crypto/rand"
	"io"
	"log"
	"os"
	"strings"

	"golang.org/x/crypto/chacha20poly1305"
)

func EncryptFile(path string, name string, key []byte) error {
	// TODO
	// - 1 leggere il file
	// - 2 cifrarlo con un nonce differente
	// - 3 crea la cartella se non esiste
	// - 4 salvare il file cifrato con il nome originale + .kino
	// note: utilizziamo xchacha20poly1305
	inputFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return err
	}

	defer inputFile.Close()

	nonce := make([]byte, chacha20poly1305.NonceSizeX)
	if _, err := rand.Read(nonce); err != nil {
		log.Fatal(err)
		return err
	}

	// crea la cartella se non esiste
	if _, err := os.Stat("./cifrati"); os.IsNotExist(err) {
		os.Mkdir("./cifrati", 0755)
	}

	noExtName := strings.Split(name, ".")[0]

	uCantCMeFile, err := os.OpenFile("./cifrati/"+noExtName+".kino", os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
		return err
	}

	defer uCantCMeFile.Close()

	if _, err := uCantCMeFile.Write(nonce); err != nil {
		log.Fatal(err)
		return err
	}

	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		log.Fatal(err)
		return err
	}

	buf := make([]byte, 1024*32)
	ad_counter := 0

	for {
		// leggi il file
		n, err := inputFile.Read(buf)

		if n > 0 {
			msg := buf[:n]
			// cifra il file
			ciphertext := aead.Seal(nil, nonce, msg, []byte(string(ad_counter)))
			// scrivi il file cifrato
			if _, err := uCantCMeFile.Write(ciphertext); err != nil {
				log.Fatal(err)
				return err
			}

			ad_counter++
		}

		if err == io.EOF {
			break
		}

		if err != nil {
			panic(err)
		}
	}

	return nil
}
