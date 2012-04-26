package lrucache

import (
	"testing"
	"fmt"
)

type MyInt int
var Store *CacheStore

func (mI MyInt) Save () {
	fmt.Println ("Saving ", mI)
}

func TestCache (t *testing.T) {
var value MyInt
Store := NewCacheStore (10)
for i := 0; i <= 1000 ; i++ {
	
        value    = MyInt(i)
	valueElem := NewValueElem (value, Store)
	Store.Add (string (i), *valueElem)
	t.Log ("Is this ?")
	}
}
