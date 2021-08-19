/*
 Linked list
 
 Head
 [2]->[7]->[11]->[16]->[18] 

 Going to the Kth element => Slower O(n)

 Adding and removing values at the beginning of the list
 => Faster O(1)  // faster then normal list 
 "[99]"->[2]->[7]->[11]->[16]->[18]    :  add 99 at the beginning 

 Doubly linked list
 [2]<->[7]<->[11]<->[16]<->[18] 
*/

import "fmt"

type node struct{
	data int 
	next *node // Address of the next node
}


type linkedList struct{
	head *node 
	length int 
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second 
	l.length++
}

func (l linkedList) printListData() {
	toPrint := l.head
	
	for l.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func (l *linkedList) delectWithValue(value int) {
	if l.length == 0{
		return 
	}

	if l.head.data == value {
		l.head = l.head.next
		l.length--
		return 
	}

	previousToDelect := l.head 
	for previousToDelect.next.data != value {
		if previousToDelect.next.next == nil {
			return
		}
		previousToDelect = previousToDelect.next
	}
	previousToDelect.next = previousToDelect.next.next
	l.length--
}

/*
func main() {
	mylist := linkedList{}
	node1 := &node{data:58}
	node2 := &node{data:18}
	node3 := &node{data:34}
	node4 := &node{data:5}
	node5 := &node{data:99}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)

	mylist.printListData()
	mylist.delectWithValue(99)
	mylist.printListData()

	emptylist := linkedList{}
	emptylist.delectWithValue(99)
	emptylist.printListData()
}
*/