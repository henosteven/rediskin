package main

import (
    "fmt"
)

func noempty(s []string) []string {
    i := 0
    for _, v := range s {
        if v != "" {
            s[i] = v
            i++
        }
    }
    return s[:i]
} 

func realnoempty(s []string) []string {
    var z []string
    for _, v := range s {
        if v != "" {
            z = append(z, v)
        }
    }
    return z
}

func reverse(s []int) []int {
    var i, j int
    for i, j = 0, len(s) - 1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i] 
    }
    return s
}

func remove(s []int, i int) []int {
    if len(s) <= i {
        panic("out of range")
    }
    copy(s[i:], s[i+1:])
    return s[:(len(s)-1)]
}

func removeIgnoreOrder(s []int, i int) []int {
    if len(s) <= i {
        panic("out of range")
    }
    s[i] = s[len(s)-1]
    return s[:len(s) - 1]
}

func main () {
    
    var s = []string{"heno", "", "xiaopang"}
    //fmt.Println(noempty(s)) //[heno xiaopang]
    //fmt.Println(s) // [heno xiaopang xiaopang]

    fmt.Println(realnoempty(s)) //[heno xiaopang]
    fmt.Println(s) // [heno  xiaopang]

    var ns = []int{2, 4, 6, 8, 10, 12}
    //fmt.Println(reverse(ns))
    //fmt.Println(remove(ns, 2))
    fmt.Println(removeIgnoreOrder(ns, 2))
}
