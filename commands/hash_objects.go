package commands

import (
	"fmt"
	"os"

	"main.go/objects"
)

func HashObject(args []string) {
	if len(args) < 2 || args[0] != "-w" {
		fmt.Fprintln(os.Stderr, "usage : mygit hash-object -w <file>")
		os.Exit(1)
	}

	filePath := args[1]

	hash, err := objects.CreateBlob(filePath)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error creating blob %s\n ", err)
		os.Exit(1)
	}

	fmt.Println(hash)
}
