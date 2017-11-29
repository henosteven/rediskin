package main

import (
    "fmt"
    "time"
    "log"
    "os/exec"
    "strings"
    "bytes"
    "context"
)

func main() {
    cmd := exec.Command("tr", "a-z", "A-Z")
    cmd.Stdin = strings.NewReader("some input")
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("in all caps: %q\n", out.String())

    ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
    defer cancel()
    if err := exec.CommandContext(ctx, "sleep", "5").Run(); err != nil {
        log.Fatal("failed-timeout-my-god")
    }
}
