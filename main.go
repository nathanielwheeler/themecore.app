package main

import (
	"fmt"
	"os"

	"themecore.app/server"

	_ "github.com/heroku/x/hmetrics/onload"
)

func main() {
	if err := server.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}
