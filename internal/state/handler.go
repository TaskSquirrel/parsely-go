package state

type Node struct {
    value string
    previous, next *Node
}

func (n *Node) SetNext(next *Node) *Node {
    n.next = next
    next.previous = n

    return n
}

func (n *Node) SetPrevious(previous *Node) *Node {
    n.previous = previous
    previous.next = n

    return n
}

