package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	inputs := strings.Split(s.Text(), " ")
}
