package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"regexp"
	"strconv"
	"time"
)

func main() {
	mathMode := flag.String("m", "", "Example: [0, 20)")
	standardMode := flag.String("s", "", "Example: 0-21")
	flag.Parse()

	inMathMode := *mathMode != ""
	inStandardMode := *standardMode != ""

	if inMathMode && inStandardMode {
		_, _ = fmt.Fprintln(os.Stderr, "You can't generate two different ranges. Pick a mode.")
		return
	}

	if inMathMode {
		generateMath(*mathMode)
	} else if inStandardMode {
		generateStandard(*standardMode)
	} else {
		_, _ = fmt.Fprintln(os.Stderr, "You must input a range. See -h")
	}
}

// generateStandard takes a string s in the form begin-end and generates a random number in the range [begin, end]
func generateStandard(s string) {
	r := regexp.MustCompile(`^\s*(\d+)\s*-\s*(\d+)\s*$`)
	match := r.FindStringSubmatch(s)
	if len(match) != 3 {
		_, _ = fmt.Fprintln(os.Stderr, "Could not parse the string. Check -h for an example")
	}
	// Bad practice, but I think the regex handles this.
	begin, _ := strconv.Atoi(match[1])
	end, _ := strconv.Atoi(match[2])
	generateRange(begin, end+1) // Add +1 because inclusive range
}

// generateRange generates a random number within the range [begin, end)
func generateRange(begin int, end int) {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(end-begin) + begin
	fmt.Println(num)
}

// generateMath takes a string s in interval notation and generates a random number in that range
func generateMath(s string) {
	r := regexp.MustCompile(`^\s*([\[\(])\s*(\d+)\s*,\s*(\d+)([\]\)])\s*$`)
	match := r.FindStringSubmatch(s)
	if len(match) != 5 {
		_, _ = fmt.Fprintln(os.Stderr, "Could not parse the string. Check -h for an example")
	}
	begin, _ := strconv.Atoi(match[2])
	if match[1] == "(" { // If the range is open, don't include the starting value
		begin++
	}
	end, _ := strconv.Atoi(match[3])
	if match[4] == "]" {
		end++
	}
	generateRange(begin, end)
}
