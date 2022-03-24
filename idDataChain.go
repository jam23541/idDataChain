package idDataChain

type IdDataChain struct {
	arrayOfData [32]IdDataNode
	p_head      *IdDataNode
	p_tail      *IdDataNode
	mapOfData   map[string]*IdDataNode
}

func (myChain *IdDataChain) Init() {
	myChain.p_head = &myChain.arrayOfData[0]
	myChain.p_tail = &myChain.arrayOfData[0]
	myChain.arrayOfData[0].p_next = &myChain.arrayOfData[1]
	myChain.arrayOfData[cap(myChain.arrayOfData)-1].p_next = nil
	myChain.arrayOfData[0].p_pre = nil
	myChain.arrayOfData[cap(myChain.arrayOfData)-1].p_pre = &myChain.arrayOfData[cap(myChain.arrayOfData)-2]
	for i := 1; i < cap(myChain.arrayOfData)-1; i++ {
		myChain.arrayOfData[i].p_pre = &myChain.arrayOfData[i-1]
		myChain.arrayOfData[i].p_next = &myChain.arrayOfData[i+1]
	}
	myChain.mapOfData = make(map[string]*IdDataNode)
}
func NewIdDataChain() *IdDataChain {
	return &IdDataChain{
		arrayOfData: [32]IdDataNode{},
		p_head:      nil,
		p_tail:      nil,
		mapOfData:   nil,
	}
}

// put: return -1 if data already exists
// 1 if the data is eventually added to the chain
func (myChain *IdDataChain) Put(inputId string) int {

	_, exist := myChain.mapOfData[inputId]
	if exist {
		return -1
	} else {
		// if is full, remove the first node
		if myChain.p_tail.p_next == nil {
			myChain.Delete(myChain.p_head.myData)
		}
		// put data in
		myChain.p_tail.myData = inputId
		myChain.p_tail = myChain.p_tail.p_next
		myChain.mapOfData[inputId] = myChain.p_tail.p_pre
		return 1
	}
}

// delete
func (myChain *IdDataChain) Delete(inputId string) int {
	p_node, exist := myChain.mapOfData[inputId]
	if exist {
		// make next node to have same connection as current node

		if p_node.p_pre != nil {
			p_node.p_pre.p_next = p_node.p_next
		} else {
			// if it is the first element
			myChain.p_head = p_node.p_next
		}
		if p_node.p_next != nil {
			p_node.p_next.p_pre = p_node.p_pre
		}
		// clear content
		p_node.myData = ""

		// move to be tail.nextNode
		p_node.p_pre = myChain.p_tail
		p_node.p_next = myChain.p_tail.p_next
		myChain.p_tail.p_next = p_node
		// if original tail.next exists
		if p_node.p_next != nil {
			p_node.p_next.p_pre = p_node
		}
		delete(myChain.mapOfData, inputId)
		return 1
	} else {
		return -1
	}
}

// check
func (myChain *IdDataChain) Check(inputId string) int {

	_, exist := myChain.mapOfData[inputId]
	if exist {
		return 1
	} else {
		return -1
	}
}

// clear all the data
func (myChain *IdDataChain) Reset() {
	myChain.arrayOfData = [32]IdDataNode{}
	myChain.mapOfData = nil
	myChain.Init()
}
