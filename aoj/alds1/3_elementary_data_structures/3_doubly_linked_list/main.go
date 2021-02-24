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
		ll.first = &val
		ll.last = &val
	} else {
		val.next = ll.first
		ll.first.prev = &val
		ll.first = &val
	}
}

func (ll *linkedList) delete(x int) {
	val := ll.first
	for val != nil {
		if val.v != x {
			val = val.next
			continue
		}

		if val.prev == nil {
			ll.first = val.next
		} else {
			val.prev.next = val.next
		}
		if val.next == nil {
			ll.last = val.prev
		} else {
			val.next.prev = val.prev
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

func (ll *linkedList) print() {
	val := ll.first
	if val == nil {
		return
	}
	fmt.Print(val.v)
	val = val.next
	for val != nil {
		fmt.Printf(" %d", val.v)
		val = val.next
	}
	fmt.Println()
}

func main() {
	var n int
	fmt.Scan(&n)
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)

	ll := linkedList{}
	for i := 0; i < n; i++ {
		s.Scan()
		com := s.Text()
		if com == "insert" {
			s.Scan()
			x, _ := strconv.Atoi(s.Text())
			ll.insert(x)
		} else if com == "delete" {
			s.Scan()
			x, _ := strconv.Atoi(s.Text())
			ll.delete(x)
		} else if com == "deleteFirst" {
			ll.deleteFirst()
		} else if com == "deleteLast" {
			ll.deleteLast()
		}
	}
	ll.print()
}
