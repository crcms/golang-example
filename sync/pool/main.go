package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main()  {
	pool := sync.Pool{
		New: func() interface{} {
			return make([]byte,10)
		},
	}


	var bytes = []byte(`abc`)
	pool.Put(bytes)
	runtime.GC()
	pool.Put([]byte("sss"))
	fmt.Printf("%s\n",pool.Get())
	fmt.Printf("%s\n",pool.Get())
}
