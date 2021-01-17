package main

import (
    "fmt"
    "sort"
)

type Store struct {
    Name string
    Point int
    Id int
}

type Stores []Store

func (s Stores) Len() int { return len(s) }
func (s Stores) Swap(i, j int) { s[i], s[j] = s[j], s[i] }
func (s Stores) Less(i, j int) bool {
    if (s[i].Name < s[j].Name) {
        return true
    } else if (s[i].Name > s[j].Name) {
        return false
    } else {
        return s[i].Point > s[j].Point
    }
}

func main() {
    var n int
    fmt.Scan(&n)
    stores := make(Stores, n)
    for i := 0; i < n; i++ {
        var s string
        var p int
        fmt.Scan(&s)
        fmt.Scan(&p)
        stores[i] = Store{s, p, i+1}
    }
    sort.Sort(stores)

    for i := 0; i < n; i++ {
        fmt.Println(stores[i].Id)
    }
}
