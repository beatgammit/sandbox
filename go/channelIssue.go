package main

import (
    "web"
    "os"
    "fmt"
    "container/vector"
    "time"
)

var listeners vector.Vector

func listen (ctx *web.Context) {
    c := make(chan []byte)
    listeners.Push(c)

    fileData := <-c
    fmt.Println("Push to client")

    for fileData != nil {
        ctx.Write(fileData)
        fileData = <-c
    }
}

func serve (path string) {
    file, err := os.Open(path, os.O_RDONLY, 0666)

    if err != nil {
        panic(err.String())
    }

    defer listeners.Resize(0, 0)
    defer file.Close()

    buf := make([]byte, 512)
    _, err = file.Read(buf)

    fmt.Println("Serve data")

    for err != os.EOF {
        for _, val := range listeners {
            c := val.(chan []byte)
            c <- buf
        }
        fmt.Println("Read next")
        _, err = file.Read(buf)
    }

    fmt.Println("End of read")

    for _, val := range listeners {
        c := val.(chan []byte)
        c <- nil
    }
}

func socketInit (path string) {
    for {
        time.Sleep(5 * 1000 * 1000000) // sleep for 5 seconds
        serve(path)
    }
}

func main() {
    go socketInit("testServer.go")

    web.Get(".*", listen)
    web.Run("localhost:13240")
}
