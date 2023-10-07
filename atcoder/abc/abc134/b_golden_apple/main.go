package main

import (
  "fmt"
  "math"
)

func main() {
    var n, d int
    fmt.Scan(&n)
    fmt.Scan(&d)
    fmt.Println(math.Ceil(float64(n) / float64(2*d+1)))
}
