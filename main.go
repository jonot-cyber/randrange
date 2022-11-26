package main

import (
	"flag"
)

func main() {
	mathMode := flag.String("m", "", "Example: [0, 20)")
	standardMode := flag.String("s", "", "Example: 0-21")
	flag.Parse()
}
