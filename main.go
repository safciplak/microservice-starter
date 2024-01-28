package main

import (
	"fmt"
	"os"
)

// main bootstraps the application
func main() {
	var application = InitializeService()

	if err := application.Start(); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
