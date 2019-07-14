package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"time"
	"unicode/utf8"

	"github.com/faiface/pixel"
	"github.com/faiface/pixel/pixelgl"
	"github.com/faiface/pixel/text"
	"golang.org/x/image/colornames"
	"golang.org/x/image/font/basicfont"
)

const maxUnicode = 137928 // http://www.babelstone.co.uk/Unicode/HowMany.html
const pauseDuration = 250 // Amount of time to pause after finding a match

var header = "Header line"
var line2 = "line2"
var line3 = "line3"
var line5 = "line5"

var completed = false

func main() {
	go bruteforce()
	pixelgl.Run(run)
}

func bruteforce() {
	start := time.Now()                                                    // Start duration timer
	stringPtr := flag.String("w", "hello world", "A string to bruteforce") // Setup command line flags
	benchmarkPtr := flag.Bool("b", false, "Enables benchmark mode if set (Use same random number generator each time)")
	flag.Parse()
	targetString := *stringPtr
	benchmarkMode := *benchmarkPtr
	buf := []byte(targetString)
	// fmt.Println("bytes =", len(buf))
	// fmt.Println("runes =", utf8.RuneCount(buf))
	targetRunes := []rune(targetString) // Get runes from string
	guesses := uint64(0)                // Count of guesses/attempts

	// Initialize the random number generator
	if benchmarkMode {
		rand.Seed(1) // Use the same RNG each time if in benchmark mode
	} else {
		rand.Seed(time.Now().UnixNano()) // Use a random seed
	}

	// fmt.Println("Target String: " + targetString + " (Length: " + strconv.Itoa(utf8.RuneCount(buf)) + ")\n")
	header = "Target String: " + targetString + " (Length: " + strconv.Itoa(utf8.RuneCount(buf)) + ")"

	count := 0
	for count < utf8.RuneCount(buf) {
		char := randChar()
		pos := 0
		currentBuf := make([]byte, utf8.RuneCount(buf))

		for pos < utf8.RuneCount(buf) {
			// fmt.Println("\033[2;0HTrying character: " + char)
			line2 = "Trying character: " + char

			if char == string(targetRunes[pos]) {
				// fmt.Println("\033[3;0H" + "Got a match!!            ")
				line3 = "Found a match!!"
				time.Sleep(pauseDuration * time.Millisecond) // Pause when match found so you can watch it work
				currentBuf = append(currentBuf[:], char...)
				pos++
				count++
			} else {
				// fmt.Println("\033[3;0H" + "No match, looking for: " + string(targetRunes[pos]))
				line3 = "No match, looking for: " + string(targetRunes[pos])
			}
			char = randChar() // Get a new random character for next time
			guesses++
			// fmt.Println("\033[5;0H" + string(currentBuf))
			line5 = string(currentBuf)
		}
	}

	pausedTime := utf8.RuneCount(buf) * pauseDuration
	elapsed := (time.Since(start) - (time.Duration(pausedTime) * time.Millisecond)) // Subtract time of match pauses

	fmt.Printf("\nBruteforcing completed in %s \nTotal combinations guessed: %d \nBenchmark mode: %s\n", elapsed, guesses, strconv.FormatBool(benchmarkMode))
	completed = true
}

func randChar() string {
	randRune := rand.Intn(maxUnicode)
	return string(randRune)
}

// Pixel stuff
func run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Bruteforce helloworld",
		Bounds: pixel.R(0, 0, 1080, 720),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	basicAtlas := text.NewAtlas(basicfont.Face7x13, text.ASCII)
	basicTxt := text.New(pixel.V(100, 600), basicAtlas)

	PrintMe := RandStringRunes(10)

	// fmt.Fprintf(basicTxt, "And I'm an %s, yay!", "io.Writer")

	for !win.Closed() { // https://github.com/faiface/pixel/wiki/Typing-text-on-the-screen
		win.Clear(colornames.Black)
		basicTxt.Clear()
		fmt.Fprintln(basicTxt, header)
		fmt.Fprintln(basicTxt, line2)
		fmt.Fprintln(basicTxt, line3)
		fmt.Fprintln(basicTxt, " ") // line4 blank for readability
		// Make a Regex to say we only want letters and numbers
		reg, err := regexp.Compile("[^a-zA-Z0-9 ] *")
		if err != nil {
			log.Fatal(err)
		}
		line5 = reg.ReplaceAllString(line5, "")

		fmt.Fprintln(basicTxt, line5)
		fmt.Fprintln(basicTxt, " ")

		if completed == false {
			for i := 0; i < 5; i++ { // gibberish for no reason
				PrintMe = RandStringRunes(50)
				fmt.Fprintln(basicTxt, PrintMe)
			}
		} else {
			fmt.Fprintln(basicTxt, "Bruteforcing complete!")
		}

		basicTxt.Draw(win, pixel.IM.Scaled(basicTxt.Orig, 2))

		win.Update()
	}
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
