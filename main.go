package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// 1. Prepare the container.
	// Its name is “port”, the default is 3000, and its description is “port number”.
	port := flag.Int("port", 3000, "port for proxy")
	origin := flag.String("origin", "", "URL origin")

	// 2. this must be called so go can read what you type in terminal
	flag.Parse()
	// Handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "reaquesting data from path: %s", r.URL.Path)
		fmt.Printf("request entered into: %s\n", r.URL.Path)
	})

	// Run the server using variabel port
	addr := fmt.Sprintf(":%d", *port)
	fmt.Println("Server running...")
	http.ListenAndServe(addr, nil)

	// 3. Check if orgin has been filled or not
	if *origin == "" {
		fmt.Println("Origin must be filled in")
		os.Exit(1)
	}

	fmt.Printf("Proxy is ready on port %d, shooting to: %s\n", port, origin)
}
