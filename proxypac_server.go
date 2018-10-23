package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

const (
	defaultPort = 58000
)

func main() {
	port := flag.Int("p", defaultPort, "Port number to listen")
	flag.Parse()

	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage:  %s [options] pacfile\n", os.Args[0])
		flag.PrintDefaults()
	}

	if len(flag.Args()) == 0 {
		flag.Usage()
		os.Exit(0)
	}

	pacfile := flag.Args()[0]
	http.HandleFunc("/proxy.pac", func(w http.ResponseWriter, r *http.Request) {
		// Use Go 1.6 or higher for preventing directory traversal.
		http.ServeFile(w, r, pacfile)
	})

	fmt.Printf("Listening port: %d\n", *port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}
