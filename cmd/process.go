package cmd

import (
	"fmt"
	"syscall"

	"github.com/spf13/cobra"
)

func killProcess(pid int) error {
	// Send SIGKILL to the specified PID
	return syscall.Kill(pid, syscall.SIGKILL)
}

func ProcessCmd() *cobra.Command {
	var pid int

	cmd := &cobra.Command{
		Use:   "process",
		Short: "Kill a process by PID",
		Run: func(cmd *cobra.Command, args []string) {
			if pid == 0 {
				fmt.Println("Please provide a valid PID using --pid")
				return
			}

			err := killProcess(pid)
			if err != nil {
				fmt.Printf("Failed to kill process %d: %v\n", pid, err)
			} else {
				fmt.Printf("Process %d terminated successfully.\n", pid)
			}
		},
	}

	cmd.Flags().IntVarP(&pid, "pid", "p", 0, "PID of the process to kill")

	return cmd
}
