package main

import (
	"bridge/api"
	"errors"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// callBridge()
	mux := http.NewServeMux()
	mux.HandleFunc("/", api.Handler)

	err := http.ListenAndServe(":3333", mux)
	if errors.Is(err, http.ErrServerClosed) {
		fmt.Printf("server closed\n")
	} else if err != nil {
		fmt.Printf("error starting server: %s\n", err)
		os.Exit(1)
	}

}
