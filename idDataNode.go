package idDataChain

type IdDataNode struct {
	p_pre  *IdDataNode
	p_next *IdDataNode
	myData string
}

func NewIdDataNode() *IdDataNode {
	return &IdDataNode{
		p_pre:  nil,
		p_next: nil,
		myData: "",
	}
}
