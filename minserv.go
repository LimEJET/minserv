package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

var path = flag.String("path", ".", "Directory to serve.")
var port = flag.Int("port", 8080, "Port to use.")

func main() {
	flag.Parse()
	wd, err := filepath.Abs(*path)
	if err != nil {
		fmt.Println(err)
		return
	}
	if info, err := os.Stat(wd); !info.IsDir() || os.IsNotExist(err) {
		fmt.Println(*path, "does not exist or is not a directory.")
		return
	}

	portStr := fmt.Sprint(*port)

	fmt.Printf("MinServ: Now serving %v on port %v...\n", wd, portStr)
	fmt.Println(http.ListenAndServe(":"+portStr, http.FileServer(http.Dir(wd))))
}
