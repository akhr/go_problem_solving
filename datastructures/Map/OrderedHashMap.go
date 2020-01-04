package mymap

type OrderedMap interface {
	Put(key, val interface{})
	Get(key interface{}) interface{}
	Entries() [][]interface{}
	Delete(key interface{})
}
