package cmd

import (
	"fmt"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
)

func stressNetwork(interfaceName string, delay time.Duration) {
	// Check if 'tc' command is available
	if _, err := exec.LookPath("tc"); err != nil {
		fmt.Println("Error: 'tc' command not found in PATH")
		return
	}

	fmt.Printf("Adding %s network latency on %s...\n", delay, interfaceName)
	cmd := exec.Command("tc", "qdisc", "add", "dev", interfaceName, "root", "netem", "delay", fmt.Sprintf("%dms", delay.Milliseconds()))
	if err := cmd.Run(); err != nil {
		fmt.Println("Error applying network stress:", err)
		return
	}
	time.Sleep(delay)
	fmt.Println("Removing network latency...")
	exec.Command("tc", "qdisc", "del", "dev", interfaceName, "root", "netem").Run()
}

func NetworkCmd() *cobra.Command {
	var interfaceName string
	var delay time.Duration
	cmd := &cobra.Command{
		Use:   "network",
		Short: "Introduce network latency",
		Run: func(cmd *cobra.Command, args []string) {
			stressNetwork(interfaceName, delay)
		},
	}
	cmd.Flags().StringVarP(&interfaceName, "interface", "i", "eth0", "Network interface to impact")
	cmd.Flags().DurationVarP(&delay, "delay", "d", 100*time.Millisecond, "Network latency duration")
	return cmd
}
