package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    buf := bufio.NewWriterSize(os.Stdout, 0)
    fmt.Println(buf.Available(), buf.Buffered()) //4096 0
    
    buf.WriteString("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
    fmt.Println(buf.Available(), buf.Buffered()) //4070 26

    buf.Flush() //ABCDEFGHIJKLMNOPQRSTUVWXYZ
}

