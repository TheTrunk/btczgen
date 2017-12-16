package main

import (
	"flag"
	"log"
	"fmt"
	"bufio"
	"os"

	"github.com/TheTrunk/btczgen/btczcrypto"
)

func main() {
	//	var networkId btczcrypto.NetworkId
	boolPtr := flag.Bool("test", false, "generate a testnet wallet")
	nPtr := flag.Int("n", 1, "Number of addresses to generate up to 100")
	boolPtr3 := flag.Bool("o", false, "enable output to file outputbtczgen.txt")
	flag.Parse()

	var output bool = *boolPtr3

	// Generate the wallet
	wallet, err := btczcrypto.CreateWallet(!(*boolPtr), *nPtr)

	if err != nil {
		log.Panicln(err.Error())
	}

	log.Println("Wallet generated!")
	fmt.Println("Passphrase:", wallet.Passphrase)
	fmt.Println("Address\t\t\t\tPrivate key")

		file, err := os.OpenFile("outputbtczgen.txt", os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil && output == true {
        fmt.Println("File does not exists or cannot be created")
        os.Exit(1)
		}
	w := bufio.NewWriter(file)
	if output == true {
	fmt.Fprintln(w,"Passphrase:", wallet.Passphrase)
	fmt.Fprintln(w,"Address\t\t\t\t\t\t\t\tPrivate key")
	w.Flush()
	}

	for i := 0; i <= len(wallet.Addresses)-1; i++ {
		fmt.Println(wallet.Addresses[i].Value, wallet.Addresses[i].PrivateKey)
			if output == true {
				fmt.Fprintln(w,wallet.Addresses[i].Value, wallet.Addresses[i].PrivateKey)
				w.Flush()
			}
	}
}
