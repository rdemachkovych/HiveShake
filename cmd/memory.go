package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

func stressMemory(sizeMB int, duration time.Duration) {
	fmt.Printf("Allocating %d MB of memory for %s...\n", sizeMB, duration)
	data := make([]byte, sizeMB*1024*1024)
	for i := range data {
		data[i] = 1 // Fill memory to force allocation
	}
	time.Sleep(duration)
	fmt.Println("Memory stress test completed.")
}

func MemoryCmd() *cobra.Command {
	var sizeMB int
	var duration time.Duration
	cmd := &cobra.Command{
		Use:   "memory",
		Short: "Simulate memory stress",
		Run: func(cmd *cobra.Command, args []string) {
			stressMemory(sizeMB, duration)
		},
	}
	cmd.Flags().IntVarP(&sizeMB, "size", "s", 100, "Memory size in MB")
	cmd.Flags().DurationVarP(&duration, "duration", "d", 10*time.Second, "Duration of memory stress test")
	return cmd
}
