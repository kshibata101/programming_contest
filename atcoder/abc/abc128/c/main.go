package main

import (
    "fmt"
)

func main() {
    var n, m int
    fmt.Scan(&n)
    fmt.Scan(&m)
    s := make([][]int, m)
    for i := 0; i < m; i++ {
       var k int
       fmt.Scan(&k)
       s[i] = make([]int, k)
       for j := 0; j < k; j++ {
           fmt.Scan(&s[i][j])
       }
    }
    p := make([]int, m)
    for i := 0; i < m; i++ {
        fmt.Scan(&p[i])
    }
    total := 0
    sw := make([]int, n)
    for i := 0; i < (1 << uint(n)); i++ {
        for j := 0; j < n; j++ {
            sw[j] = (i >> uint(j)) & 1
        }
        result := true
        for j := 0; j < m; j++ {
            count := 0
            for k := 0; k < len(s[j]); k++ {
                if sw[s[j][k]-1] == 1 {
                    count += 1
                }
            }
            if count % 2 != p[j] {
                result = false
                break
            }
        }
        if result {
            total += 1
        }
    }
    fmt.Println(total)
}
