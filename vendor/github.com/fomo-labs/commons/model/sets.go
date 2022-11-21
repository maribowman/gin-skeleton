package model

import "container/list"

type Set interface {
	Add(value interface{})
	Remove(value interface{})
	Has(value interface{}) bool
	Values() []interface{}
}

// HashSet Impl

type hashSetMap map[interface{}]struct{}

type HashSet struct {
	set hashSetMap
}

func NewHashset(values ...interface{}) *HashSet {
	hashSet := HashSet{make(hashSetMap)}
	for _, value := range values {
		hashSet.set[value] = struct{}{}
	}
	return &hashSet
}

func (set *HashSet) Add(value interface{}) {
	set.set[value] = struct{}{}
}

func (set *HashSet) Remove(value interface{}) {
	delete(set.set, value)
}

func (set *HashSet) Has(value interface{}) bool {
	_, ok := set.set[value]
	return ok
}

func (set *HashSet) Values() []interface{} {
	keys := make([]interface{}, 0, len(set.set))
	for key := range set.set {
		keys = append(keys, key)
	}
	return keys
}

// OrderedSet Impl

type keys map[interface{}]*list.Element

type OrderedSet struct {
	keys       keys
	linkedList *list.List
}

func NewOrderedSet(values ...interface{}) *OrderedSet {
	set := OrderedSet{
		make(keys),
		list.New(),
	}
	for _, value := range values {
		set.Add(value)
	}
	return &set
}

func (set *OrderedSet) Add(value interface{}) {
	_, exists := set.keys[value]
	if !exists {
		set.keys[value] = set.linkedList.PushBack(value)
	}
}

func (set *OrderedSet) Remove(value interface{}) {
	element, exists := set.keys[value]
	if exists {
		set.linkedList.Remove(element)
		delete(set.keys, value)
	}
}

func (set *OrderedSet) Has(value interface{}) bool {
	_, ok := set.keys[value]
	return ok
}

func (set *OrderedSet) Size() int {
	return len(set.keys)
}

func (set *OrderedSet) Values() []interface{} {
	values := make([]interface{}, len(set.keys))
	element := set.linkedList.Front()
	for index := 0; element != nil; index++ {
		values[index] = element.Value
		element = element.Next()
	}
	return values
}
