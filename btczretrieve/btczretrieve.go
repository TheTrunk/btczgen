// retrieves the addresses and priv keys associated with a mneumonic
package main

import (
	"flag"
	"log"
	"math"
	"regexp"
	"os"
	"time"
	"fmt"
	"bufio"

	"github.com/TheTrunk/btczgen/btczcrypto"
)

func main() {
	boolPtr := flag.Bool("test", false, "generate a testnet wallet")
	strPtr := flag.String("passphrase", "", "Passphrase for the wallet is REQUIRED between 128 and 512 bits")
	nPtr := flag.Int("n", 1, "Number of addresses to retrieve")
	strPtr2 := flag.String("match", "", "generate addresses infinitely until a regex match is made to an address")
	boolPtr2 := flag.Bool("i", false, "case insensitive regex match")
	boolPtr3 := flag.Bool("o", false, "enable output to file outputbtczretrieve.txt")

	flag.Parse()
	var passphrase string = *strPtr
	var test bool = *boolPtr
	var numAddresses uint32
	var match string = *strPtr2
	var caseInsensitive bool = *boolPtr2
	var numGenerate int = int(*nPtr)
	var output bool = *boolPtr3

	if passphrase == "" {
		log.Fatalln("Passphrase must be specified")
	}

	log.Println("Wallet retrieved")
	fmt.Println("Passphrase:", passphrase)
	// Try up to max number represented in an unsigned 32 bit integer
	var reg *regexp.Regexp
	if match != "" {
		var err error
		numAddresses = math.MaxUint32

		var regexpString string
		if caseInsensitive == true {
			fmt.Println("Searching for an address case insensitive for pattern:", match)
			regexpString = "(?i)" + match
		} else {
			fmt.Println("Searching for an address case sensitive for pattern:", match)
			regexpString = match
		}
		reg, err = regexp.Compile(regexpString)

		if err != nil {
			log.Println("Invalid regex")
			log.Panicln(err.Error())
		}
	}

	file, err := os.OpenFile("outputbtczretrieve.txt", os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil && output == true {
        fmt.Println("File does not exists or cannot be created")
        os.Exit(1)
		}
	w := bufio.NewWriter(file)
	if output == true {
	fmt.Fprintln(w,"Passphrase:", passphrase)
	fmt.Fprintln(w,"Address\t\t\t\t\t\t\t\tPrivate key")
	w.Flush()
	}

	fmt.Println("Address\t\t\t\t\tPrivate key")

	var i uint32
	var a int
	start := time.Now()

	for i = 0; i <= numAddresses-1; i++ {

		wallet, err := btczcrypto.GetWalletFromPassphrase(!test, passphrase, uint32(i))


		if err != nil {
			log.Panicln(err.Error())
		}

		if match != "" {
			if reg.MatchString(wallet.Addresses[0].Value) == true {
				fmt.Println(wallet.Addresses[0].Value, wallet.Addresses[0].PrivateKey)
					if output == true {
				fmt.Fprintln(w,wallet.Addresses[0].Value, wallet.Addresses[0].PrivateKey)
				w.Flush()
				}
				a++
			}

		} else {		
			fmt.Println(wallet.Addresses[0].Value, wallet.Addresses[0].PrivateKey)
				if output == true {
			fmt.Fprintln(w,wallet.Addresses[0].Value, wallet.Addresses[0].PrivateKey)
			w.Flush()
			}
			a++
		}

		if a == numGenerate {
		os.Exit(1)
		}

		elapsed := time.Since(start)
		totalelapsed := elapsed.Seconds()

		if i%20000 == 0 && i!=0 {
		b:= int64((float64(i)/totalelapsed))
			fmt.Println("Tested:", i, " Running for:",elapsed, " Sol/s:",b)
		}

	}
}
