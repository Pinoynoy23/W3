package crack

import (
	"bufio"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha512"
	"errors"
	"fmt"
	"hash"
	"os"
	"strings"
)

// computeHexHash converts hash.Hash to lowercase hex string
func computeHexHash(h hash.Hash) string {
	return fmt.Sprintf("%x", h.Sum(nil))
}

// CrackWithWordlist reads the wordlist and tries each candidate.
// - algo: "md5", "sha1", "sha512"
// - targetHash: expected hex string (lower/upper OK)
// - wordlistPath: path to file
// - verbose: if true, prints some progress lines
// - progressEvery: print progress every N lines if >0
func CrackWithWordlist(algo, targetHash, wordlistPath string, verbose bool, progressEvery int) (string, error) {
	target := strings.ToLower(strings.TrimSpace(targetHash))

	f, err := os.Open(wordlistPath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	// increase buffer if you expect very long lines (rare for password lists)
	const maxCapacity = 1024 * 1024 // 1MB
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, maxCapacity)

	lineNo := 0
	for scanner.Scan() {
		lineNo++
		candidate := strings.TrimSpace(scanner.Text())
		if candidate == "" {
			continue
		}

		var h hash.Hash
		switch strings.ToLower(algo) {
		case "md5":
			h = md5.New()
		case "sha1":
			h = sha1.New()
		case "sha512":
			h = sha512.New()
		default:
			return "", errors.New("unsupported algorithm: " + algo)
		}
		h.Write([]byte(candidate))
		computed := computeHexHash(h)

		if computed == target {
			if verbose {
				fmt.Printf("[+] Line %d: matched candidate=%s\n", lineNo, candidate)
			}
			return candidate, nil
		}

		if verbose && progressEvery > 0 && (lineNo%progressEvery == 0) {
			fmt.Printf("[*] scanned %d lines, last candidate=%s\n", lineNo, candidate)
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", errors.New("not found")
}
