package mymap

import (
	"container/list"
)

type OrderedMapUsingList struct {
	m map[string]int
	l *list.List
}

func NewOrderedMapUsingList() OrderedMapUsingList {
	return OrderedMapUsingList{
		m: map[string]int{},
		l: list.New(),
	}
}

func (o *OrderedMapUsingList) Put(k string, v int) {
	o.m[k] = v
	o.l.PushBack(k)
}

func (o *OrderedMapUsingList) Get(k string) int {
	return o.m[k]
}

func (o *OrderedMapUsingList) Delete(k string) {
	if _, ok := o.m[k]; ok {
		delete(o.m, k)
		o.l.Init()
	}
}

func (o *OrderedMapUsingList) Entries() ([]string, []int) {
	keys := []string{}
	vals := []int{}
	for e := o.l.Front(); e != nil; e = e.Next() {
		keys = append(keys, e.Value.(string))
		vals = append(vals, o.Get(e.Value.(string)))
	}
	return keys, vals
}
