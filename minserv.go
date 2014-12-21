package main

import (
	"flag"
	"fmt"
	"net/http"
	"path/filepath"
)

var path = flag.String("path", ".", "directory to serve")
var port = flag.Int("port", 8080, "port to use")

func main() {
	flag.Parse()
	wd, err := filepath.Abs(*path)
	if err != nil {
		fmt.Println(err)
	}
	portStr := fmt.Sprintf(":%v", *port)
	fmt.Printf("Serving %v on port %v...\n", wd, portStr)
	fmt.Println(http.ListenAndServe(portStr, http.FileServer(http.Dir(wd))))
}
