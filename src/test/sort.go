package main

import (
    "sort"
    "fmt"
)

type usergrade []int

/*
type Interface interface {
    Len() int
    Less(i, j int) bool
    Swap(i, j int)
}
*/

func (ug usergrade) Len() int {
    return len(ug)
}

func (ug usergrade) Less(i, j int) bool {
    result := false
    if ug[i] > ug[j] {
        result = true
    }
    return result
}

func (ug usergrade)Swap(i, j int) {
    ug[i], ug[j] = ug[j], ug[i]
}

func main () {
    /*
    var grade = usergrade{
        "heno": 90,
        "haoran": 100
        "xiaopang" : 95,
        }
    */
    var grade = usergrade{3, 4, 2, 8}
    fmt.Println(grade) //[3 4 2 8]
    sort.Sort(grade) //[8 4 3 2]
    fmt.Println(grade)
}
