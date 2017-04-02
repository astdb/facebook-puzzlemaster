// - n number of programmers (ID'd from 1-n) are working on decoding an n number of wiretaps
//   (given in input file by victim's first names) - find the optimal matching of victims to
//   programmers so the total time needed to decode is minimized.

// - exactly the same number of programmers as wiretap victims
// - each programmer decodes exactly one wiretap
// - to decode a wiretap it takes at least 1 server hour per letter in the victim's name
// - programmers work one at a time
// - if programmer ID is even, an extra 1.5hours of work is required per vowel in victim name
// - if programmer ID is odd, an extra 1.5hours of work is required per consonant in victim name
// - if the prime factors of a programmer's ID is shared with the number of letters in the victim's name,
//   additional 2 hours are required per shared prime factor

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// ensure input filename entered on command line
	if len(os.Args) <= 1 {
		fmt.Println("Usage: > go run IllegalWiretaps.go <input.file>")
		return
	}

	// capture input file name
	filename := os.Args[1]

	// number of wiretaps
	fmt.Printf("%d\n", getWireTapCount(filename))

	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
		return
	}

	scanner := bufio.NewScanner(file)
	// k := 1 // line numbers (programmer IDs in a later incarnation)
	for scanner.Scan() {
		// read next line
		line := scanner.Text()
		fmt.Println(strings.TrimSpace(line))
	}
}

// given a victim name and a programmer ID, return calculation cost in hours for decoding
func wireTapDecodeCost(name string, programmerId int) float64 {
	decodeCost := 0.0
	for _, ch := range name {
		// adding the standard decoding cost per letter - Vowelitosis and Consonentia costs are mentioned as 'in addition'
		decodeCost += 1.0

		if programmerId%2 == 0 {
			// Vowelitosis - extra 1.5hr per vowel
			if isVowel(string(ch)) {
				decodeCost += 1.5
			}
		} else {
			// Consonentia - extra 1hr per consonant
			if isVowel(string(ch)) {
				decodeCost += 1.0
			}
		}

		// TODO: cost due to severe phobia where there are shared prime factors between programmerId and len(name)
		
	}
}

func getWireTapCount(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input file: %v\n", err)
		return -1
	}
	defer file.Close()

	// determine number of wiretaps
	numTaps := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// read next line
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			numTaps++
		}
	}

	return numTaps
}

func isVowel(letter string) boolean {
	letter = strings.TrimSpace(strings.ToLower(letter))
	if len(letter) <= 0 {
		return false
	}

	vowels := []str{"a", "e", "i", "o", "u"}
	for _, vow := range vowels {
		if vow != letter {
			return false
		}
	}

	return true
}
