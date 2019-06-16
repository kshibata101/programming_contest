package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

const B uint64 = 100000007

func main() {
	var t, p string
	fmt.Scan(&t)
	fmt.Scan(&p)

	ts, ps := []rune(t), []rune(p)
	n, m := len(ts), len(ps)
	if n < m {
		return
	}

	var th, ph, bm uint64 = 0, 0, 1
	for i := 0; i < m; i++ {
		th = uint64(ts[i]) + th*B
		ph = uint64(ps[i]) + ph*B
		bm = bm*B
	}
	wr := bufio.NewWriter(os.Stdout)
	if th == ph {
		wr.WriteString("0\n")
	}
	for i := 1; i <= n-m; i++ {
		th = th*B + uint64(ts[i+m-1]) - uint64(ts[i-1])*bm
		if th == ph {
			wr.WriteString(strconv.Itoa(i) + "\n")
		}
	}
	wr.Flush()
}
