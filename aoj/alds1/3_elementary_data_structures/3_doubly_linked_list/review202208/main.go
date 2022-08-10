package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type value struct {
	v    int
	next *value
	prev *value
}

type linkedList struct {
	first *value
	last  *value
}

func (ll *linkedList) insert(x int) {
	val := value{v: x}
	if ll.first == nil {
		ll.last = &val
	} else {
		val.next = ll.first
		val.next.prev = &val
	}
	ll.first = &val
}

func (ll *linkedList) delete(x int) {
	val := ll.first
	for val != nil {
		if val.v != x {
			val = val.next
			continue
		}

		if val.next == nil {
			ll.last = val.prev
		} else {
			val.next.prev = val.prev
		}
		if val.prev == nil {
			ll.first = val.next
		} else {
			val.prev.next = val.next
		}
		break
	}
}

func (ll *linkedList) deleteFirst() {
	if ll.first.next == nil {
		ll.last = nil
	} else {
		ll.first.next.prev = nil
	}
	ll.first = ll.first.next
}

func (ll *linkedList) deleteLast() {
	if ll.last.prev == nil {
		ll.first = nil
	} else {
		ll.last.prev.next = nil
	}
	ll.last = ll.last.prev
}

func main() {
	var n int
	fmt.Scan(&n)

	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	ll := linkedList{}
	var x int
	for i := 0; i < n; i++ {
		s.Scan()
		switch s.Text() {
		case "insert":
			s.Scan()
			x, _ = strconv.Atoi(s.Text())
			ll.insert(x)
		case "delete":
			s.Scan()
			x, _ = strconv.Atoi(s.Text())
			ll.delete(x)
		case "deleteFirst":
			ll.deleteFirst()
		case "deleteLast":
			ll.deleteLast()
		}
	}

	val := ll.first
	fmt.Print(val.v)
	for val.next != nil {
		val = val.next
		fmt.Printf(" %d", val.v)
	}
	fmt.Println()
}
