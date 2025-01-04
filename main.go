package main

import (
	"github.com/getcharzp/light-proxy/cmd"
	"log"
)

func main() {
	rootCmd := cmd.NewRootCommand()

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
