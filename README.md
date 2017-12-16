# BtcZgen

Offline BIP32 HD wallet and vanity address generator for BitcoinZ. Originally ported from [Hushgen.](https://github.com/TheTrunk/hushgen) 

##Pre-requisites
* Golang 1.7.3 (altought lower versions may work)
* Git

##Build
~~~~
go get -u github.com/TheTrunk/btczgen
go build github.com/TheTrunk/btczgen
go build github.com/TheTrunk/btczgen/btczretrieve
~~~~

##Update an Existing Install
~~~~
go clean github.com/TheTrunk/btczgen
go build github.com/TheTrunk/btczgen
go build github.com/TheTrunk/btczgen/btczretrieve
~~~~

##Usage
To generate a wallet:
~~~~
btczgen [-test] [-n 1] [-o]

Options
-test generate testnet addresses
-n number of addresses to generate. Defaults to 1
-o enable output to file outputbtczgen.txt
~~~~

To retrieve addresses generated from your HD wallet:
	
~~~~
<<<<<<< HEAD >>>>>>>
btczretrieve -passphrase="your desired passphrase" [-test] [-n 1] [-match="t1yourdesiredstring"] [-i] [-o]

Options
-passphrase Passphrase for the wallet is REQUIRED between 128 and 512 bits
-test generate testnet addresses	
-n number of addresses to retrieve. Defaults to 1
-match regex string to search for in the address
-i case insensitive string matching
-o enable output to file outputbtczretrieve.txt
~~~~

eg. Search case insensitive for a vanity address which starts with the string "t1jl"
~~~~
btczretrieve -passphrase="board start difference answer blossom roll powerful million rough butterfly bedroom beam" -match "t1jl" -i
~~~~

Note: The maximum number of addresses that can be searched given a wallet passphrase is restricted to 4,294,967,295 (unsigned 32 bit integer). 

To import the private key into BitcoinZ:
~~~~
./zcash-cli importprivkey "private_key_from_btczgen"
~~~~
zcashd will automatically rescan the blockchain for transactions
