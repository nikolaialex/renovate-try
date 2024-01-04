package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, Modules!")

		},
	}

	fmt.Println("Calling cmd.Execute()!")
	cmd.Execute()
}
