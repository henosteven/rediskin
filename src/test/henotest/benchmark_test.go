package henotest

import (
    "testing"
)

func Benchmark_Add1(t *testing.B) {
   for i := 0; i < 1000000000; i++ {
       Add(1, i)
   } 
}

func Benchmark_Add2(t *testing.B) {
   for i := 0; i < 1000000; i++ {
       Add(1, i)
   } 
}
