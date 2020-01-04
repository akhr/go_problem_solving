package mymap

import (
	"fmt"
	"testing"
)

func TestOrderedMapUsingList_Put(t *testing.T) {
	oM := NewOrderedMapUsingList()
	oM.Put("a", 1)
	oM.Put("b", 3)
	oM.Put("c", 6)
	oM.Put("a", 5)
	keys, vals := oM.Entries()
	fmt.Printf("Keys :: %+v\n, Values :: %+v\n", keys, vals)
}
