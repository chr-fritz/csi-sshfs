package main

import (
	"flag"
	"fmt"
	"github.com/chr-fritz/csi-sshfs/pkg/sshfs"
	"os"

	"github.com/spf13/cobra"
)

var (
	endpoint  string
	nodeID    string
	BuildTime string
)

func init() {
	flag.Set("logtostderr", "true")
}

func main() {

	flag.CommandLine.Parse([]string{})

	cmd := &cobra.Command{
		Use:   "NFS",
		Short: "CSI based NFS driver",
		Run: func(cmd *cobra.Command, args []string) {
			handle()
		},
	}

	cmd.Flags().AddGoFlagSet(flag.CommandLine)

	cmd.PersistentFlags().StringVar(&nodeID, "nodeid", "", "node id")
	cmd.MarkPersistentFlagRequired("nodeid")

	cmd.PersistentFlags().StringVar(&endpoint, "endpoint", "", "CSI endpoint")
	cmd.MarkPersistentFlagRequired("endpoint")

	versionCmd := &cobra.Command{
		Use:   "version",
		Short: "Prints information about this version of csi sshfs plugin",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf(`CSI-SSHFS Plugin
Version:    %s
Build Time: %s
`, sshfs.Version, BuildTime)
		},
	}

	cmd.AddCommand(versionCmd)
	versionCmd.ResetFlags()

	cmd.ParseFlags(os.Args[1:])
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}

func handle() {
	d := sshfs.NewDriver(nodeID, endpoint)
	d.Run()
}
