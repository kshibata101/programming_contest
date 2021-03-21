package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)
	h := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &h[i])
	}

	for i := 0; i < n; i++ {
		fmt.Fprintf(w, "node %d: key = %d, ", i+1, h[i])
		if i > 0 {
			fmt.Fprintf(w, "parent key = %d, ", h[(i-1)/2])
		}
		if 2*i+1 < n {
			fmt.Fprintf(w, "left key = %d, ", h[2*i+1])
		}
		if 2*i+2 < n {
			fmt.Fprintf(w, "right key = %d, ", h[2*i+2])
		}
		fmt.Fprintln(w)
	}
}
