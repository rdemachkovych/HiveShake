package cmd

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

func stressCPU(duration time.Duration) {
	var wg sync.WaitGroup
	numCPU := runtime.NumCPU()
	fmt.Printf("Starting CPU stress test on %d cores for %s...\n", numCPU, duration)

	for i := 0; i < numCPU; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			end := time.Now().Add(duration)
			for time.Now().Before(end) {
				_ = 1 // Busy loop to keep CPU busy
			}
		}()
	}
	wg.Wait()
	fmt.Println("CPU stress test completed.")
}

func CPUCmd() *cobra.Command {
	var duration time.Duration
	cmd := &cobra.Command{
		Use:   "cpu",
		Short: "Stress test CPU",
		Run: func(cmd *cobra.Command, args []string) {
			stressCPU(duration)
		},
	}
	cmd.Flags().DurationVarP(&duration, "duration", "d", 10*time.Second, "Duration of CPU stress test")
	return cmd
}
