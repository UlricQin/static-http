package main

import (
	"fmt"
	"net/http"
	"os"
	"runtime"
)

func main() {
	args := os.Args
	l := len(args)

	pwd, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "os.Getwd() fail. error: %s", err)
		os.Exit(1)
	}

	port := "8888"
	if l == 2 {
		// http port given
		port = args[1]
	}

	runtime.GOMAXPROCS(runtime.NumCPU())

	fmt.Printf("static http server serve at: %s ...\n", port)
	panic(http.ListenAndServe(":"+port, http.FileServer(http.Dir(pwd))))
}
