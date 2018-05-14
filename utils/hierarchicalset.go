package utils

// HierarchicalSet : object
type HierarchicalSet struct {
	structure map[string](*Set)
}

// CreateHierarchicalSet : construct
func CreateHierarchicalSet() *HierarchicalSet {
	return &HierarchicalSet{structure: make(map[string](*Set))}
}

// AddHierarchicalSet : add another set
func (hierarchicalSet *HierarchicalSet) AddHierarchicalSet(secondHierarchicalSet *HierarchicalSet) {

	for k, v := range hierarchicalSet.structure {

		secondSet, found := secondHierarchicalSet.structure[k]

		if found {
			v.AddSet(secondSet)
		}
	}
}

// Add : add item to set if it does not already exist
func (hierarchicalSet *HierarchicalSet) Add(key string, item interface{}) {

	set, found := hierarchicalSet.structure[key]

	if found {
		set.Add(item)
	} else {

		set = CreateSet()
		set.Add(item)
		hierarchicalSet.structure[key] = set
	}
}

// GetKeys : return all the top level keys
func (hierarchicalSet *HierarchicalSet) GetKeys() []string {

	keys := make([]string, 0, 100)

	for k := range hierarchicalSet.structure {
		keys = append(keys, k)
	}

	return keys
}

// GetSecondLevelSet : return a Set given a key
func (hierarchicalSet *HierarchicalSet) GetSecondLevelSet(key string) *Set {
	value, found := hierarchicalSet.structure[key]

	if found {
		return value
	} else {
		return nil
	}
}

// VisitAll : visit every part of the tree and call the visitFunc on node
func (hierarchicalSet *HierarchicalSet) VisitAll(visitFunc Visitor) {
	for k := range hierarchicalSet.structure {
		set := hierarchicalSet.GetSecondLevelSet(k)

		set.VisitAll(k, visitFunc)
	}
}
