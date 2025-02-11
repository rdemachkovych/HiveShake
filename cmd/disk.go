package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func DiskCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "disk",
		Short: "Fill disk space for testing",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Disk space filling started")
			// TODO: Implement disk filling function
		},
	}
}
