package set

// Set : object
type Set struct {
	slice []interface{}
}

// Create : construct
func Create(protoypeSlice []interface{}) *Set {
	return &Set{slice: protoypeSlice}
}

// AddSet : add another set
func (set *Set) AddSet(secondSet Set) {

	for i := range secondSet.slice {
		set.Add(secondSet.slice[i])
	}
}

// Add : add item to set if it does not already exist
func (set *Set) Add(item interface{}) {

	if set.Contains(item) == false {
		set.slice = append(set.slice, item)
	}
}

// Contains : checks if this set contains the item
func (set *Set) Contains(item interface{}) bool {

	for i := range set.slice {
		if set.slice[i] == item {
			return true
		}
	}

	return false
}
