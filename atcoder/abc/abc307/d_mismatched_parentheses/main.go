package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	fmt.Scan(&n)

	sc := bufio.NewScanner(os.Stdin)
	sc.Buffer(make([]byte, 64*1024), 200001)
	sc.Scan()
	s := sc.Text()

	sindexes := []int{}
	ans := ""
	for i := 0; i < n; i++ {
		si := string(s[i])
		ans += si
		if si == "(" {
			sindexes = append(sindexes, len(ans)-1)
		} else if si == ")" && len(sindexes) > 0 {
			last := sindexes[len(sindexes)-1]
			ans = ans[:last]
			sindexes = sindexes[:len(sindexes)-1]
		}
	}
	fmt.Println(ans)
}
