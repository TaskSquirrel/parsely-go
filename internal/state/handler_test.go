package state

import (
    "testing"
)

func TestNext(t *testing.T) {
    a := Node{value: "A", previous: nil, next: nil}
    b := Node{value: "B", previous: nil, next: nil}

    a.SetNext(&b)

    if a.next != &b || b.previous != &a {
        t.Errorf("It's wrong...")
    } else {
        t.Logf("It's right! %v, %v", a.next, b)
    }
}

func TestPrevious(t *testing.T) {
    a := Node{value: "A", previous: nil, next: nil}
    b := Node{value: "B", previous: nil, next: nil}

    a.SetPrevious(&b)

    if a.previous != &b || b.next != &a {
        t.Errorf("References are wrong...")
    } else {
        t.Logf("Testing previous worked!")
    }
}

