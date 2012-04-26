package lrucache 

import (
	"sync"
)

// All Cache variables must have a save method

type CacheValue interface {
	Save () 
}

type ValueElem struct {

	Priority int
	Value CacheValue
	Mu sync.RWMutex	
}

func NewValueElem (value CacheValue, store *CacheStore) (elem *ValueElem) {
	elem = &ValueElem {
		Priority : store.GetTopPriority (),
		Value    : value,
	} 
	return
}

type CacheStore struct {
	Store map[string]ValueElem
	TopPriority int
        MaxLength int
	StatusChannel chan bool
	Mu sync.RWMutex
}

func NewCacheStore (maxLength int) (store *CacheStore) {
	store = &CacheStore { 
		Store       : make (map[string]ValueElem),
		StatusChannel : make (chan bool, 2),
		TopPriority : 0,
		MaxLength   : maxLength,	
	}
	go store.Cleaner ()
	return
}

func (c* CacheStore) Cleaner () {
	for {
		<- c.StatusChannel
		// Iterate through decrease priority
		c.Mu.Lock()
		c.TopPriority = c.MaxLength/2
		for key, value := range c.Store {
			if (c.process (&value)) {
				value.Value.Save ()
				delete (c.Store, key)		
			} else {
				c.Store[key] = value
			}	
			
		}
		c.Mu.Unlock ()
	}
}

func (c* CacheStore) process (value *ValueElem) (status bool) {
	value.Priority -= c.TopPriority
	if (value.Priority < 0) {
		status = true
	} else {
		status = false
	}
	return
}

func (c* CacheStore) Add (key string, elem ValueElem) {
	elem.Mu.Lock ()
	c.Store [key] = elem
	elem.Mu.Unlock ()
}

func (c* CacheStore) GetValueElem (key string) (present bool, value ValueElem) {
	value, present = c.Store [key]
	return 
} 

func (c* CacheStore) GetValue (key string) (present bool, value CacheValue) {
	valueElem, present := c.Store [key]
	if present {
		value = valueElem.Value
	}
	return 
} 

func (c* CacheStore) GetTopPriority() (priority int) {
	c.Mu.Lock ()
	c.TopPriority ++
	priority = c.TopPriority
	c.Mu.Unlock ()
	if priority == c.MaxLength - 1 {
		c.StatusChannel <- true
	}
	return
}
