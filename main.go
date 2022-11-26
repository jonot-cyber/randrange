package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func main() {
	mathMode := flag.String("m", "", "Example: [0, 20)")
	standardMode := flag.String("s", "", "Example: 0-21")
	flag.Parse()

	inMathMode := *mathMode != ""
	inStandardMode := *standardMode != ""

	if inMathMode && inStandardMode {
		fmt.Fprintln(os.Stderr, "You can't generate two different ranges. Pick a mode.")
		return
	}

	if inMathMode {
		generateMath(*mathMode)
	} else if inStandardMode {
		generateStandard(*standardMode)
	} else {
		fmt.Fprintln(os.Stderr, "You must input a range. See -h")
		return
	}
}

func generateStandard(s string) {
	reader := strings.NewReader(s)
	begin, err := parseInt(reader)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not parse range")
		return
	}
	parseDash(reader)
	end, err := parseInt(reader)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not parse range")
	}
	fmt.Println(begin, end)
}

func parseInt(reader io.ReadSeeker) (int, error) {
	skipWhiteSpace(reader)
	buf := make([]byte, 1)
	n, err := reader.Read(buf)
	if err != nil {
		return 0, err
	}
}

func skipWhiteSpace(reader io.ReadSeeker) error {
	buf := make([]byte, 1)
	for {
		_, err := reader.Read(buf)
		if err != nil {
			return err
		}
		if !unicode.IsSpace(rune(buf[0])) {
			break
		}
	}
	reader.Seek(-1, io.SeekCurrent)
	return nil
}

func generateMath(s string) {

}
