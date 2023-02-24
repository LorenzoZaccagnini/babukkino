// note: utilizziamo xchacha20poly1305

package main

import (
	"babukkino/dec"
	"babukkino/enc"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"runtime"
)

func main() {
	// mostra numero CPU disponibili
	println("Numero CPU disponibili: ", runtime.NumCPU())

	// se passo -d come argomento decripta i file
	if len(os.Args) > 1 && os.Args[1] == "-d" {
		encryptedFiles, err := os.ReadDir("./cifrati")
		if err != nil {
			println("Errore nella lettura della cartella")
			log.Fatal(err)
		}

		for _, file := range encryptedFiles {

			//read key from file
			key, err := os.ReadFile("./key.txt")
			if err != nil {
				log.Fatal(err)
			}

			println("sto decifrando: ", file.Name())
			err = dec.DecryptFile("./cifrati/"+file.Name(), file.Name(), key)
			if err != nil {
				log.Fatal(err)
			}
		}
		return
	}

	// se passo -c come argomento cifra i file
	if len(os.Args) > 1 && os.Args[1] == "-c" {

		// generazione chiave simmetrica
		key := make([]byte, 32)
		if _, err := rand.Read(key); err != nil {
			log.Fatal(err)
		}

		//mostra chiave simmetrica
		keyStr := hex.EncodeToString(key)
		fmt.Println("Key:", keyStr)

		// salva su file la chiave simmetrica
		if err := os.WriteFile("./key.txt", key, 0644); err != nil {
			log.Fatal(err)
		}

		// Mostrami i nomi dei file nella cartella
		plainFiles, err := os.ReadDir("./ciframi")
		if err != nil {
			println("Errore nella lettura della cartella")
			log.Fatal(err)
		}

		for _, file := range plainFiles {
			println("sto cifrando: ", file.Name())
			enc.EncryptFile("./ciframi/"+file.Name(), file.Name(), key)
		}

		println("SEI Finito 666")
	}

	println("Non hai passato nessun argomento, che voi da me, te posso canta na canzone...")
}
