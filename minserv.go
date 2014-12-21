package main

import (
        "flag"
        "fmt"
        "net/http"
        "os"
        "strconv"
)

var path = flag.String("path", ".", "directory to serve")
var port = flag.Int("port", 8080, "port to use")

func main() {
        flag.Parse()
        var wd string
        if *path == "." {
                wd, _ = os.Getwd()
        } else {
                wd = *path
        }
        fmt.Printf("Serving %v on port %v...\n", wd, *port)
        http.ListenAndServe(":"+strconv.Itoa(*port), http.FileServer(http.Dir(wd)))
}
