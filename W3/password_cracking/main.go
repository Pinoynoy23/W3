package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/yourname/password_cracker/utils/crack"
)

func main() {
	algo := flag.String("algo", "md5", "algorithm: md5 | sha1 | sha512")
	target := flag.String("target", "", "target hash to crack (hex)")
	wordlist := flag.String("wordlist", "nord_vpn.txt", "path to wordlist file")
	verbose := flag.Bool("verbose", false, "verbose output")
	progressEvery := flag.Int("progressEvery", 0, "print progress every N lines (0=off)")
	flag.Parse()

	if *target == "" {
		log.Fatal("target hash required: -target <hash>")
	}

	fmt.Printf("Trying to crack target=%s using algo=%s wordlist=%s\n", *target, *algo, *wordlist)
	found, err := crack.CrackWithWordlist(*algo, *target, *wordlist, *verbose, *progressEvery)
	if err != nil {
		fmt.Printf("Result: NOT FOUND (%v)\n", err)
		return
	}
	fmt.Printf("Result: FOUND! password = %s\n", found)
}
