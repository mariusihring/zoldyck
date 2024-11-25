package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
)

func main() {

	rootCmd := cobra.Command{
		Use:     "zoldyck",
		Aliases: []string{"zol"},
		Short:   "A wrapper around Git commands",
		Args:    cobra.ArbitraryArgs,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("Usage: zoldyck [arguments]")
				os.Exit(1)
			}

			gitCmd := exec.Command("git", args...)

			gitCmd.Stdout = os.Stdout
			gitCmd.Stderr = os.Stderr

			err := gitCmd.Run()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error executing git command: %v\n", err)
				os.Exit(1)
			}

		},
	}

	commitCmd := cobra.Command{
		Use:     "commit",
		Aliases: []string{"c"},
		Short:   "Create a commit",
		Run: func(cmd *cobra.Command, args []string) {
			runCommit()
		},
	}

	// Add all the commands to the root command
	rootCmd.AddCommand(&commitCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)

	}
}
