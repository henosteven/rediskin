package main

import (
    "bufio"
    "fmt"
    "strings"
)

func main() {
    const input = "1,2,3,4,5,"
    scanner := bufio.NewScanner(strings.NewReader(input))
    scanner.Split(splitComma)
    for scanner.Scan() {
        fmt.Printf("%q\n", scanner.Text()) 
    }

    if err := scanner.Err(); err != nil {
        fmt.Println(err)
    }
}

func splitComma(data []byte, atEOF bool) (advance int, token []byte, err error) {
    if len(data) == 0 {
        return 0, nil, nil
    }

    for i := 0; i < len(data) ; i++ {
        if (data[i] == ',') {
            return i+1, data[:i], nil
        } 
    }

    if atEOF {
        return 0, data, bufio.ErrFinalToken 
    } else {
        return 0, nil, nil
    }
}
