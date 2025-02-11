package main

import (
	"fmt"
	"os"

	"hiveshake/cmd"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "hiveshake",
		Short: "HiveShake Chaos Engineering Toolkit",
	}

	rootCmd.AddCommand(
		cmd.CPUCmd(),
		cmd.MemoryCmd(),
		cmd.NetworkCmd(),
		cmd.ProcessCmd(),
		cmd.DiskCmd(),
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
