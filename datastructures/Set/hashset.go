package set

import (
	"fmt"
)

type setImpl struct {
	dataStore map[interface{}]bool
}

// NewSet :
func NewSet() Set {
	dataStore := make(map[interface{}]bool)
	impl := &setImpl{
		dataStore: dataStore,
	}
	return impl
}

func (s *setImpl) Add(key interface{}) bool {
	if s.dataStore[key] {
		fmt.Printf("Set already contains the key - %s\n", key)
		return false
	}
	s.dataStore[key] = true
	return true
}

func (s *setImpl) Contains(key interface{}) bool {
	if !s.dataStore[key] {
		fmt.Printf("Key not available in the Set - %s\n", key)
		return false
	}
	return true
}

func (s *setImpl) Delete(key interface{}) bool {
	if !s.dataStore[key] {
		fmt.Printf("Key not available in the Set - %s\n", key)
		return false
	}
	delete(s.dataStore, key)
	return true
}
