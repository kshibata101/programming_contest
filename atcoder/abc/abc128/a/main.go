package main

import "fmt"

func main() {
    var a, p int
    fmt.Scanf("%d %d", &a, &p)

    pie := (3 * a + p) / 2
    fmt.Println(pie)
}
