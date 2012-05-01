lrucache
========

Go Cache is a least recently used caching mechanism. 

Example :

package main 

import (
	"lrucache"
	"fmt"
	"runtime"
)

// End of Lib. Test Code follows

type MyInt int
var UserStore *lrucache.CacheStore

func (mI MyInt) Save () {
	fmt.Println ("Saving ", mI)
}

func main () {
var value MyInt
UserStore = lrucache.NewCacheStore (10)
for i := 0; i <= 1000; i++ {
	
        value     = MyInt(i)
	valueElem := lrucache.NewValueElem (value, UserStore)
	fmt.Println ("Adding", valueElem.Value) 
	UserStore.Add (fmt.Sprintf ("%d", i), *valueElem)
	}
//delay 

for i := 0 ; i < 1000000000; i++ {
	runtime.Gosched ()
	}

}


Details Later !!!
