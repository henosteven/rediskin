package main

import (
    "fmt"
    "bufio"
    "strings"
)

func main() {
    sr := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
    buf := bufio.NewReaderSize(sr, 0) // 0 < minReadBufferSize 所以size默认为minReadBufferSize
    fmt.Println("buf.Buffered:", buf.Buffered()) // 0

    s, _ := buf.Peek(5)
    s[0] = 'a' //'' => byte || s is slice
    fmt.Printf("buf.Buffered:%d s:%q\n", buf.Buffered(), s) //buf.Buffered:16 s:"aBCDE"
    
    /*
    6 "aBCDEFGHIJ" <nil>
    0 "KLMNOP" <nil>
    6 "QRSTUVWXYZ" <nil>
    0 "123456" <nil>
    0 "7890" <nil>
    0 "" EOF
    */
    b := make([]byte, 10)
    for n, err := 0, error(nil); err == nil; {
        n, err = buf.Read(b)  //Read方法中，如果r和w位置一致也就是读完了，会b.rd.Read(b.buf)重新读取
        fmt.Printf("%d %q %v\n", buf.Buffered(), b[:n], err)
    }
}
