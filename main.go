package main

import (
	"math/rand"
	"fmt"
	"time"
	"strconv"
	"flag"
	"unicode/utf8"
)

const maxUnicode = 137928 // http://www.babelstone.co.uk/Unicode/HowMany.html
const pauseDuration = 250 // Amount of time to pause after finding a match

func randChar() string {
		randRune := rand.Intn(maxUnicode) 
		return string(randRune)
}

func main() {
	start := time.Now() // Start duration timer
	rand.Seed(time.Now().UnixNano()) // Initialize the random number generator
	stringPtr := flag.String("w", "hello world", "A string to bruteforce") // Setup command line flags
	flag.Parse()
	targetString := *stringPtr
	buf := []byte(targetString) 
	// fmt.Println("bytes =", len(buf))
	// fmt.Println("runes =", utf8.RuneCount(buf))
	targetRunes := []rune(targetString) // Get runes from string

	fmt.Println("Target String: " + targetString + " (Length: " + strconv.Itoa(utf8.RuneCount(buf)) + ") \n")
	
	count := 0
	for count < utf8.RuneCount(buf) {
		char := randChar()
		pos := 0
		currentBuf := make([]byte, utf8.RuneCount(buf))
		
		for pos < utf8.RuneCount(buf) {
			fmt.Println("\033[2;0HTrying character: " + char)
			
			if char == string(targetRunes[pos]) {
				fmt.Println("\033[3;0H" + "Got a match!!            ")
				time.Sleep(pauseDuration * time.Millisecond) // Pause when match found so you can watch it work
				currentBuf = append(currentBuf[:], char...)
				pos ++
				count ++
				char = randChar() // Get a new random character for next time
			} else {
				fmt.Println("\033[3;0H" + "No match, looking for: " + string(targetRunes[pos]))
				char = randChar() // Get a new random character for next time
			}
			fmt.Println("\033[5;0H" + string(currentBuf))
		}
	}
	
	pausedTime := utf8.RuneCount(buf) * pauseDuration
	elapsed := (time.Since(start) - (time.Duration(pausedTime) * time.Millisecond)) // Subtract time of match pauses
	
	fmt.Printf("\nBruteforcing completed in %s.\n", elapsed)
}

