package commands

import (
	"fmt"
	"os"

	"main.go/objects"
)

func CatFile(args []string) {
	// if len(args) < 1 {
	// 	fmt.Fprintln(os.Stderr, "usage:mygit cat -p <hash>")
	// 	os.Exit(1)
	// }

	if len(args) < 2 || args[0] != "-p" {
		fmt.Fprintln(os.Stderr, "usage:mygit cat -p <hash>")
	}

	hash := args[1]

	content, err := objects.ReadBlob(hash)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading the blob: %s\n", err)
		os.Exit(1)
	}

	fmt.Print(string(content))
}
