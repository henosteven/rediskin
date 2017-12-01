package main

import (
    "fmt"
    "time"
)

func main() {
    /*
    75
    argv: some-argv
    entry: 2017-11-30 10:58:48.652150296 +0800 CST m=+0.000407751
    exit: 2017-11-30 10:58:53.653650565 +0800 CST m=+5.001811020
    */
    fmt.Println(changeResult(5))
    doBigJob()
}

func doBigJob() {
    defer trace("some-argv")()
    time.Sleep(time.Second * 5)
}

func changeResult(i int64) (result int64) {
    defer func() {
        result += 10
    } ()
    result = i + 60
    return 
}

func trace(someargv string) func() {
    fmt.Println("argv:", someargv)
    fmt.Println("entry:", time.Now())
    return func() {
        fmt.Println("exit:", time.Now())
    }
}


