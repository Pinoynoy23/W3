package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
	"strings"
	"golang.org/x/crypto/sha3"
)

func hashMD5(s string) string {
	h := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", h)
}

func hashSHA1(s string) string {
	h := sha1.Sum([]byte(s))
	return fmt.Sprintf("%x", h)
}

func hashSHA256(s string) string {
	h := sha256.Sum256([]byte(s))
	return fmt.Sprintf("%x", h)
}

func hashSHA512(s string) string {
	h := sha512.Sum512([]byte(s))
	return fmt.Sprintf("%x", h)
}

func hashSHA3_256(s string) string {
	h := sha3.Sum256([]byte(s))
	return fmt.Sprintf("%x", h)
}

func compareAndPrint(name, a, b string) {
	match := "No Match!"
	if a == b {
		match = "Match!"
	}
	fmt.Printf("Hash (%s):\nOutput A = %s\nOutput B = %s\n=> %s\n\n", name, a, b, match)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("========Name + Hashing Program========")
	fmt.Print("Please input value 1: ")
	txt1Raw, _ := reader.ReadString('\n')
	txt1 := strings.TrimSpace(txt1Raw)

	fmt.Print("Please input value 2: ")
	txt2Raw, _ := reader.ReadString('\n')
	txt2 := strings.TrimSpace(txt2Raw)

	outA := map[string]string{
		"MD5":     hashMD5(txt1),
		"SHA1":    hashSHA1(txt1),
		"SHA256":  hashSHA256(txt1),
		"SHA512":  hashSHA512(txt1),
		"SHA3-256": hashSHA3_256(txt1),
	}
	outB := map[string]string{
		"MD5":     hashMD5(txt2),
		"SHA1":    hashSHA1(txt2),
		"SHA256":  hashSHA256(txt2),
		"SHA512":  hashSHA512(txt2),
		"SHA3-256": hashSHA3_256(txt2),
	}

	for _, algo := range []string{"MD5", "SHA1", "SHA256", "SHA512", "SHA3-256"} {
		compareAndPrint(algo, outA[algo], outB[algo])
	}
}
