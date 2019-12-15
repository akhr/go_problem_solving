package set

// Set : Set interface
type Set interface {
	Add(key interface{}) bool
	Contains(key interface{}) bool
	Delete(key interface{}) bool
}
