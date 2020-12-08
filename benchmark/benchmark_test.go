package benchmark

import (
	"fmt"
	"runtime"
	"testing"
)

func init()  {
	a := new(runtime.MemStats)
	runtime.ReadMemStats(a)
	fmt.Printf("%v",a)
}

func Sum () int {
	sum := 0
	for i := 0; i<= 10000 ;i ++ {
		sum += i
	}
	return sum

}

func BenchmarkSum(b *testing.B) {
	for i:=0 ;i <=b.N ;i++ {
		Sum()
	}
}
