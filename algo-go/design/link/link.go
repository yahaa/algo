package link

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return nil
	}

	p := root

	for p != nil {
		pnext := p.Next
		if p.Child != nil {
			sub := flatten(p.Child)

			last := sub
			for last.Next != nil {
				last = last.Next
			}

			p.Next = sub
			sub.Prev = p

			last.Next = pnext
			if pnext != nil {
				pnext.Prev = last
			}

			p.Child = nil
		}

		p = pnext
	}

	return root
}
