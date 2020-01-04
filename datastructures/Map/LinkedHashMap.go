package mymap

type Node struct {
	key   interface{}
	value interface{}
	prev  *Node
	next  *Node
}

type LinkedHashMap struct {
	dataMap map[interface{}]interface{}
	head    *Node
	tail    *Node
}

func (l *LinkedHashMap) Put(key, val interface{}) {
	if _, ok := l.dataMap[key]; !ok {
		n := &Node{
			key:   key,
			value: val,
			prev:  l.tail,
		}
		l.dataMap[key] = n
		l.tail.next = n
		l.tail = n
	} else {
		n := l.dataMap[key].(*Node)
		n.value = val
	}
}
